package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{1.0, 2.1, 3.8, 4.9, 50.1}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}
