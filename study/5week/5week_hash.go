package main 

import "fmt"

const M = 10    // 나머지 연산의 분모가 되는 숫자

func hash(d int) int {
	return d % M     // 나머지 연산
}

func main() {
	m := [M]int{}    // 값을 저장할 배열을 생성

	m[hash(25)] = 10    //키 25에 값을 설정
	m[hash(35)] = 50   // 키 35에 값을 설정
	m[hash(36)] = 70   // 키 36에 값을 하나 더 생성

	fmt.Printf("%d = %d\n", 25, m[hash(25)])
	fmt.Printf("%d = %d\n", 35, m[hash(35)]) // 값을 출력
	fmt.Printf("%d = %d\n", 36, m[hash(36)])
}
