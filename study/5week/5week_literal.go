package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {

		return func(a, b int) int { //함수 리터널을 사용해 더하기 함수 정의하고 반환함
			return a + b
		}
	} else if op == "*" {

		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("+")

	result := fn(20, 30) //함수 타입 변수를 사용하여 함수 호출
	fmt.Println(result)

}
