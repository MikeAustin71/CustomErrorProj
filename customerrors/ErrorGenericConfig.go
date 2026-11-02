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
		sourceFile:   sourceFile,
		parentType:   parentType,
		thisFunction: thisFunction,
	}
}

// GetErrorInputParameterOutOfRange
//
// This method is designed to be called when a function's input
// parameter is out of range.
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

	return errors.New(egc.buildErrorStr())
}

// GetErrorInputParameterNilPointer
//
// This method is designed to be called when a function's input
// parameter is identified as a nil pointer.
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

	return errors.New(egc.buildErrorStr())
}

// GetErrorFromFunctionCall
//
// Returns a wrapped error created by adding context to the error
// returned from a called function.
func (egc *ErrorGenericConfig) GetErrorFromFunctionCall(
	calledFunction string,
	errorContext string,
	errorMessage string,
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
