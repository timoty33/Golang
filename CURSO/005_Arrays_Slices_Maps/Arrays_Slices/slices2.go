package Arrays_Slices

import "fmt"

func Slices2() {
	str := "timoty"
	slc := []byte(str)

	fmt.Println(len(slc), slc)      // 6 [116 105 109 111 116 121]
	fmt.Println(slc[2], str[2])     // 109 109
	fmt.Println(slc[2:], str[2:])   // [109 111 116 121] moty
	fmt.Println(slc[:2], str[:2])   // [116 105] ti
	fmt.Println(slc[2:5], str[2:5]) // mot
}
