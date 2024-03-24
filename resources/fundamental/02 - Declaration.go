package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Declaration() {
	utils.LineBreak()
	fmt.Println("Multiple Declartion with Same Type")
	fmt.Println("")
	var a, b, c int = 10, 20, 30
	fmt.Println("Variable a : ", a)
	fmt.Println("Variable b : ", b)
	fmt.Println("Variable c : ", c)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Multiple Declartion with Different Type")
	fmt.Println("")
	var d, e, f = 10, "Rangga", true
	g, h := 20, "Wiraguna"
	fmt.Println("Variable d : ", d)
	fmt.Println("Variable e : ", e)
	fmt.Println("Variable f : ", f)
	fmt.Println("Variable g : ", g)
	fmt.Println("Variable h : ", h)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Multiple Declartion in Block")
	fmt.Println("")
	var (
		x bool
		y int    = 20
		z string = "Rangga"
	)

	fmt.Println("Variable x : ", x)
	fmt.Println("Variable y : ", y)
	fmt.Println("Variable z : ", z)
	utils.LineBreak()
}
