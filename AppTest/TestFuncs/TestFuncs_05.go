package TestFuncs

import (
	"fmt"

	ePref "github.com/MikeAustin71/errpref/v2"
)

type TestSeries05Functions struct {
}

func (tf *TestSeries05Functions) TestFunc05() {

	fmt.Printf("\nExecuting TestFunc04() ...\n")

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_05.go",
		"TestSeries05Functions",
		"TestFunc05()")

	err := tf.Series05Func01()

	if err != nil {

		combinedWrappedErr := errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series05Func01()",
			"",
			"",
			err)

		fmt.Printf("%s\n", combinedWrappedErr)

		innerMostErr := errGenericCfg.GetInnerMostError(combinedWrappedErr)

		fmt.Println("Inner Most Error:")

		fmt.Printf("%s\n", innerMostErr)

		return
	}

	fmt.Printf("\n&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n")
	fmt.Println("TestFunc04() completed without error!")
	fmt.Println("That's BAD!. TestFunc05 FAILED!!")
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	return
}

func (tf *TestSeries05Functions) Series05Func01() error {

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_05.go",
		"TestSeries05Functions",
		"Series05Func01")

	err := tf.Series05Func02()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries05Functions",
			"Series05Func02()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries05Functions) Series05Func02() error {

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_05.go",
		"TestSeries05Functions",
		"Series05Func02")

	err := tf.Series05Func03()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series05Func03()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries05Functions) Series05Func03() error {

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_05.go",
		"TestSeries05Functions",
		"Series05Func03")

	err := tf.Series05Func04()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries05Functions",
			"Series05Func04()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries05Functions) Series05Func04() error {

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_05.go",
		"TestSeries05Functions",
		"Series05Func04")

	err := tf.Series05Func05()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries05Functions",
			"Series05Func05()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries05Functions) Series05Func05() error {

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_05.go",
		"TestSeries05Functions",
		"Series05Func05")

	err := tf.Series05Func06()

	if err != nil {

		return errGenericCfg.GetErrorFromFunctionCall(
			"TestSeries05Functions",
			"Series05Func06()",
			"Some Context From Series05Func05()",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries05Functions) Series05Func06() error {

	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_05.go",
		"TestSeries05Functions",
		"Series05Func06")

	err := errGenericCfg.GetErrorNativeToFunction(
		"Some really Funky Parameter",
		"",
		"",
		"",
		"Serious Error Occurred!!",
		"I'm feeling down - and sent this error.")

	return err
}
