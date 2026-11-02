[TheOriginalAmarilloMike](./Assets/sittingduck003.png)



# CustomErrorProj: Develop New Custom Error Types


This project includes a number of custom error types for Go programming language.

The purpose of this project is to develop and test new custom error types that can
be used in various Go applications. The custom error types are designed to provide
more detailed and informative error messages, making it easier to debug and
troubleshoot issues in Go code.



## ***ErrorGenericFactory***

Based on testing to date, the ***ErrorGenericFactory*** custom error type has been found to be the most effective in providing detailed and informative error messages for various scenarios. It allows developers to create custom error messages with specific error codes and descriptions, making it easier to identify and resolve issues in their Go code.



### Be Advised: ***ErrorGenericFactory*** is Still Undergoing Testing

[Caution30](./Assets/underconstruction003.png)

Testing and evaluation of ***ErrorGenericFactory*** is still ongoing. Progress updates will follow. In the meantime, this is looking very promising. 



## ***ErrorGenericFactory*** Example Usage

### Call Tree for Test

```go
TestFunc04() // Starting point.TestFunc04() calls Series04Func01() 
            |
            |----> Series04Func01() // Receives and Returns Error
                 |
                 ----> Series04Func02() // Receives and Returns Error
                     |
                     ----> Series04Func03()  // Receives and Returns Error
                         |
                          ----> Series04Func04() // Base error generated here
```



### Test Source Code

```go
import (
	"fmt"

	"github.com/MikeAustin71/CustomErrorProj/customerrors"
)

type TestSeries04Functions struct {
}

func (tf *TestSeries04Functions) TestFunc04() {

	fmt.Printf("\nExecuting TestFunc04() ...\n")

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"TestFunc02")

	err := tf.Series04Func01()

	if err != nil {

		combinedWrappedErr := errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func01()",
			"",
			"",
			err)

		fmt.Printf("%s\n", combinedWrappedErr)

		innerMostErr := errGenericCfg.GetInnerMostError(combinedWrappedErr)

		fmt.Println("Inner Most Error:")

		fmt.Printf("%s\n", innerMostErr)

		return
	}

	fmt.Println("TestFunc02() - Series04Func01() completed without error!")

	return
}

func (tf *TestSeries04Functions) Series04Func01() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func01")

	err := tf.Series04Func02()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func02()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries04Functions) Series04Func02() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func02")

	err := tf.Series04Func03()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func03()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries04Functions) Series04Func03() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func03")

	err := tf.Series04Func04()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func04()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries04Functions) Series04Func04() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func04")

	err := errGenericCfg.GetErrorNativeToFunction(
		"",
		"",
		"",
		"",
		"Serious Error Occurred!!",
		"I'm feeling down - and sent this error.")

	return err
}

```



## ***ErrorGenericFactory*** Source Code

```go
// Source Code File: ErrorGenericFactory_0010.go

package customerrors

import (
	"errors"
	"fmt"
)

// ============================================================================
// 1. Core Configuration Types
// ============================================================================

type ErrorGenericFactory struct {
	sourceFile         string // Source file containing 'thisFunction'.
	thisFunctionType   string // The 'type' in which 'thisFunction' (method) resides.
	thisFunction       string // The current function processing this error.
	calledFunctionType string // The 'type' of the method identified in
	// 'calledFunction', below.
	calledFunction string // The method or function called by 'thisFunction'.
	parameter      string // Parameter associated with this error.
	path           string // File path associated with this error.
	fileName       string // File name associated with this error.
	url            string // URL associated with this error.
	errorContext   string // Error context associated with this error.
	errorMessage   string // The error message associated with this error.
}

// ============================================================================
// 2. Specialized Struct-Based Error Types
// ============================================================================

// ErrorInputParameterOutOfRange represents a parameter value out of bounds.
type ErrorInputParameterOutOfRange struct {
	ErrorGenericFactory // Value embedding ensures an isolated, safe snapshot
}

func (e *ErrorInputParameterOutOfRange) Error() string {
	return e.buildErrorStr()
}

// ErrorInputParameterNilPointer represents a required pointer passing a nil value.
type ErrorInputParameterNilPointer struct {
	ErrorGenericFactory
}

func (e *ErrorInputParameterNilPointer) Error() string {
	return e.buildErrorStr()
}

// ErrorNativeToFunction represents an internal operational failure within a function.
type ErrorNativeToFunction struct {
	ErrorGenericFactory
}

func (e *ErrorNativeToFunction) Error() string {
	return e.buildErrorStr()
}

// ErrorFromFunctionCall represents an error context layer wrapping a downstream failure.
type ErrorFromFunctionCall struct {
	ErrorGenericFactory
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
func (egc *ErrorGenericFactory) NewErrorGenericConfig(
	sourceFile string,
	parentType string,
	thisFunction string) *ErrorGenericFactory {

	return &ErrorGenericFactory{
		sourceFile:       sourceFile,
		thisFunctionType: parentType,
		thisFunction:     thisFunction,
	}
}

// GetErrorInputParameterOutOfRange yields a unique ErrorInputParameterOutOfRange error.
func (egc *ErrorGenericFactory) GetErrorInputParameterOutOfRange(
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
	return &ErrorInputParameterOutOfRange{ErrorGenericFactory: *egc}
}

// GetErrorInputParameterNilPointer yields a unique ErrorInputParameterNilPointer error.
func (egc *ErrorGenericFactory) GetErrorInputParameterNilPointer(
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

	return &ErrorInputParameterNilPointer{ErrorGenericFactory: *egc}
}

// GetErrorNativeToFunction yields a unique ErrorNativeToFunction error.
func (egc *ErrorGenericFactory) GetErrorNativeToFunction(
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

	return &ErrorNativeToFunction{ErrorGenericFactory: *egc}
}

// GetErrorFromFunctionCall returns an inspectable, wrapped FunctionCallError.
func (egc *ErrorGenericFactory) GetErrorFromFunctionCall(
	calledFunctionType string,
	calledFunction string,
	errorContext string,
	errorMessage string,
	returnedCalledFunctionError error) error {

	egc.calledFunctionType = calledFunctionType
	egc.calledFunction = calledFunction
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	return &ErrorFromFunctionCall{
		ErrorGenericFactory: *egc,
		unwrapped:           returnedCalledFunctionError,
	}
}

// GetInnerMostError
//
// Returns the innermost (or last) error from a series of
// wrapped errors.
func (egc *ErrorGenericFactory) GetInnerMostError(
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
func (egc *ErrorGenericFactory) SetDetailErrorContext(
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
func (egc *ErrorGenericFactory) SetErrorMessageElements(
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
func (egc *ErrorGenericFactory) buildErrorStr() string {
	errStr := ""
	foundParametersCount := 0

	if egc.sourceFile != "" {
		errStr = "\n"
		errStr += fmt.Sprintf("Source File: %s\n", egc.sourceFile)
		foundParametersCount++
	}

	if egc.thisFunctionType != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("This Function Type: %s\n", egc.thisFunctionType)
		foundParametersCount++
	}

	if egc.thisFunction != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("This Function: %s\n", egc.thisFunction)
		foundParametersCount++
	}

	if egc.calledFunctionType != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("Called Function Type: %s\n", egc.calledFunctionType)
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





## CustomErrorProj: Operational Configuration

**CustomErrorProj** was compiled and tested with Go Version 1.27.3 on Windows 11. 
