package function

import (
	"os"
	"path/filepath"
	"pusherGo/domain"
)

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeFile(filename, content string) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filename, []byte(content), 0644)
}

func WriteFile(request *domain.SaveRequest) error {
	return writeFile(request.FileName, request.Content)
}

func ReadFile(request *domain.ReadRequest) (string, error) {
	return readFile(request.FileName)
}
