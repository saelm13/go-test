package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 //1 바이트  (작은 타입부터 큰타입 순으로 정렬)
	B int8
	C int8
	D int //8 바이트
	E int
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
