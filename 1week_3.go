package main

import "fmt"

func main() {
	const 파이1 float64 = 3.1415926535
	var 파이2 float64 = 3.1415926535

	파이2 = 4

	fmt.Printf("원주율: %f\n", 파이1)
	fmt.Printf("원주율 %f\n", 파이2)

}
