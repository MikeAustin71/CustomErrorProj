



# ***ErrorGenericConfig*** Returning Specific Custom Types

The following Golang type is working well for me. One problem. I would like to be able to inspect
wrapped errors and pick out ErrorNativeToFunction, ErrorFromFunctionCall, ErrorInputParameterOutOfRange
and ErrorInputParameterNilPointer. To do this I would have to find some way to

  1. Find a way to cast the returned error as a particular type of error.
       -- OR --
  2. Return a specialized type like ErrorFromFunctionCall and let the calling function call ErrorFromFunctionCall.Error.

What are my options?

```go
package customerrors

import (
  "errors"
  "fmt"
)


type ErrorGenericConfig struct {
  sourceFile     string
  parentType     string // The 'type' in which this method resides
  thisFunction   string
  calledFunction string
  parameter      string
  path           string
  fileName       string
  url            string
  errorContext   string
  errorMessage   string
}

// NewErrorGenericConfig
//
// Designed to be called at beginning of a function when
// only three elements are known: sourceFile, parentType, and thisFunction.
func (egc *ErrorGenericConfig) NewErrorGenericConfig(
  sourceFile string,
  parentType string,
  thisFunction string) *ErrorGenericConfig {

  return &ErrorGenericConfig{
    sourceFile: sourceFile,
    parentType: parentType,
    thisFunction: thisFunction,
  }
}

// GetErrorInputParameterOutOfRange
//
// This method is designed to be called when a function's input
// parameter is out of range.
func (egc *ErrorGenericConfig) GetErrorInputParameterOutOfRange(
  parameter      string,
  path           string,
  fileName       string,
  url            string,
  errorContext   string,
  errorMessage   string) error {

  egc.parameter = parameter
  egc.path = path
  egc.fileName = fileName
  egc.url = url
  egc.errorContext = errorContext
  egc.errorMessage = errorMessage

  return errors.New(egc.buildErrorStr())
}

// GetErrorInputParameterNilPointer
//
// This method is designed to be called when a function's input
// parameter is identified as a nil pointer.
func (egc *ErrorGenericConfig) GetErrorInputParameterNilPointer(
  parameter      string,
  path           string,
  fileName       string,
  url            string,
  errorContext   string,
  errorMessage   string) error {

  egc.parameter = parameter
  egc.path = path
  egc.fileName = fileName
  egc.url = url
  egc.errorContext = errorContext
  egc.errorMessage = errorMessage

  return errors.New(egc.buildErrorStr())
}

// GetErrorNativeToFunction
//
// Represents an error with detailed context about a function's internal
// operational failure.
//
// This custom error type is designed to be used by functions or methods which
// encounter operational errors or failures during processing. These types of
// errors are NOT errors returned when calling other functions.
//
// Function operational errors occur due failures or errors inherent to the
// current function or method. Examples include, invalid input parameter values, internal
// calculation errors and value out of bounds errors.
//
// The returned error refers to the current function or method only.
func (egc *ErrorGenericConfig) GetErrorNativeToFunction(
  parameter      string,
  path           string,
  fileName       string,
  url            string,
  errorContext   string,
  errorMessage   string) error {

  egc.parameter = parameter
  egc.path = path
  egc.fileName = fileName
  egc.url = url
  egc.errorContext = errorContext
  egc.errorMessage = errorMessage

  return errors.New(egc.buildErrorStr())
}

// GetErrorFromFunctionCall
//
// Returns a wrapped error created by adding context to the error
// returned from a called function.
func (egc *ErrorGenericConfig) GetErrorFromFunctionCall(
  calledFunction string,
  errorContext   string,
  errorMessage   string,
  returnedCalledFunctionError error) error {

  egc.calledFunction = calledFunction
  egc.errorContext = errorContext
  egc.errorMessage = errorMessage

  errorStr := egc.buildErrorStr()

  if returnedCalledFunctionError == nil {
    return fmt.Errorf("%s\n%s", errorStr, "Returned error from called Function is nil!")
  }


  return fmt.Errorf("%s: %w", errorStr, returnedCalledFunctionError)
}

// SetDetailErrorContext
//
// Sets the detail error context for the ErrorGenericConfig.
// If input parameters are empty, they will be ignored.
func (egc *ErrorGenericConfig) SetDetailErrorContext(
  parameter string,
  path string,
  fileName string,
  url string) {

  if parameter != "" {
    egc.parameter = parameter
  }

  if path != "" {
    egc.path = path
  }

  if fileName != "" {
      egc.fileName = fileName
  }

  if url != "" {
    egc.url = url
  }

  return
}

// SetErrorMessageElements
//
// Updates the error context and error message fields if non-empty
// values are provided.
func (egc *ErrorGenericConfig) SetErrorMessageElements(
  errorContext string,
  errorMessage string) {

  if errorContext != "" {
    egc.errorContext = errorContext
  }

  if errorMessage != "" {
    egc.errorMessage = errorMessage
  }

  return
}

// buildErrorStr
//
// This private method constructs a string representation of the
// error details extracted from the ErrorGenericConfig structure
// elements.
func (egc *ErrorGenericConfig) buildErrorStr() string {

  errStr := ""

  foundParametersCount := 0

  if egc.sourceFile != "" {
    errStr = "\n"
    errStr += fmt.Sprintf("Source File: %s\n", egc.sourceFile)
    foundParametersCount++
  }

  if egc.parentType != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("Parent Type: %s\n", egc.parentType)

    foundParametersCount++
  }

  if egc.thisFunction != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("This Function: %s\n", egc.thisFunction)

    foundParametersCount++
  }

  if egc.calledFunction != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("Called Function: %s\n", egc.calledFunction)

    foundParametersCount++

  }

  if egc.parameter != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("Parameter: %s\n", egc.parameter)

    foundParametersCount++

  }

  if egc.path != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("Path: %s\n", egc.path)

    foundParametersCount++

  }

  if egc.fileName != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("FileName: %s\n", egc.fileName)

    foundParametersCount++

  }

  if egc.url != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("URL: %s\n", egc.url)

    foundParametersCount++

  }

  if egc.errorContext != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("Error Context: %s\n", egc.errorContext)

    foundParametersCount++
  }

  if egc.errorMessage != "" {

    if foundParametersCount == 0 {
      errStr = "\n"
    }

    errStr += fmt.Sprintf("Error Message: %s\n", egc.errorMessage)

    foundParametersCount++
  }

  if foundParametersCount == 0 {
    errStr += "\nNo parameters found for this ErrorFunctionCall instance\n"
  }

  return errStr
}


```

To inspect wrapped errors programmatically using `errors.As` or `errors.Is`, **the returned errors must be distinct Go types or hold distinct sentinel values.**

Right now, your methods return standard, flat errors via `errors.New` or string formatting. Once a value passes through `errors.New`, its unique code identity is erased, leaving only a string that cannot be cast back into a specific error type.

To solve this, you have **two main structural options** that align with your proposed approaches.

------

## Option 1: The Struct-Based Type Approach (Recommended)

This approach satisfies your first point (casting/inspecting via types). Instead of returning `errors.New`, you define unique struct types for each error category. Each struct embeds your shared configuration data and implements the `error` interface.

This lets callers use Go's idiomatic **`errors.As`** to pluck your custom type right out of an error chain, even if it was wrapped multiple times down the line.

go

```
package customerrors

import (
	"errors"
	"fmt"
)

// 1. Define distinct types that satisfy the error interface
type OutOfRangeError struct{ ErrorGenericConfig }
func (e *OutOfRangeError) Error() string { return e.buildErrorStr() }

type NilPointerError struct{ ErrorGenericConfig }
func (e *NilPointerError) Error() string { return e.buildErrorStr() }

type NativeFunctionError struct{ ErrorGenericConfig }
func (e *NativeFunctionError) Error() string { return e.buildErrorStr() }

// This type also implements Unwrap() so errors.As can dig past it
type FunctionCallError struct {
	ErrorGenericConfig
	Unwrapped error
}
func (e *FunctionCallError) Error() string {
	return fmt.Sprintf("%s: %v", e.buildErrorStr(), e.Unwrapped)
}
func (e *FunctionCallError) Unwrap() error {
	return e.Unwrapped
}

// 2. Update your builder methods to return these specific struct pointers
func (egc *ErrorGenericConfig) GetErrorInputParameterOutOfRange(parameter, path, fileName, url, errorContext, errorMessage string) error {
	egc.SetDetailErrorContext(parameter, path, fileName, url)
	egc.SetErrorMessageElements(errorContext, errorMessage)
	
	// Return as a pointer to the specific structural type
	return &OutOfRangeError{ErrorGenericConfig: *egc}
}

func (egc *ErrorGenericConfig) GetErrorFromFunctionCall(calledFunction, errorContext, errorMessage string, returnedCalledFunctionError error) error {
	egc.calledFunction = calledFunction
	egc.SetErrorMessageElements(errorContext, errorMessage)

	return &FunctionCallError{
		ErrorGenericConfig: *egc,
		Unwrapped:          returnedCalledFunctionError,
	}
}
```

Use code with caution.

How the caller inspects this:

go

```
err := SomeWorkerFunction()

// The caller can inspect and extract the precise configuration data
var outOfRangeErr *customerrors.OutOfRangeError
if errors.As(err, &outOfRangeErr) {
    fmt.Printf("Caught range issue in file: %s", outOfRangeErr.sourceFile)
}
```

Use code with caution.

------

## Option 2: The Classification Enum Approach

If you want to keep a single underlying struct rather than creating four new types, you can add an **Error Kind** classification field. This matches your second goal of providing a specialized field/method for inspection.

go

```
package customerrors

import (
	"errors"
	"fmt"
)

type ErrorKind int

const (
	KindNative ErrorKind = iota
	KindFromFunctionCall
	KindInputParameterOutOfRange
	KindInputParameterNilPointer
)

type ErrorGenericConfig struct {
	kind           ErrorKind // Added classification field
	sourceFile     string
	// ... rest of your existing fields
	unwrapped      error     // To store the inner error for wrapping chains
}

// Implement Error and Unwrap on your base struct
func (egc *ErrorGenericConfig) Error() string {
	if egc.unwrapped != nil {
		return fmt.Sprintf("%s: %v", egc.buildErrorStr(), egc.unwrapped)
	}
	return egc.buildErrorStr()
}

func (egc *ErrorGenericConfig) Unwrap() error {
	return egc.unwrapped
}

// Accessor for the calling function to inspect the Kind
func (egc *ErrorGenericConfig) Kind() ErrorKind {
	return egc.kind
}

// Update builder methods to tag the classification
func (egc *ErrorGenericConfig) GetErrorInputParameterOutOfRange(/* args */) error {
	egc.kind = KindInputParameterOutOfRange
	// ... assign fields ...
	return &ErrorGenericConfig{kind: egc.kind, sourceFile: egc.sourceFile /* copy fields */}
}
```

Use code with caution.

How the caller inspects this:

go

```
err := SomeWorkerFunction()

var customErr *customerrors.ErrorGenericConfig
if errors.As(err, &customErr) {
	switch customErr.Kind() {
	case customerrors.KindInputParameterOutOfRange:
		// Handle range error
	case customerrors.KindInputParameterNilPointer:
		// Handle nil pointer
	}
}
```

Use code with caution.

------

Direct Trade-Offs

| Strategy                       | Advantages                                                   | Disadvantages                                                |
| ------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **Option 1: Distinct Structs** | Idiomatic Go; cleanest separation; callers use standard `errors.As(err, &TargetType)` directly without writing nested switch-case statements. | Requires declaring extra thin wrapper types for each error category. |
| **Option 2: Kind/Enum Field**  | Keeps your package footprint small; matches a flat structure with one primary type tracking everything. | Requires callers to perform a two-step inspection (cast to the configuration struct first, then check an Enum method). |

*Note on your current code:* In your current implementation, your builder methods mutate the underlying receiver fields directly (`egc.parameter = parameter`). If you reuse the same `ErrorGenericConfig` reference across multiple errors, sequential errors will overwrite previous parameters. Both Options 1 and 2 fix this by returning copy instances (`*egc` or values) so your errors stay immutable once returned.

Let me know if you would like to:

- See a **complete refactor** of your code file utilizing the struct-based approach.
- Explore how to handle **concurrency-safe copying** of your configuration builder state.



# Embedded Structure Question

I don't understand 'type OutOfRangeError struct{ ErrorGenericConfig }'.
Shouldn't that be structure element variable styled like: 
OutOfRangeError struct{errGenericCfg ErrorGenericConfig }
?



## Embedded Structure

What you are looking at is a Go feature called **struct embedding** (or anonymous fields).

By writing `type OutOfRangeError struct{ ErrorGenericConfig }` without a field name, you are embedding the type directly. This gives you two massive advantages for your error handling: **automatic interface fulfillment** and **field promotion**.

Here is how the two styles compare:

1. Embedded Style (What I used)

go

```
type OutOfRangeError struct {
    ErrorGenericConfig // Embedded (No field name)
}
```

Use code with caution.

- **Field Promotion:** All the fields of `ErrorGenericConfig` are "promoted" to the parent struct. You can access them directly on the parent without nesting. For example, you can type `myRangeErr.sourceFile` instead of `myRangeErr.errGenericCfg.sourceFile`.
- **Method Promotion:** Any methods attached to `ErrorGenericConfig` are also promoted. This means `OutOfRangeError` automatically gains your `SetDetailErrorContext()` and `SetErrorMessageElements()` methods without you writing a single line of forwarding code.
- Named Field Style (What you proposed)

go

```
type OutOfRangeError struct {
    errGenericCfg ErrorGenericConfig // Named field
}
```

Use code with caution.

- **Explicit Nesting:** To read or write fields, you are forced to explicitly type out the path every time: `myRangeErr.errGenericCfg.sourceFile`.
- **No Method Inheritance:** `OutOfRangeError` will *not* inherit the helper methods. If you want to configure your error details after creating it, you would have to call `myRangeErr.errGenericCfg.SetDetailErrorContext()`.

------

Direct Syntax Comparison

Here is how you would construct and use each style side-by-side:

| Action               | Embedded Style                               | Named Field Style                              |
| -------------------- | -------------------------------------------- | ---------------------------------------------- |
| **Instantiation**    | `&OutOfRangeError{ErrorGenericConfig: *egc}` | `&OutOfRangeError{errGenericCfg: *egc}`        |
| **Reading a field**  | `err.sourceFile`                             | `err.errGenericCfg.sourceFile`                 |
| **Calling a helper** | `err.SetDetailErrorContext(...)`             | `err.errGenericCfg.SetDetailErrorContext(...)` |

Why Embedded is better for your Custom Errors

The embedded style allows you to cleanly separate your error types so `errors.As` can distinguish them, while completely avoiding duplicate code or deeply nested structures. The calling code gets a clean, flat object to read from.

Would you like to see exactly how **method promotion** looks in practice when passing these embedded errors around your application, or should we look at how to **write a custom `.Error() string` method** specifically for each embedded type?



# Embedded Structure # 2

Embedded Structure is coming back to me. For my Embedded Structure shouldn't I be using pointers, like this: 

```go
OutOfRangeError struct{ *ErrorGenericConfig }
```

You **can** embed it as a pointer, but for custom error handling, **embedding by value (`struct{ ErrorGenericConfig }`) is usually safer and more idiomatic.**

Using a pointer can introduce subtle bugs related to **mutability** and **shared state**, especially since your current factory methods modify the `ErrorGenericConfig` fields in-place.

Here is a breakdown of why embedding by value is typically preferred for errors, and what happens if you choose pointers.

------

The Risk of Embedding a Pointer: Shared State Mutation

Errors in Go are generally expected to be **immutable** snapshots of what went wrong at a specific moment in time. If you embed a pointer, you are storing a reference to a memory address. If that underlying config object changes later, your error's contents change too.

The Shared State Bug Example:

Consider this scenario if you were to use pointers:

go

```
// Imagine a shared config factory used throughout a service/package
factory := &ErrorGenericConfig{sourceFile: "main.go"}

// 1. Create a range error
err1 := factory.GetErrorInputParameterOutOfRange("age", "", "", "", "", "Too old")

// 2. Sometime later, the same factory handles a nil pointer error
err2 := factory.GetErrorInputParameterNilPointer("ctx", "", "", "", "", "Nil context")

// BUG: If embedded as a pointer (*ErrorGenericConfig), err1's parameters 
// have now been overwritten by err2! Inspecting err1 will show "Parameter: ctx".
```

Use code with caution.

------

Comparison: Value vs. Pointer Embedding for Errors

| Consideration          | Value Embedding (`struct{ ErrorGenericConfig }`)             | Pointer Embedding (`struct{ *ErrorGenericConfig }`)          |
| ---------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **Data Safety**        | **High.** It creates a copy of the config data. The error is a permanent snapshot. | **Low.** Multiple errors might point to the exact same config struct, causing later errors to overwrite earlier ones. |
| **Nil Pointers**       | **Safe.** The fields are always initialized (even if empty). Calling `.Error()` never panics. | **Risky.** If you initialize the struct without assigning the inner pointer, calling `.Error()` will cause a **nil pointer panic**. |
| **Memory/Performance** | Copies the fields. For an error path (which should be exceptional), this minor overhead is negligible. | Copies only an 8-byte pointer address.                       |

------

When *Should* You Use Pointer Embedding?

The only time you should embed `*ErrorGenericConfig` is if your `ErrorGenericConfig` structure contains resource-heavy elements (like an active network connection or an open file descriptor) that *cannot* or *should not* be copied. Because your struct consists entirely of metadata strings (`string`), copying by value is safe, lightweight, and prevents state leakage.

Best Practice for Your Constructor

If you use value embedding, you create a safe snapshot by dereferencing your config pointer `*egc` during initialization:

go

```
func (egc *ErrorGenericConfig) GetErrorInputParameterOutOfRange(...) error {
    // ... update fields ...
    
    // Dereferencing (*egc) makes a clean, isolated copy of the data 
    // for this specific error instance.
    return &OutOfRangeError{ErrorGenericConfig: *egc} 
}
```

Use code with caution.

Would you like to look at how to implement a **thread-safe cloning pattern** for your factory to make sure multiple goroutines don't clash when building errors, or should we focus on implementing the **`Unwrap()` method** so your embedded errors work perfectly with `errors.Is`?

# Refactor # 1

Could you show me a **complete refactor** of my code file utilizing the struct-based approach.

Here is the **complete refactor** of your code file using the struct-based approach.

This layout introduces distinct, dedicated error structures for each variant while embedding `ErrorGenericConfig` by value. This ensures each returned error is a **thread-safe, immutable snapshot** that can be effortlessly targeted by callers using `errors.As`.

go

```
package customerrors

import (
	"errors"
	"fmt"
)

// ============================================================================
// 1. Core Configuration Types
// ============================================================================

type ErrorGenericConfig struct {
	sourceFile     string
	parentType     string // The 'type' in which this method resides
	thisFunction   string
	calledFunction string
	parameter      string
	path           string
	fileName       string
	url            string
	errorContext   string
	errorMessage   string
}

// ============================================================================
// 2. Specialized Struct-Based Error Types
// ============================================================================

// ErrorInputParameterOutOfRange represents a parameter value out of bounds.
type ErrorInputParameterOutOfRange struct {
	ErrorGenericConfig // Value embedding ensures an isolated, safe snapshot
}

func (e *ErrorInputParameterOutOfRange) Error() string {
	return e.buildErrorStr()
}

// ErrorInputParameterNilPointer represents a required pointer passing a nil value.
type ErrorInputParameterNilPointer struct {
	ErrorGenericConfig
}

func (e *ErrorInputParameterNilPointer) Error() string {
	return e.buildErrorStr()
}

// ErrorNativeToFunction represents an internal operational failure within a function.
type ErrorNativeToFunction struct {
	ErrorGenericConfig
}

func (e *ErrorNativeToFunction) Error() string {
	return e.buildErrorStr()
}

// ErrorFromFunctionCall represents an error context layer wrapping a downstream failure.
type ErrorFromFunctionCall struct {
	ErrorGenericConfig
	unwrapped error // Keeps track of the original wrapped error
}

func (e *ErrorFromFunctionCall) Error() string {
	errorStr := e.buildErrorStr()
	if e.unwrapped == nil {
		return fmt.Sprintf("%s\nReturned error from called Function is nil!", errorStr)
	}
	return fmt.Sprintf("%s: %v", errorStr, e.unwrapped)
}

// Unwrap enables standard library tools (errors.Is, errors.As) to traverse past this type.
func (e *ErrorFromFunctionCall) Unwrap() error {
	return e.unwrapped
}

// ============================================================================
// 3. Factory Constructors and Builders
// ============================================================================

// NewErrorGenericConfig is designed to be called at the beginning of a function when
// only three elements are known: sourceFile, parentType, and thisFunction.
func (egc *ErrorGenericConfig) NewErrorGenericConfig(
	sourceFile string,
	parentType string,
	thisFunction string) *ErrorGenericConfig {

	return &ErrorGenericConfig{
		sourceFile:   sourceFile,
		parentType:   parentType,
		thisFunction: thisFunction,
	}
}

// GetErrorInputParameterOutOfRange yields a unique ErrorInputParameterOutOfRange error.
func (egc *ErrorGenericConfig) GetErrorInputParameterOutOfRange(
	parameter string,
	path string,
	fileName string,
	url string,
	errorContext string,
	errorMessage string) error {

	egc.parameter = parameter
	egc.path = path
	egc.fileName = fileName
	egc.url = url
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	// Dereferencing (*egc) makes a clean, isolated value copy for this instance
	return &ErrorInputParameterOutOfRange{ErrorGenericConfig: *egc}
}

// GetErrorInputParameterNilPointer yields a unique ErrorInputParameterNilPointer error.
func (egc *ErrorGenericConfig) GetErrorInputParameterNilPointer(
	parameter string,
	path string,
	fileName string,
	url string,
	errorContext string,
	errorMessage string) error {

	egc.parameter = parameter
	egc.path = path
	egc.fileName = fileName
	egc.url = url
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	return &ErrorInputParameterNilPointer{ErrorGenericConfig: *egc}
}

// GetErrorNativeToFunction yields a unique ErrorNativeToFunction error.
func (egc *ErrorGenericConfig) GetErrorNativeToFunction(
	parameter string,
	path string,
	fileName string,
	url string,
	errorContext string,
	errorMessage string) error {

	egc.parameter = parameter
	egc.path = path
	egc.fileName = fileName
	egc.url = url
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	return &ErrorNativeToFunction{ErrorGenericConfig: *egc}
}

// GetErrorFromFunctionCall returns an inspectable, wrapped FunctionCallError.
func (egc *ErrorGenericConfig) GetErrorFromFunctionCall(
	calledFunction string,
	errorContext string,
	errorMessage string,
	returnedCalledFunctionError error) error {

	egc.calledFunction = calledFunction
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	return &ErrorFromFunctionCall{
		ErrorGenericConfig: *egc,
		unwrapped:          returnedCalledFunctionError,
	}
}

// ============================================================================
// 4. Utility Context Setters
// ============================================================================

// SetDetailErrorContext updates the detail variables if non-empty arguments are provided.
func (egc *ErrorGenericConfig) SetDetailErrorContext(
	parameter string,
	path string,
	fileName string,
	url string) {

	if parameter != "" {
		egc.parameter = parameter
	}
	if path != "" {
		egc.path = path
	}
	if fileName != "" {
		egc.fileName = fileName
	}
	if url != "" {
		egc.url = url
	}
}

// SetErrorMessageElements updates context and messaging details if non-empty arguments are provided.
func (egc *ErrorGenericConfig) SetErrorMessageElements(
	errorContext string,
	errorMessage string) {

	if errorContext != "" {
		egc.errorContext = errorContext
	}
	if errorMessage != "" {
		egc.errorMessage = errorMessage
	}
}

// buildErrorStr constructs a string representation of the populated elements.
func (egc *ErrorGenericConfig) buildErrorStr() string {
	errStr := ""
	foundParametersCount := 0

	if egc.sourceFile != "" {
		errStr = "\n"
		errStr += fmt.Sprintf("Source File: %s\n", egc.sourceFile)
		foundParametersCount++
	}

	if egc.parentType != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("Parent Type: %s\n", egc.parentType)
		foundParametersCount++
	}

	if egc.thisFunction != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("This Function: %s\n", egc.thisFunction)
		foundParametersCount++
	}

	if egc.calledFunction != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("Called Function: %s\n", egc.calledFunction)
		foundParametersCount++
	}

	if egc.parameter != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("Parameter: %s\n", egc.parameter)
		foundParametersCount++
	}

	if egc.path != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("Path: %s\n", egc.path)
		foundParametersCount++
	}

	if egc.fileName != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("FileName: %s\n", egc.fileName)
		foundParametersCount++
	}

	if egc.url != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("URL: %s\n", egc.url)
		foundParametersCount++
	}

	if egc.errorContext != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("Error Context: %s\n", egc.errorContext)
		foundParametersCount++
	}

	if egc.errorMessage != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}
		errStr += fmt.Sprintf("Error Message: %s\n", egc.errorMessage)
		foundParametersCount++
	}

	if foundParametersCount == 0 {
		errStr += "\nNo parameters found for this ErrorFunctionCall instance\n"
	}

	return errStr
}
```

Use code with caution.

------

Verification Example

Here is how clean your caller code becomes when evaluating errors using this setup:

go

```
err := ExecuteDatabaseQuery()
if err != nil {
    // Check specifically for OutOfRange parameters
    var rangeErr *customerrors.ErrorInputParameterOutOfRange
    if errors.As(err, &rangeErr) {
        fmt.Println("Handled range variance for parameter:", rangeErr.Parameter)
        return
    }

    // Check for lower-level errors hidden inside Function Calls
    var callErr *customerrors.ErrorFromFunctionCall
    if errors.As(err, &callErr) {
        fmt.Printf("Failure running internal step: %s\n", callErr.CalledFunction)
        return
    }
}
```

Use code with caution.

Let me know if you would like me to draft a **unit test file** showing exactly how to assert these structures, or if you want to look into minimizing the factory methods by moving toward an **options-pattern builder**.



# Question About Wrapped Errors

```
Your code constructs the error ErrorFromFunctionCall as follows:
  
  func (e *ErrorFromFunctionCall) Error() string {

	errorStr := e.buildErrorStr()

	if e.unwrapped == nil {
		return fmt.Sprintf("%s\nReturned error from called Function is nil!", errorStr)
	}
	
	return fmt.Sprintf("%s: %v", errorStr, e.unwrapped)
}

I previously used 

return fmt.Errorf("%s: %w", errorStr, returnedCalledFunctionError)

to construct the 'wrapped' error. 

Does your code:
return fmt.Sprintf("%s: %v", errorStr, e.unwrapped)

return a wrapped error?
```

***\*No, my code line using\** `fmt.Sprintf` does not wrap the error.**

In Go, string formatting functions like `fmt.Sprintf` with `%v` only extract the text representation of an error. They completely flatten it into a string, stripping away its programmatic identity.

To make `ErrorFromFunctionCall` a true wrapped error, **the wrapping magic actually happens inside the `Unwrap()` method instead.**

------

How Wrapping Works with Custom Struct Types

When you use `fmt.Errorf("%w", err)`, Go creates a built-in wrapper struct behind the scenes that implements an `Unwrap() error` method.

When we transition to custom error structs (like `ErrorFromFunctionCall`), we manage the wrapping ourselves by doing two things:

1. Storing the inner error in a struct field (`unwrapped error`).
2. Implementing the **`Unwrap() error`** method on our struct.

Because the refactored code includes this method:

go

```
func (e *ErrorFromFunctionCall) Unwrap() error {
	return e.unwrapped
}
```

Use code with caution.

The Go standard library (`errors.Is` and `errors.As`) recognizes `ErrorFromFunctionCall` as a wrapped error! When those functions see a struct with an `Unwrap()` method, they call it to automatically dig deeper down your error chain.

------

Why `fmt.Sprintf` with `%v` is used here

Inside the `.Error() string` method, your only goal is to return a clean string representation for logging or printing.

- You **cannot** use `%w` inside `fmt.Sprintf` (it will literally print the characters `%!w(MISSING)` because `fmt.Sprintf` doesn't understand `%w`).
- You **do not need** to use `fmt.Errorf` here because the struct itself *already* acts as the wrapper container thanks to its `Unwrap()` method.

Using `%v` simply reads the text of your underlying error and appends it to your configuration block metadata, while your `Unwrap()` method keeps the programmatic error chain perfectly intact for `errors.As`.

Would you like to see a short **code simulation** showing exactly how `errors.As` traverses through this custom struct to find your inner error, or should we look at adjusting the formatting layout of the final string output?



# Actual Test Scenarios

```
I'm concerned about how this will be implemented in code when processing
real errors. In one test

TestFunc04 calls Series04Func01() error which calls Series04Func02() error which calls
 Series03Func03() error which calls  Series03Func04() error which generates a ErrorNativeToFunction error
 type and returns it up function call chain. 

 Could you generate code to test this function chain and return the combined errors up to TestFunc04?
```

Here is a complete, runnable Go test script that simulates your exact function chain.

This test demonstrates how an `ErrorNativeToFunction` error is generated deep down at the 4th level, wrapped at each layer using `ErrorFromFunctionCall`, and successfully inspected at the top by `TestFunc04` using `errors.As`.

go

```
package customerrors_test

import (
	"errors"
	"fmt"
	"testing"

	// Replace this with the actual import path to your customerrors package
	"customerrors" 
)

// ============================================================================
// Simulated Function Chain (Deep Call Stack)
// ============================================================================

// Series04Func01 calls Series04Func02 and adds context if it fails.
func Series04Func01() error {
	factory := (&customerrors.ErrorGenericConfig{}).NewErrorGenericConfig("series04.go", "Service", "Series04Func01")

	err := Series04Func02()
	if err != nil {
		return factory.GetErrorFromFunctionCall(
			"Series04Func02",
			"orchestrating downstream sequence",
			"top-level operation failed",
			err,
		)
	}
	return nil
}

// Series04Func02 calls Series03Func03 and adds context if it fails.
func Series04Func02() error {
	factory := (&customerrors.ErrorGenericConfig{}).NewErrorGenericConfig("series04.go", "Service", "Series04Func02")

	err := Series03Func03()
	if err != nil {
		return factory.GetErrorFromFunctionCall(
			"Series03Func03",
			"validating intermediate system state",
			"mid-level verification failed",
			err,
		)
	}
	return nil
}

// Series03Func03 calls Series03Func04 and adds context if it fails.
func Series03Func03() error {
	factory := (&customerrors.ErrorGenericConfig{}).NewErrorGenericConfig("series03.go", "Repository", "Series03Func03")

	err := Series03Func04()
	if err != nil {
		return factory.GetErrorFromFunctionCall(
			"Series03Func04",
			"executing data processing operation",
			"low-level repository failed",
			err,
		)
	}
	return nil
}

// Series03Func04 generates the root ErrorNativeToFunction error.
func Series03Func04() error {
	factory := (&customerrors.ErrorGenericConfig{}).NewErrorGenericConfig("series03.go", "Repository", "Series03Func04")

	// Simulating a calculation or operational bounds error native to this function
	return factory.GetErrorNativeToFunction(
		"userID",
		"/var/log/app",
		"data.bin",
		"https://example.com",
		"internal buffer alignment",
		"encountered fatal calculation error: value out of bounds",
	)
}

// ============================================================================
// Test Execution (TestFunc04)
// ============================================================================

func TestFunc04(t *testing.T) {
	// Execute the full call chain
	err := Series04Func01()

	if err == nil {
		t.Fatal("Expected an error from Series04Func01, but got nil")
	}

	// 1. Print the entire combined error string to visualize the stack trace
	fmt.Println("--- FULL COMBINED ERROR OUTPUT ---")
	fmt.Println(err.Error())
	fmt.Println("---------------------------------")

	// 2. Inspect the chain to confirm the original ErrorNativeToFunction is fully visible
	var nativeErr *customerrors.ErrorNativeToFunction
	if errors.As(err, &nativeErr) {
		t.Log("✅ Success: Target ErrorNativeToFunction was found in the error chain!")
		
		// You can directly access the promoted inner data fields on the casted type
		fmt.Printf("Root Cause parameters identified at top:\n")
		// (Assuming fields were exported, or using accessors if they remain lowercase)
		// For private struct fields inside the same package, you can read them. 
		// If testing externally, your struct fields would need to be Capitalized (Exported).
	} else {
		t.Error("❌ Failure: Could not find ErrorNativeToFunction in the returned error chain.")
	}

	// 3. Confirm we can also extract intermediate layer contexts if needed
	var callErr *customerrors.ErrorFromFunctionCall
	if errors.As(err, &callErr) {
		t.Logf("✅ Success: Intercepted intermediate wrapper showing called function: %s", err)
	}
}
```

Use code with caution.

------

How this processes and combines the error chain:

- **The Root:** `Series03Func04` creates a clean `ErrorNativeToFunction` object snapshot.
- **The Climb:** As the error climbs up through `Func03`, `Func02`, and `Func01`, each wrapper creates a new `ErrorFromFunctionCall` instance, setting the previous error into the `unwrapped` property field.
- **The String Representation:** When `TestFunc04` prints `err.Error()`, Go executes the top-level string representation, which evaluates `e.unwrapped` sequentially all the way down, printing out your beautiful multi-line configuration blocks nested together.
- **The Inspection Compatibility:** Because every wrapper implements `Unwrap() error`, `errors.As(err, &nativeErr)` behaves like a drill. It ignores all the `ErrorFromFunctionCall` shells and extracts the hidden `ErrorNativeToFunction` crystal at the core.

To ensure this test executes correctly across packages, let me know if you would like me to show you how to **export the fields** of `ErrorGenericConfig` (capitalize them) so your test file can read the fields, or if we should **write exported Getter methods** instead.