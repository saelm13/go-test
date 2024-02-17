package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	/* 기존 바이트 패딩 문제점
	A int8
	B int
	C float64
	D uint16
	E int
	F float32
	*/
	A int8    // 1바이트
	B uint16  // 2바이트
	C float32 // 4바이트
	D float64 // 8바이트
	E int     // 8바이트
	F int     // 8바이트
}

func main() {
	user := User{1, 2, 3, 4, 5, 6}
	fmt.Println(unsafe.Sizeof(user))
}
