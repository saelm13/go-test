package main

import (
	"fmt"
	"pro_netauto/database"
	"pro_netauto/file"
	"pro_netauto/ssh"
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

func main() {
	// Cisco 장비 정보 설정
	ciscoInfo := &ssh.CiscoInfo{
		Host:     "192.168.109.128", // SSH 연결할 호스트 주소 설정
		Port:     22,                // SSH 연결할 포트 번호 설정
		Username: "root",            // 하드코딩된 사용자 이름
		Password: "eve",             // 하드코딩된 패스워드
	}

	// SSH 연결 및 명령 실행
	output, err := ssh.ExecuteCommandWithSSH(ciscoInfo, "netstat -nr")
	if err != nil {
		fmt.Println("명령 실행 실패:", err)
		return
	}

	// 결과를 파일에 저장
	err = file.SaveToFile(output, "output.json")
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
		return
	}

	err = file.SaveToFile(output, "output.txt")
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
		return
	}

	// DB에 저장
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("데이터베이스 연결 실패:", err)
		return
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("DB 인스턴스 가져오기 실패:", err)
			return
		}
		sqlDB.Close() // 데이터베이스 연결을 닫습니다.
	}()

	// 데이터베이스에 저장
	ciscoDBInfo := &database.CiscoInfo{
		Host:    ciscoInfo.Host,
		Command: "netstat -nr",
		Output:  string(output),
	}
	result := database.CreateCiscoInfo(ciscoDBInfo)
	if result.Error != nil {
		fmt.Println("데이터베이스에 저장하는 중 오류 발생:", result.Error)
		return
	}

	fmt.Println("DB에 저장완료")
}
