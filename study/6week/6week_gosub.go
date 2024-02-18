package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // ❶ waitGroup 객체 생성

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d이다.\n", a, b, sum)
	wg.Done() // ❸ 작업이 완료됨을 표시
}

func main() {
	wg.Add(5) // ❷ 총 작업 개수 설정
	for i := 0; i < 5; i++ {
		go SumAtoB(1, 100)
	}

	wg.Wait() // ❹ 모든 작업이 완료되길 기다림.
	fmt.Println("모든 계산이 완료되었습니다.")
}
