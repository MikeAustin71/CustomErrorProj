package customerrors

import (
	"fmt"
	"strings"
)

// FuncReturnError
// Custom error message used to specify the
// function returning the error message
type FuncReturnError struct {
	ErrPrefix  string
	ReturnFunc string
	ErrContext string
	ErrMessage string
}

// Error
// The following format is used to initialize
// this type of error:
//
// EXAMPLE #1:
//
//	 dec, err := new(Decimal).NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)
//
//		if err != nil {
//
//			return Decimal{},
//				&FuncReturnError{
//					ErrPrefix:  ePrefix.String(),
//					ReturnFunc: "dec, err := new(Decimal).NewBigInt(\n" +
//					"  big.NewInt(0).Set(bNum.bigInt), bNum.precision)\n",
//					ErrContext: "",
//					ErrMessage:   err.Error(),
//				}
//		}
//
//		EXAMPLE #2:
//
//	 "dec, err := new(Decimal).NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)
//
//		if err != nil {
//
//			return Decimal{},
//			&FuncReturnError{
//				ErrPrefix: ePrefix.String(),
//				ReturnFunc: "dec, err := new(Decimal).NewBigInt(\n" +
//				"    big.NewInt(0).Set(bNum.bigInt), bNum.precision)",
//				ErrContext: "",
//				ErrMessage: err.Error(),
//			}
//		}
//
//		NOTE:
//		Element 'ErrContext' is optional
func (e *FuncReturnError) Error() string {

	var errStr string
	foundCnt := 0

	if e.ErrPrefix != "" {
		errStr = e.ErrPrefix + "\n"
		foundCnt++
	}

	if e.ReturnFunc != "" {
		errStr += "Error returned by: \n  " + e.ReturnFunc + "\n"
		foundCnt++
	}

	if e.ErrContext != "" {
		errStr += e.ErrContext + "\n"
		foundCnt++
	}

	if e.ErrMessage != "" {
		errStr += "Error: \n  " + e.ErrMessage + "\n"
		foundCnt++
	}

	if foundCnt == 0 {

		errStr = "No Error parameters provided!\n"

	} else {

		if len(errStr) > 0 {

			errStr = strings.TrimRight(errStr, "\n")

			errStr += "\n"
		}
	}

	return errStr
}

func (e *FuncReturnError) GetError() error {
	//return fmt.Errorf("%w", e.Error())
	return fmt.Errorf("%s", e.Error())
}

func (e *FuncReturnError) Unwrap() error {

	return fmt.Errorf("%s", e.Error())
}

// InputPtrNilError
// Custom error message used to format an error message
// in the case of a nil input parameter pointer
//
// Example Usage:
//
//	if aNumPrecision == nil {
//		return agMean, agMeanPrecision, gValue, gValuePrecision, cycles,
//		&InputPtrNilError{
//		ErrPrefix:     ePrefix.String(),
//		ParameterName: "'aNumPrecision'", // ParameterName is the name of the nil input parameter pointer
//		}
//	}
type InputPtrNilError struct {
	ErrPrefix     string
	ErrContext    string
	ParameterName string
}

func (i *InputPtrNilError) Error() string {

	var errStr string
	foundCnt := 0

	if i.ErrPrefix != "" {
		errStr = i.ErrPrefix + "\n"
		foundCnt++
	}

	if i.ErrContext != "" {
		errStr += i.ErrContext + "\n"
		foundCnt++
	}

	var isParameterName = false

	if len(i.ParameterName) > 0 {
		isParameterName = true
	}

	if foundCnt == 0 {
		errStr = "No Error parameters provided!\n"
	} else {

		var parmName string

		if isParameterName {
			parmName = i.ParameterName
		} else {
			parmName = "unknown parameter"
		}

		errStr += fmt.Sprintf("FATAL ERROR: Input parameter '%v' is a nil pointer.\n",
			parmName)
	}

	return errStr
}

func (i *InputPtrNilError) Unwrap() error {
	return fmt.Errorf("%s", i.Error())
}
