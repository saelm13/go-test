package main  

import "fmt"

func main () {
	str := "net 워크"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}
