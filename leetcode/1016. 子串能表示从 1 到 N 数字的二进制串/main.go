package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(queryString("0110", 3))
	fmt.Println(queryString("0110", 4))

}

func queryString(s string, n int) bool {
	var res = true
	sbye := stringtobyte(s)
	for i := n; i >= 1; i-- {
		bs := stringtobyte(fmt.Sprintf("%b", i))
		if db(sbye, bs) {
		} else {
			return false
		}
	}
	return res
}
func stringtobyte(string string) []byte {
	return *(*[]byte)(unsafe.Pointer(&string))
}
func db(s, y []byte) bool {
	slen := len(s)
	ylen := len(y)
	if slen < ylen {
		return false
	}
	for i := 0; i <= slen-ylen; i++ {
		var ss int
		for i1 := 0; i1 < ylen; i1++ {
			if s[i+i1] == y[i1] {
				ss++
			} else {
				break
			}
		}
		if ss == ylen {
			return true
		}
	}
	return false
}
