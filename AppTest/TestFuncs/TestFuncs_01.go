package TestFuncs

import (
	"fmt"

	"github.com/MikeAustin71/CustomErrorProj/customerrors"
)

type TestSeries01Functions struct{}

func (tf *TestSeries01Functions) TestFunc01() {

	fmt.Printf("\nExecuting TestFunc01() ...\n")

	err := tf.Series01Func01()

	if err != nil {

		fmt.Printf("%s\n", err)

		return
	}

	fmt.Println("TestFunc_01() - Series01Func01() completed without error!")

	return
}

func (tf *TestSeries01Functions) Series01Func01() error {

	callErr := tf.Series01Func02()

	if callErr != nil {

		err, initErr := new(customerrors.FunctionCallError).NewFunctionCallError(
			"",
			"TestFuncs_01.go",
			"",
			"Series01Func01()",
			"Series01Func02()",
			"",
			"",
			callErr)

		if initErr != nil {
			return initErr
		}

		return err
	}

	return nil
}

func (tf *TestSeries01Functions) Series01Func02() error {

	callErr := tf.Series01Func03()

	if callErr != nil {

		err, initErr := new(customerrors.FunctionCallError).NewFunctionCallError(
			"",
			"TestFuncs_01.go",
			"",
			"Series01Func02()",
			"Series01Func03()",
			"",
			"",
			callErr)

		if initErr != nil {
			return initErr
		}

		return err
	}

	return nil
}

func (tf *TestSeries01Functions) Series01Func03() error {

	callErr := tf.Series01Func04()

	if callErr != nil {

		err, initErr := new(customerrors.FunctionCallError).NewFunctionCallError(
			"",
			"TestFuncs_01.go",
			"",
			"Series01Func03()",
			"Series01Func04()",
			"",
			"",
			callErr)

		if initErr != nil {
			return initErr
		}

		return err
	}

	return nil
}

func (tf *TestSeries01Functions) Series01Func04() error {

	err, initErr := new(customerrors.FunctionNativeError).NewFunctionNativeError(
		"",
		"TestFuncs01.go",
		"",
		"Series01Func04()",
		"Inner Error",
		"Returned an error 'Just Because'")

	if initErr != nil {
		return initErr
	}

	return err
}
