package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//fmt.Println(minNumberOfFrogs("croakcroak"))
	//fmt.Println(minNumberOfFrogs("crcoakroak"))
	//fmt.Println(minNumberOfFrogs("croakcrook"))
	fmt.Println(minNumberOfFrogs("crocracokrakoak"))

}

func minNumberOfFrogs(croakOfFrogs string) int {
	var cage [5]int
	if len(croakOfFrogs) == 0 {
		return -1
	}
	max := -1
	croakOfFrogsbyte := stringToByte(croakOfFrogs)
	for i := 0; i < len(croakOfFrogsbyte); i++ {
		switch croakOfFrogsbyte[i] {
		case 99:
			cage[0]++
		case 114:
			if cage[0] <= 0 {
				return -1
			}
			cage[0]--
			cage[1]++
		case 111:
			if cage[1] <= 0 {
				return -1
			}
			cage[1]--
			cage[2]++
		case 97:
			if cage[2] <= 0 {
				return -1
			}
			cage[2]--
			cage[3]++
		case 107:
			if cage[3] <= 0 {
				return -1
			}
			cage[3]--
			cage[4]++
		default:
			return -1
		}
		if max < cage[0]+cage[1]+cage[2]+cage[3] {
			max = cage[0] + cage[1] + cage[2] + cage[3]
		}
	}
	if cage[0]+cage[1]+cage[2]+cage[3] != 0 {
		return -1
	}
	return max
}
func stringToByte(croakOfFrogs string) []byte {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(&croakOfFrogs))
	var t []byte
	tt := (*reflect.SliceHeader)(unsafe.Pointer(&t))
	tt.Data = ss.Data
	tt.Len = ss.Len
	tt.Cap = ss.Cap
	return t
}
