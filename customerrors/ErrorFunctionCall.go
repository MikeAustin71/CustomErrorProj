package customerrors

import (
	"errors"
	"fmt"
)

type ErrorFunctionCall struct {
	SourceFile     string
	ParentType     string // The 'type' in which this method resides
	ThisFunction   string
	CalledFunction string
	ErrorContext   string
	ErrorMessage   string
}

func (e *ErrorFunctionCall) Error() string {

	errStr := e.buildErrorStr()

	return errStr
}

func (e *ErrorFunctionCall) Unwrap() error {

	err := errors.New(e.buildErrorStr())

	return fmt.Errorf("%w", err)

}

func (e *ErrorFunctionCall) buildErrorStr() string {

	errStr := ""

	foundParametersCount := 0

	if e.SourceFile != "" {
		errStr = "\n"
		errStr += fmt.Sprintf("SourceFile: %s\n", e.SourceFile)
		foundParametersCount++
	}

	if e.ParentType != "" {

		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("ParentType: %s\n", e.ParentType)

		foundParametersCount++
	}

	if e.ThisFunction != "" {

		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("ThisFunction: %s\n", e.ThisFunction)

		foundParametersCount++
	}

	if e.CalledFunction != "" {

		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("CalledFunction: %s\n", e.CalledFunction)

		foundParametersCount++
	}

	if e.ErrorContext != "" {

		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("ErrorContext: %s\n", e.ErrorContext)

		foundParametersCount++
	}

	if e.ErrorMessage != "" {

		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("ErrorMessage: %s\n", e.ErrorMessage)

		foundParametersCount++
	}

	if foundParametersCount == 0 {
		errStr += "\nNo parameters found for this ErrorFunctionCall instance\n"
	}

	return errStr
}
