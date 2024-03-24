package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Output() {
	utils.LineBreak()
	fmt.Println("Output Print")
	fmt.Println("")
	fmt.Print("Rangga")
	fmt.Print("Wiraguna")
	fmt.Println("")
	fmt.Print("Rangga\n")
	fmt.Print("Wiraguna\n\n")
	fmt.Print("Rangga", " ", "Wiraguna", " ", 20, " ", 06, " ", 00, "\n")
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Output Println")
	fmt.Println("")
	fmt.Println("Rangga")
	fmt.Println("Wiraguna")
	fmt.Println("Rangga", "Wiraguna", 20, 06, 00)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Output Printf")
	a, b, c := "Rangga", "Wiraguna", 20

	fmt.Printf("%s %s %d \n", a, b, c)
	fmt.Printf("Type Variable a : %T\n", a)
	fmt.Printf("Type Variable b : %T\n", b)
	fmt.Printf("Type Variable c : %T\n", c)

	fmt.Println("")
	utils.LineBreak()
}
