package main 

import "fmt"

func main () {
str := "net 워크"          // 한영 문자가 섞인 문자열
for _, v := range str {   //range를 이용한 순회
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", v, v, v)  // 출력
	}
}