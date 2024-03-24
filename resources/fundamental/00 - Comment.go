package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Comment() {
	utils.LineBreak()
	// Single Comment
	fmt.Println("[>] Single Comment")
	/*
		Multiline Comment
	*/
	fmt.Println("[>] Multiple Comment")
	utils.LineBreak()
}
