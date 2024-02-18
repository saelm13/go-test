package main  

import (           //둘 이상의 패키지는 소괄호로 묶음
	"fmt" 
	"math/rand"   // 패키지명은 rand 임
)
func main () {
	fmt.Println(rand.Intn(10) + 1)  // 1부터 10 사이의 랜덤 숫자를 출력}
}