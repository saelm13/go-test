package main

import "fmt"

func main() {
	i := 0

	f := func() {
		i += 100 // i에 100을 더하기
	}
	fmt.Println("결과값1 :", i)
	i++
	fmt.Println("결과값2 :", i)
	f() // f함수를 이용해 함수 리터널 실행

	fmt.Println("결과값3 :", i)
}
