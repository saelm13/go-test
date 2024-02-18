package main 

import "fmt"

type account struct {
	balance int

}

func withrawFunc(a *account, amount int) {   // 일반 함수의 표현
	a.balance -= amount
}

func (a *account) withrawMethod(amount int) {  // 메서드의 표현
 a.balance -= amount
}   

func main() {
	a := &account {100 }   //balance 가 100인 account 포인터 변수 생성
	withrawFunc (a, 30)   //함수 형태 호출
	a.withrawMethod(30)   //메서드 형태 호출

	fmt.Printf("%d \n", a.balance)
}
