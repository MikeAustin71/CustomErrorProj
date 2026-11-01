package customerrors

import "fmt"

// FunctionCallError
//
// Represents a structured custom error type for capturing details
// of an error returned by a call to another function.
//
// FunctionCallChain provides a chain of functions where the error occurred.
// SourceFile specifies the source file where the error originated.
// ParentType indicates the type of the parent that invoked the function.
// ThisFunction specifies the name of the function that encountered the error.
// CalledFunction specifies the name of the function called by this function
//
//	which generated an error.
//
// ErrContext gives contextual details relevant to the error.
// Message describes the error in a human-readable format.
// WrappedErr contains the underlying wrapped error, if any.
type FunctionCallError struct {
	FunctionCallChain string // ePrefix Function Call Chain
	SourceFile        string
	ParentType        string
	ThisFunction      string
	CalledFunction    string
	ErrContext        string
	Message           string
	WrappedErr        error
}

// Error
//
// Required by error interface.
func (e FunctionCallError) Error() string {

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

		if e.CalledFunction != "" {
			errString += fmt.Sprintf("Called Function: %s\n", e.CalledFunction)
		}

		if e.ErrContext != "" {
			errString += fmt.Sprintf("Error Context: %s\n", e.ErrContext)
		}

		if e.Message != "" {
			errString += fmt.Sprintf("%s\n", e.Message)
		}

		errString += fmt.Sprintf("%v\n", e.WrappedErr)
	}

	return errString
}

func (e FunctionCallError) Unwrap() error {
	return e.WrappedErr
}
