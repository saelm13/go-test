package main

import (
	"fmt"
	"time"
)

func PrintHanguls() {
	hanguls := []rune{'도', '레', '미', '파', '솔', '라', '시'} //슬라이스 룬 타입 정의
	for _, v := range hanguls {
		time.Sleep(500 * time.Millisecond) // 0.5초 동안 일시정지
		fmt.Printf("%c ", v)               //  v의 변수에 대해 문자열을 출력

	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(510 * time.Millisecond) // 0.51초 동안 일시정지
		fmt.Printf("%d ", i)               // i의 변수에 숫자값을 출력
	}
}

func main() {
	go PrintHanguls() //고 루틴 한글 실행
	go PrintNumbers() //고 루틴 숫자 실행

	time.Sleep(4 * time.Second) // 4초 동안 대기한 후 고루틴 종료
}
