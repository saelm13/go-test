package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	// int 타입 인수 2개를 받아 int 타입을 반환하는 함수 타입변수
	var operator func(int, int) int
	operator = getOperator("+")

	var result = operator(10, 20) //함수타입 변수를 사용하여 함수 호출
	fmt.Println(result)
}
