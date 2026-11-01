package TestFuncs

import (
	"fmt"

	"github.com/MikeAustin71/CustomErrorProj/customerrors"
)

type TestFunctions struct{}

func (tf *TestFunctions) Testfunc01() {

	err := Series01Func01()

	if err != nil {

		fmt.Printf("%s\n", err)

		return
	}

	fmt.Println("TestFunc_01() - Series01Func01() completed without error!")

	return
}

func Series01Func01() error {

	callErr := Series01Func02()

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

func Series01Func02() error {

	callErr := Series01Func03()

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

func Series01Func03() error {

	callErr := Series01Func04()

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

func Series01Func04() error {

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
