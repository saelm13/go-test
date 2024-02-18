package main 

import (
	"fmt"
	"strings"
)

func ToUpper(str string) string {
	var builder strings.Builder
	for _, v := range str   {      //한글자씩 순회
	if v >= 'a' && v <= 'z' {   // 소문자인지 확인
		builder.WriteRune('A' + (v - 'a'))  //소문자이면 대문자로 변경
	} else {
		builder.WriteRune(v)
	}
}
	return builder.String()
}
func main() {
	result := ToUpper("Hello, World!")
	fmt.Println(result)
}
