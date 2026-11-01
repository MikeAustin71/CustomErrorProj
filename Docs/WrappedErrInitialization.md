
# Wrapped Error Issue

Given the following Golang type:

```go
// PointerNilError
//
// is a custom error type for nil pointer errors.
//
// Typically used when a function is called with a
// nil pointer as an argument.
type PointerNilError struct {
FunctionCallChain string // ePrefix Function Call Chain
SourceFile       string
ParentType       string
ThisFunction     string
NilParameterName string
ErrContext       string
Message          string
WrappedErr       error
}

func(e *PointerNilError) SetErrorParameters(
functionCallChain string,
sourceFile string,
parentType string,
thisFunction string,
nilParameterName string,
errorContext string,
errorMessage string) error {


if nilParameterName == "" {
return fmt.Errorf("\nPointerNilError\n" +
"nilParameterName is required")
}

if errorMessage == "" {
errorMessage = fmt.Sprintf("%v is a nil pointer!" , e.NilParameterName)
}

e.FunctionCallChain = functionCallChain
e.SourceFile = sourceFile
e.ParentType = parentType
e.ThisFunction = thisFunction
e.NilParameterName = nilParameterName
e.ErrContext = errorContext
e.Message = errorMessage

return nil
}


func (e *PointerNilError) Error() string {

errString := ""

if e.WrappedErr != nil {

errString = "\n"

if e.FunctionCallChain != "" {
errString += fmt.Sprintf("Function Call Chain: %s\n", e.FunctionCallChain)
}

if e.SourceFile != "" {
errString += fmt.Sprintf("Source File: %s\n", e.SourceFile)
}

if e.ParentType != "" {
errString += fmt.Sprintf("Parent Type: %s\n", e.ParentType)
}

if e.ThisFunction != "" {
errString += fmt.Sprintf("Calling Function: %s\n", e.ThisFunction)
}

if e.NilParameterName == "" {
e.NilParameterName = "Unknown Parameter"
}

errString += fmt.Sprintf("Nil Parameter Name: %s\n", e.NilParameterName)

if e.ErrContext != "" {
errString += fmt.Sprintf("Error Context: %s\n", e.ErrContext)
}

if e.Message != "" {
errString += fmt.Sprintf("%s\n", e.Message)
}

e.WrappedErr = errors.New(errString)
}

return errString
}

func (e *PointerNilError) Unwrap() error {

return e.WrappedErr
}
```

The problem with this type is that if method 'Unwrap()' is called before Error(), e.WrappedErr will always be nil.

As you can see, e.WrappedErr is set in method Error().

I could call Error() from Unwrap(), but I'm concerned about deadlocks and thread safe operations. I could use a sync mutex and perform the configuration in a separate type.

What options for solution do you see?


# Google Search Response
The design flaw you highlighted stems from lazy initialization paired with state mutation inside read-only methods (Error() and Unwrap()). In Go, custom error types should ideally treat their internal fields as immutable once created, or handle dynamic string generation without mutating the struct state.
Calling Error() from Unwrap() won't cause a deadlock on a single goroutine, but mutating e.WrappedErr inside these methods makes your custom error completely thread-unsafe. If multiple goroutines log this error simultaneously, your application will crash with a concurrent map/struct write panic.
Here are the best architectural options to fix this issue, ranked from simplest to most advanced.

## Option 1: Eager Initialization (Recommended)

The cleanest, most idiomatic solution in Go is to populate WrappedErr immediately inside SetErrorParameters rather than delaying it. This eliminates the need for any sync mechanisms, as the struct becomes read-only by the time Error() or Unwrap() are called.

```go
func(e *PointerNilError) SetErrorParameters(/* ... args ... */) error {
if nilParameterName == "" {
return fmt.Errorf("\nPointerNilError\nnilParameterName is required")
}
if errorMessage == "" {
// FIX: Use nilParameterName variable, not e.NilParameterName which is empty right now
errorMessage = fmt.Sprintf("%v is a nil pointer!", nilParameterName)
}

e.FunctionCallChain = functionCallChain
e.SourceFile = sourceFile
e.ParentType = parentType
e.ThisFunction = thisFunction
e.NilParameterName = nilParameterName
e.ErrContext = errorContext
e.Message = errorMessage

// EAGERLY INITIALIZE: Generate the string and wrap it right here
e.WrappedErr = errors.New(e.buildErrorString())

return nil
}

// Private helper to build the string cleanly without mutating state
func (e *PointerNilError) buildErrorString() string {
errString := "\n"
if e.FunctionCallChain != "" { errString += fmt.Sprintf("Function Call Chain: %s\n", e.FunctionCallChain) }
if e.SourceFile != ""        { errString += fmt.Sprintf("Source File: %s\n", e.SourceFile) }
if e.ParentType != ""        { errString += fmt.Sprintf("Parent Type: %s\n", e.ParentType) }
if e.ThisFunction != ""      { errString += fmt.Sprintf("Calling Function: %s\n", e.ThisFunction) }

param := e.NilParameterName
if param == "" { param = "Unknown Parameter" }
errString += fmt.Sprintf("Nil Parameter Name: %s\n", param)

if e.ErrContext != ""        { errString += fmt.Sprintf("Error Context: %s\n", e.ErrContext) }
if e.Message != ""           { errString += fmt.Sprintf("%s\n", e.Message) }
return errString
}

func (e *PointerNilError) Error() string {
if e.WrappedErr == nil { return "" }
return e.WrappedErr.Error()
}

func (e *PointerNilError) Unwrap() error {
return e.WrappedErr
}

```

Use code with caution.

## Option 2: Stateless Error() and Dynamic Unwrap()

If you want Unwrap() to return an underlying cause error (like a standard io.EOF or sentinel error), you shouldn't overwrite WrappedErr with a brand-new string error inside the method. Instead, let Error() dynamically compile the string on-the-fly without mutating the struct.


```go
func (e *PointerNilError) Error() string {
// Directly build and return the string without saving it to e.WrappedErr
return e.buildErrorString()
}

func (e *PointerNilError) Unwrap() error {
// Returns the actual underlying error passed during initialization (if any)
return e.WrappedErr
}
Use code with caution.
• Why this works: It removes all side effects from Error() and Unwrap(). Both methods become safe to call concurrently by multiple goroutines.

Option 3: Thread-Safe Lazy Initialization (sync.Once)

If your string generation is incredibly computationally heavy and you must lazily initialize it exactly once across multiple concurrent threads, do not use a raw sync.Mutex. Use sync.Once, which is specifically optimized for this exact use case.
go
import "sync"

type PointerNilError struct {
// ... your other fields ...
WrappedErr       error
initOnce         sync.Once // Adds protection for lazy loading
}

func (e *PointerNilError) init() {
e.initOnce.Do(func() {
// This block is guaranteed to run exactly once, safely,
// even if called by multiple goroutines at the same time.
e.WrappedErr = errors.New(e.buildErrorString())
})
}

func (e *PointerNilError) Error() string {
e.init()
return e.WrappedErr.Error()
}

func (e *PointerNilError) Unwrap() error {
e.init()
return e.WrappedErr
}

```

Use code with caution.

## Summary of Hidden Bugs in the Original Snippet

While addressing the Unwrap() sequence issue, note two other logic bugs in the original code:
1. The Catch-22 Condition: if e.WrappedErr != nil wrapping the inside of your original Error() method meant that if WrappedErr was initially nil, the string generation block would never run, and Error() would always return an empty string "".
2. Field Evaluation Order: Inside SetErrorParameters(), the errorMessage evaluation relied on e.NilParameterName before e.NilParameterName = nilParameterName was assigned, resulting in a blank value in the string.
If you want to keep refining this error type, let me know:
• Is WrappedErr intended to hold an underlying root cause error from another package, or is it strictly meant to hold the formatted text block of this specific error?
• Would you like to see how to implement a Constructor Function (NewPointerNilError) to make this instantiation cleaner and remove the two-step initialization process?

# Request # 2 - Let's look at a Constructor

Yes, please show me how to implement a Constructor Function (NewPointerNilError). Bear in mind that the
current type configuration has public string elements which means they can be manipulated externally. We
may need to refactor these string elements as private and require instance configuration through the Constructor.



# Google Response # 2

Refactoring the fields to be **private (unexported)** is a brilliant move. In Go, hiding the internal state of a struct is the standard way to guarantee **encapsulation** and prevent unauthorized external modification after creation.

By making the fields lowercase and forcing instantiation through a strict constructor function (`NewPointerNilError`), you ensure that your custom error remains **immutable** and **thread-safe** once it is returned to the caller.

Here is the complete, refactored implementation of your custom error package using a secure Constructor Function pattern.

------

## Refactored Custom Error Implementation

```go
package errors

import (
	"errors"
	"fmt"
)

// PointerNilError is a custom error type for nil pointer errors.
// All configuration fields are now private (unexported) to ensure immutability.
type PointerNilError struct {
	functionCallChain string
	sourceFile         string
	parentType         string
	thisFunction       string
	nilParameterName   string
	errContext         string
	message            string
	wrappedErr         error
}

// NewPointerNilError is the constructor function that enforces validation,
// handles encapsulation, and eagerly prepares the error state.
func NewPointerNilError(
	functionCallChain string,
	sourceFile string,
	parentType string,
	thisFunction string,
	nilParameterName string,
	errContext string,
	errorMessage string,
) (*PointerNilError, error) {

	// Enforce strict required validations immediately
	if nilParameterName == "" {
		return nil, fmt.Errorf("PointerNilError initialization failed: nilParameterName is required")
	}

	// Default fallback message if empty
	if errorMessage == "" {
		errorMessage = fmt.Sprintf("%s is a nil pointer!", nilParameterName)
	}

	// Construct the concrete struct pointer
	err := &PointerNilError{
		functionCallChain: functionCallChain,
		sourceFile:         sourceFile,
		parentType:         parentType,
		thisFunction:       thisFunction,
		nilParameterName:   nilParameterName,
		errContext:         errContext,
		message:            errorMessage,
	}

	// Eagerly compute the string error block and wrap it.
	// This makes Error() and Unwrap() completely read-only and thread-safe.
	err.wrappedErr = errors.New(err.buildErrorString())

	return err, nil
}

// buildErrorString acts as an internal helper to format the output block
func (e *PointerNilError) buildErrorString() string {
	errString := "\n"

	if e.functionCallChain != "" {
		errString += fmt.Sprintf("Function Call Chain: %s\n", e.functionCallChain)
	}
	if e.sourceFile != "" {
		errString += fmt.Sprintf("Source File: %s\n", e.sourceFile)
	}
	if e.parentType != "" {
		errString += fmt.Sprintf("Parent Type: %s\n", e.parentType)
	}
	if e.thisFunction != "" {
		errString += fmt.Sprintf("Calling Function: %s\n", e.thisFunction)
	}

	errString += fmt.Sprintf("Nil Parameter Name: %s\n", e.nilParameterName)

	if e.errContext != "" {
		errString += fmt.Sprintf("Error Context: %s\n", e.errContext)
	}
	if e.message != "" {
		errString += fmt.Sprintf("%s\n", e.message)
	}

	return errString
}

// Error satisfies Go's built-in error interface.
// Because e.wrappedErr was initialized eagerly, this is completely thread-safe.
func (e *PointerNilError) Error() string {
	if e.wrappedErr == nil {
		return ""
	}
	return e.wrappedErr.Error()
}

// Unwrap satisfies Go's standard unwrap interface (`errors.Unwrap`).
// Safe to call at any time, in any order relative to Error().
func (e *PointerNilError) Unwrap() error {
	return e.wrappedErr
}
```

Use code with caution.

------

## Architectural Design Improvements

1. **Unexported Struct Fields:** By changing the field names to lowercase (`functionCallChain`, `sourceFile`, etc.), code outside of your package can no longer run `myErr.Message = "hacked"` or bypass your validation.
2. **Failure in the Constructor:** If a developer tries to create this error without providing a `nilParameterName`, the constructor returns `nil, error` explicitly. This flags problems immediately during runtime testing.
3. **Elimination of Side Effects:** Both `Error()` and `Unwrap()` now have a single job: reading `e.wrappedErr`. They never modify structural variables. No mutexes or atomic flags are required, maximizing performance.

------

## How to use it in your code

Here is an example of a function catching a nil pointer argument and utilizing your new constructor pattern:

go

```go
package main

import (
	"fmt"
	"yourproject/errors" // Assuming your custom error package name
)

type Database struct{}

func ProcessUser(db *Database, userID string) error {
	if db == nil {
		// Use the constructor pattern safely
		err, initErr := errors.NewPointerNilError(
			"main() -> ProcessUser()",
			"user_service.go",
			"Database",
			"ProcessUser",
			"db",
			"Fetching active user record",
			"", // Leaving blank to leverage your default fallback message
		)
		if initErr != nil {
			return initErr // Handles constructor configuration failure
		}
		return err
	}

	return nil
}
```

Use code with caution.

Now that the fields are tightly encapsulated, let me know:

- Do you want to learn how to add **custom Getter methods** (e.g., `func (e *PointerNilError) NilParameterName() string`) if other packages still need read-only access to those values?
- Would you like to explore how to extract the `SourceFile` and `ThisFunction` **automatically** using Go's built-in `runtime.Caller` package, so developers don't have to type them out manually?
