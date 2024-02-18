package file

import "os"

// SaveToFile 함수는 주어진 내용을 파일에 저장
func SaveToFile(content []byte, filename string) error {
	err := os.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
