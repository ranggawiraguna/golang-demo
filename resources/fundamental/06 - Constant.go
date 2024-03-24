package resources

import (
	"fmt"
	"golang-demo/utils"
)

const UNTYPED_CONSTANT = 4.20
const TYPED_CONSTANT int = 100
const (
	MULTIPLE_CONSTANT_1        = 10.2
	MULTIPLE_CONSTANT_2 int    = 200
	MULTIPLE_CONSTANT_3 string = "CONSTANT"
)

func Constant() {
	utils.LineBreak()
	fmt.Println("Constant")
	fmt.Println("")
	fmt.Println("[>] Untyped Constant\t:", UNTYPED_CONSTANT)
	fmt.Println("[>] Typed Constant\t:", TYPED_CONSTANT)
	fmt.Println("[>] Multiple Constant\t:")
	fmt.Printf("    %.1f\n", MULTIPLE_CONSTANT_1)
	fmt.Printf("    %d\n", MULTIPLE_CONSTANT_2)
	fmt.Printf("    %s\n", MULTIPLE_CONSTANT_3)
	utils.LineBreak()
}
