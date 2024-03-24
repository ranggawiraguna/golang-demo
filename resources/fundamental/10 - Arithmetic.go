package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Arithmetic() {
	utils.LineBreak()
	fmt.Println("Arithmetic")
	fmt.Println("")
	numberA, numberB := 13, 59

	fmt.Printf("%v + %v\t= %v\n", numberA, numberB, numberA+numberB)
	fmt.Printf("%v - %v\t= %v\n", numberA, numberB, numberA-numberB)
	fmt.Printf("%v * %v\t= %v\n", numberA, numberB, numberA*numberB)
	fmt.Printf("%v / %v\t= %v\n", numberA, numberB, numberA/numberB)
	fmt.Printf("%v %% %v\t= %v\n", numberA, numberB, numberA%numberB)
	numberA2 := numberA
	numberA2++
	fmt.Printf("%v++\t= %v\n", numberA, numberA2)
	numberB2 := numberB
	numberB2--
	fmt.Printf("%v--\t= %v\n", numberB, numberB2)

	utils.LineBreak()
}
