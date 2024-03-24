package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Array() {
	var arrayString = [10]string{}            //Inital with default value
	var arrayNumberA = [3]int{20, 6, 2000}    //Initial with assign value
	var arrayNumberB = [5]int{100}            //Initial with partially assign value
	var arrayNumberC = [4]int{0: 200, 2: 300} //Initial with specific index

	utils.LineBreak()
	fmt.Println("Array")
	fmt.Println("")
	fmt.Println("arrayString\t: ", arrayString)
	fmt.Println("arrayNumberA\t: ", arrayNumberA)
	fmt.Println("arrayNumberB\t: ", arrayNumberB)
	fmt.Println("arrayNumberC\t: ", arrayNumberC)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Array Manipulation")
	fmt.Println("")
	arrayString[0] = "Rangga"
	arrayString[2] = "Wiraguna"
	arrayNumberA[0] = 40
	arrayNumberB[1] = 1000
	arrayNumberC[1] = 250
	fmt.Println("arrayString\t: ", arrayString)
	fmt.Println("arrayNumberA\t: ", arrayNumberA)
	fmt.Println("arrayNumberB\t: ", arrayNumberB)
	fmt.Println("arrayNumberC\t: ", arrayNumberC)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Array Length")
	fmt.Println("")
	fmt.Println("arrayString Length\t: ", len(arrayString))
	fmt.Println("arrayNumberA Length\t: ", len(arrayNumberA))
	fmt.Println("arrayNumberB Length\t: ", len(arrayNumberB))
	fmt.Println("arrayNumberC Length\t: ", len(arrayNumberC))
	utils.LineBreak()
}
