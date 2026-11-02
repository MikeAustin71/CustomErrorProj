package customerrors

import "fmt"

type ErrorBaseConfig struct {
	sourceFile     string
	parentType     string // The 'type' in which this method resides
	thisFunction   string
	calledFunction string
	errorContext   string
	errorMessage   string
}

// NewErrorBaseConfig
//
// Designed to be called at beginning of a function when
// only three elements are known: sourceFile, parentType, and thisFunction.
func (ebc *ErrorBaseConfig) NewErrorBaseConfig(
	sourceFile string,
	parentType string,
	thisFunction string) *ErrorBaseConfig {

	errBaseConfig := &ErrorBaseConfig{
		sourceFile:   sourceFile,
		parentType:   parentType,
		thisFunction: thisFunction,
	}

	return errBaseConfig
}

func (ebc *ErrorBaseConfig) GetErrorNativeToFunction(
	errorContext string,
	errorMessage string) *ErrorNativeToFunction {

	return &ErrorNativeToFunction{
		SourceFile:   ebc.sourceFile,
		ParentType:   ebc.parentType,
		ThisFunction: ebc.thisFunction,
		ErrorContext: errorContext,
		ErrorMessage: errorMessage,
	}
}

func (ebc *ErrorBaseConfig) GetErrorFunctionCall(
	calledFunction string,
	errorContext string,
	errorMessage string) ErrorFunctionCall {

	return ErrorFunctionCall{
		SourceFile:     ebc.sourceFile,
		ParentType:     ebc.parentType,
		ThisFunction:   ebc.thisFunction,
		CalledFunction: calledFunction,
		ErrorContext:   errorContext,
		ErrorMessage:   errorMessage,
	}
}

func (ebc *ErrorBaseConfig) GetCombinedTwoErrors(
	errString string,
	err2 error) error {

	return fmt.Errorf("%s: %s", errString, err2)
}

func (ebc *ErrorBaseConfig) GetWrappedCombinedTwoErrors(
	errString string,
	err2 error) error {

	return fmt.Errorf("%s: %w", errString, err2)
}
