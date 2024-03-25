package resources

import (
	"fmt"
	"golang-demo/utils"
	"math/rand"
	"strings"
	"time"
)

func Switch() {
	utils.LineBreak()
	fmt.Println("Single Case")
	fmt.Println("")
	number := rand.Intn(3) + 1
	switch number {
	case 1:
		fmt.Printf("Your Number is %d > Grade A\n", number)
	case 2:
		fmt.Printf("Your Number is %d > Grade B\n", number)
	case 3:
		fmt.Printf("Your Number is %d > Grade C\n", number)
	default:
		fmt.Println("Wrong Number")
	}
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Multi Case")
	fmt.Println("")
	day := time.Now().Weekday()
	switch strings.ToLower(day.String()) {
	case "saturday", "sunday":
		fmt.Println("This is a weekday!")
	default:
		fmt.Println("This is not a weekday!")
	}
	utils.LineBreak()
}
