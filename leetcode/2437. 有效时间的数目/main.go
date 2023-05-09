package main

import (
	"fmt"
	"unsafe"
)

var h = map[byte]map[byte]int{
	'?': {
		'?': 24,
		'0': 3,
		'1': 3,
		'2': 3,
		'3': 3,
		'4': 3,
		'5': 2,
		'6': 2,
		'7': 2,
		'8': 2,
		'9': 2,
	},
	'1': {
		'?': 24,
	},
}

func main() {

	fmt.Println(countTime("?5:00"))
	fmt.Println(countTime("2?:??"))

}
func countTime(time string) int {
	timeb1 := stringtobyte(&time)
	timeb := *timeb1
	h := 1
	f := 1

	if timeb[0] == '?' {
		if timeb[1] <= '3' {
			h = 3
		} else {
			h = 2
		}
	}
	if timeb[1] == '?' {
		if timeb[0] == '?' {
			h = 24
		} else if timeb[0] == '2' {
			h = 4
		} else {
			h = 10
		}
	}

	if timeb[3] == '?' {
		f *= 6
	}
	if timeb[4] == '?' {
		f *= 10
	}

	return h * f
}
func stringtobyte(in *string) *[]byte {
	return (*[]byte)(unsafe.Pointer(in))
}
