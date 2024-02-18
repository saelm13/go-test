package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker //기본값은 nil
	att.Attack()     //att가 nil이기 때문에 런타임 오류발생
}
