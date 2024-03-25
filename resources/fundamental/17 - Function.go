package resources

import (
	"fmt"
	"golang-demo/utils"
	"math/rand"
	"time"
)

func sayGreeting() {
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

func getScoreGrade(name string, score int) {
	fmt.Printf("Your Name\t: %s\n", name)
	fmt.Printf("Your Value\t: %d\n", score)
	if score >= 80 && score <= 100 {
		fmt.Println("Your Grade\t: A")
	} else if score >= 68 && score <= 79 {
		fmt.Println("Your Grade\t: B")
	} else if score >= 50 && score <= 67 {
		fmt.Println("Your Grade\t: C")
	} else if score >= 40 && score <= 49 {
		fmt.Println("Your Grade\t: D")
	} else {
		fmt.Println("Your Grade\t: E")
	}
}

func Function() {
	utils.LineBreak()
	fmt.Println("Basic Function")
	fmt.Println("")
	sayGreeting()
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Parameter & Argument")
	fmt.Println("")
	name := "Rangga"
	score := rand.Intn(101)
	getScoreGrade(name, score)
	utils.LineBreak()
}
