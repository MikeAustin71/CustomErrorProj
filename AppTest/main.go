package main

import (
	"fmt"

	"github.com/MikeAustin71/CustomErrorProj/AppTest/TestFuncs"
)

func main() {

	fmt.Printf("Running TestFunc_01()...\n\n")
	tf := new(TestFuncs.TestFunctions)

	tf.Testfunc01()
}
