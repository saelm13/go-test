package main

import "fmt"

func main() {

	// 큰따옴표에서 여러줄을 표현하려면 \n을 사용해야 함
	poet1 := "동해물과 백두산이 마르고 닳도록\n하느님이 보우하사 우리나라만세\n"

	// 백쿼트에서는 여러 줄 표현에 특수문자가 필요없음
	poet2 := `동해물과 백두산이 마르고 닳도록
	하느님이 보우하사 우리나라만세`

	fmt.Println(poet1)
	fmt.Println(poet2)
}
