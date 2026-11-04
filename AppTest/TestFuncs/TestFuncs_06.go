package TestFuncs

import (
	"fmt"

	"github.com/MikeAustin71/CustomErrorProj/customerrors"
	ePref "github.com/MikeAustin71/errpref/v2"
)

type TestSeries06Functions struct {
}

func (tf *TestSeries06Functions) TestFunc06() {

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_06.go",
		"TestSeries06Functions",
		"TestFunc06")

	returnedErr := tf.Series06Func01()

	if returnedErr != nil {

		combinedWrappedErr := errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries06Functions",
			"Series06Func01()",
			"",
			"",
			returnedErr)

		if combinedWrappedErr == nil {

			fmt.Printf("\nTest FAILED!\n")
			fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n")
			fmt.Printf("\ncombinedWrappedErr is nil!\n")
			fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n")

			return
		}

		fmt.Printf("\nTest Succeeded!\n")
		fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
		fmt.Printf("\ncombinedWrappedErr:\n%+v\n", combinedWrappedErr)

	} else {
		fmt.Printf("\nTest FAILED!\n")
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n")
		fmt.Printf("returnedErr==nil\n")
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n")

	}

	return
}

func (tf *TestSeries06Functions) Series06Func01() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_06.go",
		"TestSeries06Functions",
		"Series06Func01")

	err := tf.Series06Func02()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries06Functions",
			"Series06Func02()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries06Functions) Series06Func02() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_06.go",
		"TestSeries06Functions",
		"Series06Func02")

	err := tf.Series06Func03()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries06Functions",
			"Series06Func03()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries06Functions) Series06Func03() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_06.go",
		"TestSeries06Functions",
		"Series06Func03")

	err := tf.Series06Func04()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries06Functions",
			"Series06Func04()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries06Functions) Series06Func04() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_06.go",
		"TestSeries06Functions",
		"Series06Func04")

	err := tf.Series06Func05()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries06Functions",
			"Series06Func05()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries06Functions) Series06Func05() error {

	errGenericCfg := new(customerrors.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_06.go",
		"TestSeries06Functions",
		"Series06Func05")

	err := errGenericCfg.GetErrorNativeToFunction(
		"somethingBroke",
		"",
		"",
		"",
		"Serious Error Occurred!!",
		"I'm feeling down - and sent this error.")

	return err
}
