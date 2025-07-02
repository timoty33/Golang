package Arrays_Slices

import "fmt"

func do(arr [3]int, slc []int) []int {
	// arr = slc SYNTAX ERROR, um array não pode receber(=) um []slice

	arr[0] = 1
	slc[0] = 2

	c := make([]int, 5)

	c[4] = 3

	copy(c, slc)

	return c
}

func Slices3() {
	arr1 := [...]int{0, 1, 2}
	slc1 := []int{0, 0, 0}

	slc2 := do(arr1, slc1)

	fmt.Println(arr1)
	fmt.Println(slc1)
	fmt.Println(slc2)
}
