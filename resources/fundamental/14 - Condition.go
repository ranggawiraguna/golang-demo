package resources

import (
	"fmt"
	"golang-demo/utils"
	"slices"
	"strings"
	"time"
)

func Condition() {
	utils.LineBreak()
	fmt.Println("Basic Condition")
	fmt.Println("")
	hour := time.Now().Hour()
	if hour >= 5 && hour < 10 {
		fmt.Println("Hi, Good Morning!")
	} else if hour >= 10 && hour < 17 {
		fmt.Println("Hi, Good Afternoon!")
	} else if hour >= 17 && hour < 19 {
		fmt.Println("Hi, Good Evening!")
	} else {
		fmt.Println("Hi, Good Night!")
	}
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Nested Condition")
	fmt.Println("")
	day := time.Now().Weekday()
	if !slices.Contains([]string{"saturday", "sunday"}, strings.ToLower(day.String())) {
		hour := time.Now().Hour()
		if hour >= 8 && hour <= 17 {
			fmt.Println("Now is working time")
		} else {
			fmt.Println("Now is not working time")
		}
	} else {
		fmt.Println("Now is not working time")
	}
	utils.LineBreak()
}
