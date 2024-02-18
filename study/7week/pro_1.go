package main

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/crypto/ssh"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// CiscoInfo 구조체 정의
type CiscoInfo struct {
	gorm.Model
	Host     string `json:"192.168.109.128"`
	Port     int    `json:"22"`
	Username string `json:"root"`
	Password string `json:"-"`
	Command  string `json:"netstat -nr"`
	Output   string `json:"output"`
}

var (
	passwordMutex sync.Mutex
)

func main() {
	// Cisco 장비 정보 설정
	ciscoInfo := &CiscoInfo{
		Host:     "192.168.109.128", // SSH 연결할 호스트 주소 설정
		Port:     22,                // SSH 연결할 포트 번호 설정
		Username: "root",            // 하드코딩된 사용자 이름
		Password: "eve",             // 하드코딩된 패스워드
	}

	// SSH 연결 및 명령 실행
	output, err := executeCommandWithSSH(ciscoInfo, "netstat -nr")
	if err != nil {
		fmt.Println("명령 실행 실패:", err)
		return
	}

	// 결과를 파일에 저장
	saveToFile(output, "output.json")
	saveToFile(output, "output.txt")

	// DB에 저장
	db, err := initDB()
	if err != nil {
		fmt.Println("데이터베이스 연결 실패:", err)
		return
	}
	defer dbClose(db)

	// GORM 디버그 로깅 활성화
	db.Logger = logger.Default.LogMode(logger.Info)

	// 트랜잭션 사용
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 데이터베이스에 저장
	ciscoInfo.Command = "netstat -nr"
	ciscoInfo.Output = string(output)
	result := tx.Create(&ciscoInfo)
	if result.Error != nil {
		fmt.Println("데이터베이스에 저장하는 중 오류 발생:", result.Error)
		tx.Rollback()
		return
	}

	// 트랜잭션 커밋
	tx.Commit()

	fmt.Println("DB에 저장완료")
}

// executeCommandWithSSH 함수는 SSH 클라이언트를 사용하여 호스트에 연결하고 명령을 실행
func executeCommandWithSSH(ciscoInfo *CiscoInfo, command string) ([]byte, error) {
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

// saveToFile 함수는 주어진 내용을 파일에 저장
func saveToFile(content []byte, filename string) {
	err := os.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		fmt.Println("파일 쓰기 실패:", err)
	}
}

// initDB 함수는 MySQL 데이터베이스에 연결
func initDB() (*gorm.DB, error) {
	// MySQL 드라이버 사용
	dsn := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns 설정 추가 (선택 사항)
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10) // 필요에 따라 조절

	db.AutoMigrate(&CiscoInfo{})
	return db, nil
}

// dbClose 함수는 *gorm.DB를 닫는 함수
func dbClose(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB 인스턴스 가져오기 실패!")
		return
	}
	sqlDB.Close() // *sql.DB를 닫음
}
