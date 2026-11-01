package customerrors

import "fmt"

// FunctionCallError
//
// Represents a structured custom error type for capturing details
// of an error returned by a call to another function.
//
// functionCallChain provides a chain of functions where the error occurred.
// sourceFile specifies the source file where the error originated.
// parentType indicates the type of the parent that invoked the function.
// thisFunction specifies the name of the function that encountered the error.
// calledFunction specifies the name of the function called by this function
// which generated an error.
// errContext gives contextual details relevant to the error.
// errorMessage describes the error in a human-readable format.
// wrappedErr contains the underlying wrapped error, if any.
type FunctionCallError struct {
	functionCallChain string // ePrefix Function Call Chain
	sourceFile        string
	parentType        string // The 'type' in which this method resides
	thisFunction      string
	calledFunction    string
	errorContext      string
	errorMessage      string
	wrappedErr        error // Error returned from called function
}

func (e *FunctionCallError) NewFunctionCallError(
	functionCallChain string,
	sourceFile string,
	parentType string,
	thisFunction string,
	calledFunction string,
	errorContext string,
	errorMessage string,
	calledFunctionWrappedErr error) (*FunctionCallError, error) {

	if calledFunctionWrappedErr == nil {

		parmErr, initErr := new(FunctionNativeError).NewFunctionNativeError(
			"",
			"FunctionCallError.go",
			"FunctionCallError",
			"NewFunctionCallError",
			"",
			"Input Parameter calledFunctionWrappedErr is nil")

		if initErr != nil {
			return nil, initErr
		}

		return nil, parmErr
	}

	err := &FunctionCallError{
		functionCallChain: functionCallChain,
		sourceFile:        sourceFile,
		parentType:        parentType,
		thisFunction:      thisFunction,
		calledFunction:    calledFunction,
		errorContext:      errorContext,
		errorMessage:      errorMessage,
		wrappedErr:        calledFunctionWrappedErr,
	}

	return err, nil
}

func (e *FunctionCallError) buildErrorString() string {

	errString := ""

	if e.wrappedErr != nil {

		errString = "\n"

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

		if e.calledFunction != "" {
			errString += fmt.Sprintf("Called Function: %s\n", e.calledFunction)
		}

		if e.errorContext != "" {
			errString += fmt.Sprintf("Error Context: %s\n", e.errorContext)
		}

		if e.errorMessage != "" {
			errString += fmt.Sprintf("%s\n", e.errorMessage)
		}

		errString += fmt.Sprintf("%v\n", e.wrappedErr)
	}

	return errString
}

// Error
//
// Required by error interface.
func (e *FunctionCallError) Error() string {

	errString := ""

	if e.wrappedErr != nil {

		errString = "\n"

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

		if e.calledFunction != "" {
			errString += fmt.Sprintf("Called Function: %s\n", e.calledFunction)
		}

		if e.errorContext != "" {
			errString += fmt.Sprintf("Error Context: %s\n", e.errorContext)
		}

		if e.errorMessage != "" {
			errString += fmt.Sprintf("%s\n", e.errorMessage)
		}

		errString += fmt.Sprintf("%v\n", e.wrappedErr)
	}

	return errString
}

func (e *FunctionCallError) Unwrap() error {
	return e.wrappedErr
}
