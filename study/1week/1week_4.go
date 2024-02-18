package main

import "fmt"

func main() {
	light := "green"

	if light == "red" {
		fmt.Println("정지한다")
	} else {
		fmt.Println("길을 건너간다")
	}
}
