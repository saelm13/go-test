package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// Student의 별칭 타입 Students
type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }

// 나이 비교
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31}, {"태백", 22}, {"금강", 53},
		{"고려", 64}, {"고우", 10}}

	sort.Sort(Students(s)) //정렬
	fmt.Println(s)

}
