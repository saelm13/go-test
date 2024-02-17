package main

import "fmt"

func main() {
	var a int = 300
	var p *int // int 포인터 p 변수 선언

	p = &a // a의 메모리 주소를 b의 변수 값으로 대입 (복사)
	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리 값: %d\n", *p)

	//p가 가리키는 메모리 값 출력
	*p = 100
	fmt.Printf("a의 값: %d\n", a) //a값 확인
}
