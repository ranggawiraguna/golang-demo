package resources

import (
	"fmt"
	"golang-demo/utils"
	"math/rand"
)

func Loop() {
	utils.LineBreak()
	fmt.Println("Basic Loop")
	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Printf("Row-%d\n", i+1)
	}
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Nested Loop")
	fmt.Println("")
	for i := 0; i < 3; i++ {
		fmt.Printf("Row-%d\n", i+1)
		for j := 0; j < 3; j++ {
			fmt.Printf("Row-%d Column-%d\n", i+1, j+1)
		}
	}
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Range Keyword")
	fmt.Println("")
	fruits := [3]string{"Rangga", "Wiraguna", "Rudiyanto"}
	for idx, val := range fruits {
		fmt.Printf("Word-%v\t: %s\n", idx+1, val)
	}
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Continue Keyword")
	fmt.Println("")
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Printf("Note\t: The loop skipped because %d %% 2 == 1\n", i)
			continue
		}
		fmt.Printf("Value\t: %d\n", i)
	}
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Break Keyword")
	fmt.Println("")
	for i := rand.Intn(10); i < 10; i = rand.Intn(10) {
		fmt.Printf("Value\t: %d\n", i)
		if i%3 == 0 {
			fmt.Printf("Note\t: The loop stops because %d %% 3 == 0\n", i)
			break
		}
	}
	utils.LineBreak()
}
