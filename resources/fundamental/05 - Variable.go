package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Variable() {
	utils.LineBreak()
	fmt.Println("String Variable")
	fmt.Println()
	var textFullname string = "Rangga Wiraguna"
	fmt.Println("textFullname\t\t:", textFullname)
	textEmail := "ranggawiragunar@gmail.com"
	fmt.Println("textEmail\t\t:", textEmail)
	fmt.Println("textFullname,textEmail\t:", textFullname, textEmail)
	fullText := textFullname + " (" + textEmail + ")"
	fullTextOther := fmt.Sprintf("%s (%s)", textFullname, textEmail)
	fmt.Println("fullText\t\t:", fullText)
	fmt.Println("fullTextOther\t\t:", fullTextOther)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Integer Variable")
	fmt.Println()
	var numberOne int = 1
	fmt.Println("numberOne\t\t:", numberOne)
	numberTwo := 123
	fmt.Println("numberTwo\t\t:", numberTwo)
	fmt.Println("numberOne,numberTwo\t:", numberOne, numberTwo)
	numberThree := numberOne + numberTwo
	fmt.Println("numberOne+numberTwo\t:", numberThree)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Float Variable")
	fmt.Println()
	var decimalOne float32 = 1.9
	fmt.Println("decimalOne\t\t:", decimalOne)
	decimalTwo := 123.9 //default type is float64
	fmt.Println("decimalTwo\t\t:", decimalTwo)
	fmt.Println("decimalOne,decimalTwo\t:", decimalOne, decimalTwo)
	decimalThree := decimalOne + float32(decimalTwo) //must cast decimalTwo to float32 for calculate
	fmt.Println("decimalOne+decimalTwo\t:", decimalThree)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Boolean Variable")
	fmt.Println()
	var booleanOne bool = true
	fmt.Println("booleanOne\t\t:", booleanOne)
	booleanTwo := false
	fmt.Println("booleanTwo\t\t:", booleanTwo)
	fmt.Println("booleanOne,booleanTwo\t:", booleanOne, booleanTwo)
	booleanThree := booleanOne == booleanTwo
	fmt.Println("booleanOne==booleanTwo\t:", booleanThree)
	booleanFour := booleanOne != booleanTwo
	fmt.Println("booleanOne!=booleanTwo\t:", booleanFour)
	utils.LineBreak()
}
