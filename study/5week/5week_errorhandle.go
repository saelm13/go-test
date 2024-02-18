package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)   // 파일 열기
	if err != nil {
		return "", err      //에러 발생 하면 반환
	}
	defer file.Close()     // defer 함수로 종료직전 파일 닫기
	read := bufio.NewReader(file)  // 파일 내용을 읽기
	line, _ := read.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)   // 쓰기전용으로 파일이 생성
	if err != nil {   // 에러발생시 반환
		return err 
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)    //파일에 문자열 쓰기
	return err
}

const filename string = "data.txt"    //data.txt 파일을 생성

func main() {
	line, err := ReadFile(filename)       // 아래 내용들은 파일 검증 조건문
	if err != nil {
		err = WriteFile(filename, "이것은 쓰기파일이야")   
		if err != nil {
			fmt.Println("파일 생성실패!", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기실패!!", err)
			return
		}
	}
	fmt.Println("파일의 에러 내용:", line)     // 최종적으로 에러결과 출력
}
