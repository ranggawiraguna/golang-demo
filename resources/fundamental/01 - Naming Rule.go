package resources

import (
	"fmt"
	"golang-demo/utils"
)

func NamingRule() {
	utils.LineBreak()
	fmt.Println("[>] Camel Case >> myVariable")
	fmt.Println("    For Private Variable or Class")
	fmt.Println("")
	fmt.Println("[>] Pascal Case >> MyClass")
	fmt.Println("    For Public Variable or Class")
	fmt.Println("")
	fmt.Println("[>] Snake Case >> MY_CONST")
	fmt.Println("    For Constant Variable")
	utils.LineBreak()
}
