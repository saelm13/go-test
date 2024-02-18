package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{} // 수저와 포크 뮤텍스

	go diningProblem("A", fork, spoon, "포크", "수저") //a는 포크 먼저 획득
	go diningProblem("B", spoon, fork, "수저", "포크") //b는 수저 먼저
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥 먹을래.\n", name)
		first.Lock() //첫번째 뮤텍스 획득을 시도
		fmt.Printf("%s 획득\n", name, firstName)
		second.Lock() //두번째 뮤텍스 획득을 시도
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥 먹어야지", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock() // 뮤텍스 반납
		first.Unlock()
	}
	wg.Done()

}
