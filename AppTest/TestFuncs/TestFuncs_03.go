package TestFuncs

import (
	"fmt"

	"github.com/MikeAustin71/CustomErrorProj/customerrors"
)

type TestSeries03Functions struct {
}

func (tf *TestSeries03Functions) TestFunc03() {

	fmt.Printf("\nExecuting TestFunc02() ...\n")

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries03Functions",
		"TestFunc02")

	err := tf.Series03Func01()

	if err != nil {

		combinedWrappedErr := errBaseCfg.GetErrorFromFunctionCall(
			"Series03Func01()",
			"",
			"",
			err)

		fmt.Printf("%s\n", combinedWrappedErr)

		innerMostErr := errBaseCfg.GetInnerMostError(combinedWrappedErr)

		fmt.Println("Inner Most Error:")

		fmt.Printf("%s\n", innerMostErr)

		return
	}

	fmt.Println("TestFunc02() - Series03Func01() completed without error!")

	return
}

func (tf *TestSeries03Functions) Series03Func01() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries03Functions",
		"Series03Func01")

	err := tf.Series03Func02()

	if err != nil {

		return errBaseCfg.GetErrorFromFunctionCall(
			"Series03Func02()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries03Functions) Series03Func02() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries03Functions",
		"Series03Func02")

	err := tf.Series03Func03()

	if err != nil {

		return errBaseCfg.GetErrorFromFunctionCall(
			"Series03Func03()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries03Functions) Series03Func03() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries03Functions",
		"Series03Func03")

	err := tf.Series03Func04()

	if err != nil {

		return errBaseCfg.GetErrorFromFunctionCall(
			"Series03Func04()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries03Functions) Series03Func04() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries03Functions",
		"Series03Func04")

	err := errBaseCfg.GetErrorNativeToFunction(
		"Serious Error Occurred!!",
		"I'm feeling down - and sent this error.")

	return err
}
