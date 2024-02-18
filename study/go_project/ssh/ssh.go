package ssh

import (
	"fmt"
	"sync"

	"golang.org/x/crypto/ssh"
)

var (
	passwordMutex sync.Mutex
)

// CiscoInfo 구조체 정의
type CiscoInfo struct {
	Host     string `json:"192.168.109.128"`
	Port     int    `json:"22"`
	Username string `json:"root"`
	Password string `json:"-"`
	Command  string `json:"netstat -nr"`
	Output   string `json:"output"`
}

// executeCommandWithSSH 함수는 SSH 클라이언트를 사용하여 호스트에 연결하고 명령을 실행
func ExecuteCommandWithSSH(ciscoInfo *CiscoInfo, command string) ([]byte, error) {
	passwordMutex.Lock()
	defer passwordMutex.Unlock()

	// SSH 연결 설정
	config := &ssh.ClientConfig{
		User: ciscoInfo.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(ciscoInfo.Password),
		},
		// HostKeyCallback 설정 추가
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH 연결
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ciscoInfo.Host, ciscoInfo.Port), config)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// SSH 세션 생성
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// 명령 실행
	output, err := session.CombinedOutput(command)
	if err != nil {
		return nil, err
	}

	return output, nil
}
