package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i <= 10; i++ { // 요소를 하나씩 추가함
		slice = append(slice, i)
	}
	slice = append(slice, 11, 12, 13, 14, 15) // 한번에 여러 요소 추가
	fmt.Println(slice)
}
