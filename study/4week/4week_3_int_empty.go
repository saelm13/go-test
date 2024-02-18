package main

import "fmt"

func PrintVal(v interface{}) { // 빈 인터페이스를 인수로 받는 함수

	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))

	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		// 그 외 타입인경우 타입과 값을 출력
		fmt.Printf("oh my god type: %T:%v\n", t, t)
	}

}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("안녕")
	PrintVal(Student{13})
}
