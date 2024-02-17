package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct { // VIP 구조체
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"홍길동", "kildong", 100}
	vip := VIPUser{
		User{"넘버원", "NO.1", 30},
		7,
		300,
	}
	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만원\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
