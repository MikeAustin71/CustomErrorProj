package TestFuncs

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
