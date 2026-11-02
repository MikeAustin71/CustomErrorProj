package TestFuncs

import (
	"fmt"

	"github.com/MikeAustin71/CustomErrorProj/customerrors"
)

type TestSeries02Functions struct {
}

func (tf *TestSeries02Functions) TestFunc02() {

	fmt.Printf("\nExecuting TestFunc02() ...\n")

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries02Functions",
		"TestFunc02")

	err := tf.Series02Func01()

	if err != nil {

		errFuncCall := errBaseCfg.GetErrorFunctionCall(
			"Series02Func01()",
			"",
			"")

		combinedErr := errBaseCfg.GetWrappedCombinedTwoErrors(errFuncCall.Error(), err)

		fmt.Printf("%s\n", combinedErr)

		return
	}

	fmt.Println("TestFunc02() - Series02Func01() completed without error!")

	return
}

func (tf *TestSeries02Functions) Series02Func01() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries02Functions",
		"Series02Func01")

	err := tf.Series02Func02()

	if err != nil {

		errFuncCall := errBaseCfg.GetErrorFunctionCall(
			"Series02Func02()",
			"",
			"")

		return errBaseCfg.GetWrappedCombinedTwoErrors(errFuncCall.Error(), err)
	}

	return nil
}

func (tf *TestSeries02Functions) Series02Func02() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries02Functions",
		"Series02Func02")

	err := tf.Series02Func03()

	if err != nil {

		errFuncCall := errBaseCfg.GetErrorFunctionCall(
			"Series02Func03()",
			"",
			"")

		return errBaseCfg.GetWrappedCombinedTwoErrors(errFuncCall.Error(), err)
	}

	return nil
}

func (tf *TestSeries02Functions) Series02Func03() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries02Functions",
		"Series02Func03")

	err := tf.Series02Func04()

	if err != nil {

		errFuncCall := errBaseCfg.GetErrorFunctionCall(
			"Series02Func04()",
			"",
			"")

		return errBaseCfg.GetWrappedCombinedTwoErrors(errFuncCall.Error(), err)
	}

	return nil
}

func (tf *TestSeries02Functions) Series02Func04() error {

	errBaseCfg := new(customerrors.ErrorBaseConfig).NewErrorBaseConfig(
		"TestFuncs_02.go",
		"TestSeries02Functions",
		"Series02Func04")

	err := errBaseCfg.GetErrorNativeToFunction(
		"Serious Error Occurred!!",
		"I'm feeling down - and sent this error.")

	return err
}
