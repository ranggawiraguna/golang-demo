package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Formatting() {
	utils.LineBreak()
	fmt.Println("General Formatting")
	fmt.Println("")
	generalNumber := 15.5
	generalText := "Hello World!"
	fmt.Printf("%v \n", generalNumber)   //Prints the value in the default format
	fmt.Printf("%#v \n", generalNumber)  //Prints the value in Go-syntax format
	fmt.Printf("%v%% \n", generalNumber) //Prints the value in the default format with % sign
	fmt.Printf("%T \n", generalNumber)   //Prints the type of the value
	fmt.Printf("%v \n", generalText)     //Prints the value in the default format
	fmt.Printf("%#v \n", generalText)    //Prints the value in Go-syntax format
	fmt.Printf("%T \n", generalText)     //Prints the type of the value
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Integer Formatting")
	fmt.Println("")
	integerNumber := 15
	fmt.Printf("%b \n", integerNumber)
	fmt.Printf("%d \n", integerNumber)
	fmt.Printf("%+d \n", integerNumber)
	fmt.Printf("%o \n", integerNumber)
	fmt.Printf("%O \n", integerNumber)
	fmt.Printf("%x \n", integerNumber)
	fmt.Printf("%X \n", integerNumber)
	fmt.Printf("%#x \n", integerNumber)
	fmt.Printf("%4d \n", integerNumber)
	fmt.Printf("%-4d \n", integerNumber)
	fmt.Printf("%04d \n", integerNumber)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("String Formatting")
	fmt.Println("")
	stringText := "Hello"
	fmt.Printf("%s \n", stringText)
	fmt.Printf("%q \n", stringText)
	fmt.Printf("%8s \n", stringText)
	fmt.Printf("%-8s \n", stringText)
	fmt.Printf("%x \n", stringText)
	fmt.Printf("% x \n", stringText)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Boolean Formatting")
	fmt.Println("")
	fmt.Printf("%t \n", true)
	fmt.Printf("%t \n", false)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Float Formatting")
	fmt.Println("")
	floatNumber := 3.141
	fmt.Printf("%e\n", floatNumber)
	fmt.Printf("%f\n", floatNumber)
	fmt.Printf("%.2f\n", floatNumber)
	fmt.Printf("%6.2f\n", floatNumber)
	fmt.Printf("%g\n", floatNumber)
	utils.LineBreak()
}
