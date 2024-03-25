package resources

import (
	"fmt"
	"golang-demo/utils"
	"math/rand"
	"time"
)

func printSayGreeting() {
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
}

func printScoreGrade(name string, score int) {
	fmt.Printf("Name\t: %s\n", name)
	fmt.Printf("Value\t: %d\n", score)
	if score >= 80 && score <= 100 {
		fmt.Println("Grade\t: A")
	} else if score >= 68 && score <= 79 {
		fmt.Println("Grade\t: B")
	} else if score >= 50 && score <= 67 {
		fmt.Println("Grade\t: C")
	} else if score >= 40 && score <= 49 {
		fmt.Println("Grade\t: D")
	} else {
		fmt.Println("Grade\t: E")
	}
}

func Function() {
	utils.LineBreak()
	fmt.Println("Basic Function")
	fmt.Println("")
	printSayGreeting()
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Parameter & Argument")
	fmt.Println("")
	name := "Rangga"
	score := rand.Intn(101)
	printScoreGrade(name, score)
	utils.LineBreak()
}
