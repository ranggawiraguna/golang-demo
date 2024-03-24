package main

import (
	"fmt"
	fundamental "golang-demo/resources/fundamental"
	"golang-demo/utils"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Print("\033[H\033[2J")

	entries, err := os.ReadDir("./resources")

	if err == nil {
		utils.LineBreak()

		var paths []string
		for _, entry := range entries {
			paths = append(paths, cases.Title(language.Und, cases.NoLower).String(strings.Split(strings.ReplaceAll(entry.Name(), ".go", ""), " - ")[1]))
		}

		fmt.Println("Learning Path")
		fmt.Println()
		for index, path := range paths {
			fmt.Printf("[%d] %s\n", index, path)
		}

		var selectedNumber int
		fmt.Println()
		fmt.Print("Choose Path : ")
		fmt.Scan(&selectedNumber)

		utils.LineBreak()

		switch selectedNumber {
		case 0:
			fundamental.Comment()
		case 1:
			fundamental.NamingRule()
		case 2:
			fundamental.Declaration()
		case 3:
			fundamental.Output()
		case 4:
			fundamental.Formatting()
		case 5:
			fundamental.Variable()
		case 6:
			fundamental.Constant()
		case 7:
			fundamental.Array()
		case 8:
			fundamental.Slice()
		case 9:
			fundamental.SliceModify()
		case 10:
			fundamental.Arithmetic()
		case 11:
			fundamental.Assignment()
		case 12:
			fundamental.Comparison()
		case 13:
			fundamental.Logical()
		case 14:
			fundamental.Condition()
		case 15:
			fundamental.Switch()
		case 16:
			fundamental.Loop()
		default:
			fmt.Println("")
		}
	}

}
