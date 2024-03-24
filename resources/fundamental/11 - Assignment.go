package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Assignment() {
	utils.LineBreak()
	fmt.Println("Assignment")
	fmt.Println("")
	value := 0
	fmt.Println("value\t\t:", value)
	value += 9
	fmt.Println("value += 9\t:", value)
	value -= 4
	fmt.Println("value -= 4\t:", value)
	value *= 6
	fmt.Println("value *= 6\t:", value)
	value /= 3
	fmt.Println("value /= 3\t:", value)
	value %= 6
	fmt.Println("value %= 6\t:", value)
	value &= 1
	fmt.Println("value &= 1\t:", value)
	value |= 2
	fmt.Println("value|= 2\t:", value)
	value ^= 10
	fmt.Println("value ^= 10\t:", value)
	value >>= 1
	fmt.Println("value >>= 1\t:", value)
	value <<= 2
	fmt.Println("value <<= 2\t:", value)
	utils.LineBreak()
}
