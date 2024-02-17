package main

import "fmt"

type Student struct {
	Age   int // 대문자 시작 필드는 외부로 공개됨
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이:%d 번호:%d 점수:%.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{20, 11, 99.9}
	// student 구조체 모든 필드가 students2로 복사됨
	student2 := student

	PrintStudent(student2)

}
