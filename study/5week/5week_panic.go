package main

import "fmt"

func f() {
	fmt.Println("f() 함수시작")
	defer func () {     //패닉 복구
		if r := recover(); r != nil {
			fmt.Println("패닉 복구 -", r)

		}
	} ()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))     // h함수  호출 - 패닉상태

}

func h(a, b int) int {
	if b == 0{ 
		panic("분모는 0이 될수 없음")   //패닉 발생!!!
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램 계속실행중")    // 프로그램 무한 로딩
}

