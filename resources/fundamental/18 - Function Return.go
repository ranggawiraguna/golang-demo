package resources

import (
	"fmt"
	"golang-demo/utils"
	"math/rand"
)

func getScoreGrade(score int) (result string) {
	if score >= 80 && score <= 100 {
		return "Grade\t: A"
	} else if score >= 68 && score <= 79 {
		return "Grade\t: B"
	} else if score >= 50 && score <= 67 {
		return "Grade\t: C"
	} else if score >= 40 && score <= 49 {
		return "Grade\t: D"
	} else {
		return "Grade\t: E"
	}
}

func getScoreGradeValue(score int) (result string) {
	if score >= 80 && score <= 100 {
		result = "A"
	} else if score >= 68 && score <= 79 {
		result = "B"
	} else if score >= 50 && score <= 67 {
		result = "C"
	} else if score >= 40 && score <= 49 {
		result = "D"
	} else {
		result = "E"
	}
	return
}

func getScoreGradeValueWithStatus(score int) (result string, status string) {
	if score >= 80 && score <= 100 {
		result = "A"
		status = "Passed"
	} else if score >= 68 && score <= 79 {
		result = "B"
		status = "Passed"
	} else if score >= 50 && score <= 67 {
		result = "C"
		status = "Not Pass"
	} else if score >= 40 && score <= 49 {
		result = "D"
		status = "Not Pass"
	} else {
		result = "E"
		status = "Not Pass"
	}
	return
}

func FunctionReturn() {
	var score int
	var grade, status string

	utils.LineBreak()
	fmt.Println("Basic Function Return")
	fmt.Println("")
	score = rand.Intn(101)
	fmt.Printf("Value\t: %d\n", score)
	fmt.Println(getScoreGrade(score))
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Store Return Value at Variable")
	fmt.Println("")
	score = rand.Intn(101)
	fmt.Printf("Value\t: %d\n", score)
	grade = getScoreGradeValue(score)
	fmt.Printf("Grade\t: %s\n", grade)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Multiple Return Value")
	fmt.Println("")
	score = rand.Intn(101)
	fmt.Printf("Value\t: %d\n", score)
	grade, status = getScoreGradeValueWithStatus(score)
	fmt.Printf("Grade\t: %s\n", grade)
	fmt.Printf("Status\t: %s\n", status)
	utils.LineBreak()
}
