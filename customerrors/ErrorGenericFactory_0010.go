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
	parentType     string // The 'type' in which thisFunction resides
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

// GetInnerMostError
//
// Returns the innermost error from a wrapped error.
func (egc *ErrorGenericConfig) GetInnerMostError(
	wrappedErrors error) error {

	if wrappedErrors == nil {
		return errors.New("ErrorGenericConfig.GetInnerMostError()\n" +
			"Input parameter 'wrappedErrors' is nil")
	}

	innerMostError := wrappedErrors

	for {

		unwrappedErr := errors.Unwrap(innerMostError)

		if unwrappedErr == nil {
			break
		}

		innerMostError = unwrappedErr
	}

	return innerMostError
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
