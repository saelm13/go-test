package main

import "fmt"

func main() {
	str1 := "network"
	str2 := "Catch"

	str3 := str1 + " " + str2 // str1,  " " ,과 str2를 연결

	fmt.Println(str3)

	str1 += " " + str2 // str1에 " " + str2 문자열을 붙임
	fmt.Println(str1)
	fmt.Printf("%s\n", str3)

}
