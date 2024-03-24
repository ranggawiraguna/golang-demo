package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Logical() {
	utils.LineBreak()
	fmt.Println("Comparison")
	fmt.Println("")
	a, b, c, d, e := 11, 32, 54, 11, 32
	fmt.Printf("%v == %v && %v != %v\t: %v \n", a, d, a, b, a == d && a != b)
	fmt.Printf("%v < %v || %v > %v\t: %v \n", b, c, c, d, b < c || c > d)
	fmt.Printf("%v <= %v\t\t: not %v \n", d, e, !(d <= e))
	utils.LineBreak()
}
