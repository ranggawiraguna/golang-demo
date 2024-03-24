package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Comparison() {
	utils.LineBreak()
	fmt.Println("Comparison")
	fmt.Println("")
	a, b, c, d, e := 11, 32, 54, 11, 32
	fmt.Printf("%v == %v\t:%v \n", a, d, a == d)
	fmt.Printf("%v != %v\t:%v \n", a, b, a != b)
	fmt.Printf("%v < %v\t\t:%v \n", b, c, b < c)
	fmt.Printf("%v > %v\t\t:%v \n", c, d, c > d)
	fmt.Printf("%v <= %v\t:%v \n", d, e, d <= e)
	fmt.Printf("%v >= %v\t:%v \n", e, a, e >= a)
	utils.LineBreak()
}
