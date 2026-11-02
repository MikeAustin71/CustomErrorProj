package customerrors

import (
	"errors"
	"fmt"
)

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

func (ebc *ErrorBaseConfig) GetErrorFromFunctionCall(calledFunction string,
	errorContext string,
	errorMessage string,
	functionReturnErr error) error {

	errFuncCall := &ErrorFunctionCall{
		SourceFile:     ebc.sourceFile,
		ParentType:     ebc.parentType,
		ThisFunction:   ebc.thisFunction,
		CalledFunction: calledFunction,
		ErrorContext:   errorContext,
		ErrorMessage:   errorMessage,
	}

	if functionReturnErr == nil {
		return fmt.Errorf("%s", errFuncCall.Error())
	}

	return fmt.Errorf("%s: %w", errFuncCall.Error(), functionReturnErr)
}

func (ebc *ErrorBaseConfig) GetErrorFunctionCall(
	calledFunction string,
	errorContext string,
	errorMessage string) *ErrorFunctionCall {

	return &ErrorFunctionCall{
		SourceFile:     ebc.sourceFile,
		ParentType:     ebc.parentType,
		ThisFunction:   ebc.thisFunction,
		CalledFunction: calledFunction,
		ErrorContext:   errorContext,
		ErrorMessage:   errorMessage,
	}
}

func (ebc *ErrorBaseConfig) GetSimpleCombinedTwoErrors(
	errString string,
	err2 error) error {

	return fmt.Errorf("%s: %s", errString, err2)
}

func (ebc *ErrorBaseConfig) GetWrappedCombinedTwoErrors(
	errString string,
	err2 error) error {

	return fmt.Errorf("%s: %w", errString, err2)
}

func (ebc *ErrorBaseConfig) GetInnerMostError(
	wrappedErrors error) error {

	if wrappedErrors == nil {
		return errors.New("ErrorBaseConfig.GetInnerMostError()\n" +
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
