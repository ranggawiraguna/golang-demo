package resources

import (
	"fmt"
	"golang-demo/utils"
	"math/rand"
)

func factorialRecursion(x int) (y int) {
	if x > 0 {
		y = x * factorialRecursion(x-1)
	} else {
		y = 1
	}
	return
}

func FunctionRecursive() {
	utils.LineBreak()
	fmt.Println("Recursion Function")
	fmt.Println("")
	value := rand.Intn(10)
	fmt.Printf("Factorial Value\t\t: %d\n", value)
	fmt.Printf("Factorial Result\t: %d\n", factorialRecursion(value))
	utils.LineBreak()
}
