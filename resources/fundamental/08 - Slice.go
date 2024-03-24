package resources

import (
	"fmt"
	"golang-demo/utils"
)

func Slice() {
	utils.LineBreak()
	fmt.Println("Basic Slice")
	fmt.Println("")
	defaultSlice := []string{}
	fmt.Println("defaultSlice\t\t:", defaultSlice)
	fmt.Println("defaultSlice Length\t:", len(defaultSlice))
	fmt.Println("defaultSlice Capacity\t:", cap(defaultSlice)) //Capacity akan mengikuti jumlah dari length slice
	defaultSlice = append(defaultSlice, "Rangga")
	fmt.Println("defaultSlice\t\t:", defaultSlice)
	fmt.Println("defaultSlice Length\t:", len(defaultSlice))
	fmt.Println("defaultSlice Capacity\t:", cap(defaultSlice)) //Capacity akan mengikuti jumlah dari length slice
	defaultSlice = append(defaultSlice, "Wiraguna")
	fmt.Println("defaultSlice\t\t:", defaultSlice)
	fmt.Println("defaultSlice Length\t:", len(defaultSlice))
	fmt.Println("defaultSlice Capacity\t:", cap(defaultSlice)) //Capacity akan mengikuti jumlah dari length slice
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Slice From Array")
	fmt.Println("")
	array := [5]int{20, 6, 2000, 1999, 1998}
	arraySlice := array[0:3]
	fmt.Println("arraySlice\t\t:", arraySlice)
	fmt.Println("arraySlice Length\t:", len(arraySlice))
	fmt.Println("arraySlice Capacity\t:", cap(arraySlice)) //Capacity akan dihitung sampai akhir array yang di slicing
	arraySlice = append(arraySlice, 1)
	fmt.Println("arraySlice\t\t:", arraySlice)
	fmt.Println("arraySlice Length\t:", len(arraySlice))
	fmt.Println("arraySlice Capacity\t:", cap(arraySlice))
	arraySlice = append(arraySlice, 2)
	fmt.Println("arraySlice\t\t:", arraySlice)
	fmt.Println("arraySlice Length\t:", len(arraySlice))
	fmt.Println("arraySlice Capacity\t:", cap(arraySlice))
	arraySlice = append(arraySlice, 3)
	fmt.Println("arraySlice\t\t:", arraySlice)
	fmt.Println("arraySlice Length\t:", len(arraySlice))
	fmt.Println("arraySlice Capacity\t:", cap(arraySlice)) //Capacity akan meningkat kelipatan dua dari saat ini jika length melebihi capacity yang ada
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Slice From make() Function")
	fmt.Println("")
	makeSliceA := make([]string, 3)     //Create slice with length = 3, value = default
	makeSliceB := make([]string, 3, 10) //Create slice with length = 3, capacity = 10, value = default
	fmt.Println("makeSliceA\t\t:", makeSliceA)
	fmt.Println("makeSliceA Length\t:", len(makeSliceA))
	fmt.Println("makeSliceA Capacity\t:", cap(makeSliceA))
	fmt.Println("makeSliceB\t\t:", makeSliceB)
	fmt.Println("makeSliceB Length\t:", len(makeSliceB))
	fmt.Println("makeSliceB Capacity\t:", cap(makeSliceB))
	makeSliceA[0] = "Rangga"
	makeSliceA[1] = "Wiraguna"
	makeSliceA[2] = "Rudiyanto"
	fmt.Println("makeSliceA\t\t:", makeSliceA)
	fmt.Println("makeSliceA Length\t:", len(makeSliceA))
	fmt.Println("makeSliceA Capacity\t:", cap(makeSliceA))
	makeSliceA = append(makeSliceA, "20 Juni 2000")
	fmt.Println("makeSliceA\t\t:", makeSliceA)
	fmt.Println("makeSliceA Length\t:", len(makeSliceA))
	fmt.Println("makeSliceA Capacity\t:", cap(makeSliceA)) //Capacity akan meningkat kelipatan dua dari saat ini jika length melebihi capacity yang ada
	utils.LineBreak()
}
