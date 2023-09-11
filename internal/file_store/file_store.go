package file_store

import (
	"github.com/tomhobson/freematics2prom/internal/domain"
	"os"
)

type fileStore struct {
}

func NewFileStore() domain.FileStore {
	return fileStore{}
}

func (f fileStore) ReadFile(filePath string) (string, error) {
	// Read the content of the file located at the given filePath
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the file content to a string and return it
	fileContent := string(content)
	return fileContent, nil
}
