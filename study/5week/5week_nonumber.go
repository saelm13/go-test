package main

import "fmt"

func solution(numbers []int) int {
	if len(numbers) < 1 || len(numbers) > 9 {
		return -1
	}

	total := 0

	for i := 0; i < 10; i++ {
		total += i
	}

	for _, num := range numbers {
		total -= num
	}

	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 6, 7, 8}
	result := solution(numbers)
	fmt.Println(result) // 출력: 5 + 9 = 14
}
