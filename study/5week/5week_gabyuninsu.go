package main

import "fmt"

func sum(nums ...int) int { // 가변 인수를 받는 함수
	sum := 0
	fmt.Printf("nums 타입: %T\n", nums) // nums의 타입을 출력
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3))       // 인수 3개 사용
	fmt.Println(sum(100, 200))      // 인수 2개 사용
	fmt.Println(sum())              // 인수 0개 사용
	
}
