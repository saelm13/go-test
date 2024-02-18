package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// CiscoInfo 구조체 정의
type CiscoInfo struct {
	gorm.Model
	Host     string
	Port     int
	Username string
	Password string
	Command  string
	Output   string
}

// InitDB 함수는 MySQL 데이터베이스에 연결
func InitDB() (*gorm.DB, error) {
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

	DB = db

	db.AutoMigrate(&CiscoInfo{})
	return db, nil
}

// CreateCiscoInfo 함수는 CiscoInfo를 데이터베이스에 저장합니다.
func CreateCiscoInfo(ciscoInfo *CiscoInfo) error {
	result := DB.Create(ciscoInfo)
	return result.Error
}
