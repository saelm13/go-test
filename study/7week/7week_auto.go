package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Cisco 장비 정보 설정
	ciscoInfo := &CiscoInfo{
		Host:     "192.168.123.123",
		Port:     22, // Cisco의 SSH 포트 번호
		Username: "root",
		Password: "eve",
	}

	// SSH 연결 설정
	config := &ssh.ClientConfig{
		User: ciscoInfo.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(ciscoInfo.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 실제 운영 환경에서는 보안을 위해 수정 필요
	}

	// SSH 연결
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ciscoInfo.Host, ciscoInfo.Port), config)
	if err != nil {
		fmt.Println("Failed to dial:", err)
		return
	}
	defer client.Close()

	// SSH 세션 열기
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}
	defer session.Close()

	// 명령 실행
	cmd := "ls" // 장비에서 실행할 실제 명령
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Println("Failed to run command:", err)
		return
	}

	// 결과 출력
	fmt.Println(string(output))
}

type CiscoInfo struct {
	Host     string
	Port     int
	Username string
	Password string
}
