package main

import (
	"fmt"
	"os"
) //패키지를 여러개 사용할때 ( ) 이용

func main() {

	f, err := os.Create("nwcatch2.txt") //nwcatch2.txt 파일 생성함
	if err != nil {                     // 에러 확인
		fmt.Println("파일 열기 실패", err)
		return
	}

	defer fmt.Println(" 반드시 호출")  // 지연 수행 코드
	defer f.Close()               //지연 수행 코드
	defer fmt.Println("파일 닫기 성공") // 지연 수행 코드

	fmt.Println("파일에 네전따 쓰기.")
	fmt.Fprintln(f, "network catch") //파일에 network catch 작성

}
