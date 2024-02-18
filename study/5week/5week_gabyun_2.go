package main 

import "fmt"

func Print(args ...interface{}) string {     // 인터페이스 타입을 이용한 가변인수 
	
	for_, arg := range args{   // 모든인수 range 함수로 순회
		switch f := arg.(type) {

		 // 인수의 타입에 따른 동작
	case bool:
		val := arg.(bool)    //인터페이스 변환
	case fload64:
		val := arg.(float64)
	case int:
		val := arg.(int)
		
		//다른타입들 반복

		}

	}

}