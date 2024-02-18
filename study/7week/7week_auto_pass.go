package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh" // go get golang.org/x/crypto/ssh 을 이용해 패키지를 가져옴

	/* Go 언어의 암호화 및 SSH 프로토콜과 관련된 기능을 제공. 이 패키지는 SSH 클라이언트 및 서버를 구현하는 데 사용됩니다.
	   여러 암호화 알고리즘 및 SSH 프로토콜의 다양한 기능을 지원  */
	"golang.org/x/crypto/ssh/terminal" //go get 을 이용해 외부 패키지 terminal 패키지를 가져옴
)

// CiscoInfo 구조체 정의
type CiscoInfo struct {
	Host     string // SSH 연결할 호스트 주소
	Port     int    // SSH 연결할 포트 번호
	Username string // SSH 연결에 사용할 사용자 이름
	Password string // SSH 연결에 사용할 패스워드
}

func main() {
	var password string
	var attemptCount = 0
	maxAttempts := 3

	for {
		// 사용자로부터 ID와 Password 입력 받기
		fmt.Print("사용자명: ")
		username, err := getUserInput()
		if err != nil {
			fmt.Println("에러: 사용자명 입력 -", err)
			return
		}

		// 잘못된 사용자명 또는 패스워드를 처리하기 위한 루프
		for attemptCount < maxAttempts {
			// getPassword 함수를 통해 사용자로부터 패스워드 입력 받기
			password, err = getPassword()
			if err != nil {
				fmt.Println("에러: 패스워드 입력 -", err)
				return
			}

			// Cisco 장비 정보 설정
			ciscoInfo := &CiscoInfo{
				Host:     "192.168.109.128", // SSH 연결할 호스트 주소 설정
				Port:     22,                // SSH 연결할 포트 번호 설정
				Username: username,          // 사용자로부터 입력받은 사용자 이름 설정
				Password: password,          // 사용자로부터 입력받은 패스워드 설정
			}

			// Check the validity of the entered username
			if !isValidUsername(ciscoInfo.Username) {
				attemptCount++
				fmt.Printf("잘못된 사용자명입니다 (%d/%d)\n", attemptCount, maxAttempts)
				if attemptCount == maxAttempts {
					fmt.Println("최대 허용 시도 횟수 초과. 프로그램 종료.")
					return
				}

				// 사용자명이 잘못되었으면 다시 사용자명을 입력 받기
				break
			}

			// SSH 연결 설정
			config := &ssh.ClientConfig{
				User: ciscoInfo.Username,
				Auth: []ssh.AuthMethod{
					ssh.Password(ciscoInfo.Password),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				// 실제 운영 환경에서는 보안을 위해 수정 필요
			}

			// SSH 연결
			client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ciscoInfo.Host, ciscoInfo.Port), config)
			if err != nil {
				attemptCount++
				fmt.Printf("인증 실패 (%d/%d)\n", attemptCount, maxAttempts)
				if attemptCount == maxAttempts {
					fmt.Println("최대 허용 시도 횟수 초과. 프로그램 종료.")
					return
				}
				break // 인증 실패 시 루프를 종료하여 사용자에게 다시 입력 받음
			}

			clientClose := make(chan struct{})
			defer func() {
				client.Close()     // defer 함수 사용하여 클라이언트 닫기
				close(clientClose) // 클라이언트 닫힘을 신호하는 채널 닫기
			}()

			// SSH 세션 열기
			session, err := client.NewSession()
			if err != nil {
				fmt.Println("세션 생성 실패:", err)
				return
			}

			defer session.Close() //defer 지연함수 사용하여 세션 닫기

			// 명령 실행
			cmd := "ls" // 장비에서 실행할 실제 명령
			output, err := session.CombinedOutput(cmd)
			if err != nil {
				fmt.Println("명령 실행 실패:", err)
				fmt.Println("다시 시도하세요.")
				continue // 명령 실행 실패 시 루프를 계속 실행하여 사용자에게 다시 입력 받음
			}

			// 결과 출력
			fmt.Println(string(output))

			// 결과를 JSON 파일에 저장
			saveJSONToFile(output, "output.json")

			// 결과를 텍스트 파일에 저장
			saveTextToFile(output, "output.txt")

			return
		}
	}
}

// getUserInput 함수는 사용자로부터 키보드 입력을 받아오는 함수
func getUserInput() (string, error) {
	// bufio.NewReader(os.Stdin).ReadString('\n')을 사용하여 개행 문자까지 읽어오도록 수정
	scanner := bufio.NewReader(os.Stdin)
	input, err := scanner.ReadString('\n')
	if err != nil {
		return "", err
	}
	// 개행 문자를 제거하여 반환
	return strings.TrimSpace(input), nil
}

// getPassword 함수는 사용자로부터 안전하게 패스워드를 입력받는 함수
func getPassword() (string, error) {
	fmt.Print("비밀번호 입력: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println() // 개행을 출력하여 사용자가 다음에 입력할 때 정렬이 보기 좋게 되도록 함
	return strings.TrimSpace(string(password)), nil
}

// saveJSONToFile 함수는 주어진 내용을 JSON 파일에 저장
func saveJSONToFile(content []byte, filename string) {
	err := os.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
	}
}

// saveTextToFile 함수는 주어진 내용을 텍스트 파일에 저장
func saveTextToFile(content []byte, filename string) {
	err := os.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
	}
}

// isValidUsername 함수는 사용자명이 유효한지 검사
func isValidUsername(username string) bool {

	return username == "root" // root"와 동일하면 true를 반환하고, 그렇지 않으면 false를 반환
}
