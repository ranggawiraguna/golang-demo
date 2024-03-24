package resources

import (
	"fmt"
	"golang-demo/utils"
)

func SliceModify() {
	utils.LineBreak()
	fmt.Println("Slice Modify")
	fmt.Println("")
	sliceNumber := []int{10, 21, 32}
	fmt.Println("sliceNumber[0]\t:", sliceNumber[0])
	fmt.Println("sliceNumber[1]\t:", sliceNumber[1])
	fmt.Println("sliceNumber[2]\t:", sliceNumber[2])
	sliceNumber[0], sliceNumber[1], sliceNumber[2] = 110, 220, 330
	fmt.Println("sliceNumber[0]\t:", sliceNumber[0])
	fmt.Println("sliceNumber[1]\t:", sliceNumber[1])
	fmt.Println("sliceNumber[2]\t:", sliceNumber[2])
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Slice Append")
	fmt.Println("")
	sliceString := []string{"A", "B", "C"}
	fmt.Println("sliceString\t:", sliceString)
	sliceString = append(sliceString, "D", "E")
	fmt.Println("sliceString\t:", sliceString)
	sliceStringNew := []string{"F", "G"}
	sliceString = append(sliceString, sliceStringNew...)
	fmt.Println("sliceString\t:", sliceString)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Slice Copy")
	fmt.Println("")
	sliceFloat := []float32{4.2, 3.3, 9.1}
	fmt.Println("sliceFloat\t:", sliceFloat)
	sliceFloatNew := make([]float32, len(sliceFloat))
	copy(sliceFloatNew, sliceFloat)
	fmt.Println("sliceFloatNew\t:", sliceFloatNew)
	utils.LineBreak()
}
