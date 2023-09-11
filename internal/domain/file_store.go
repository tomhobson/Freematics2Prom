package domain

type FileStore interface {
	ReadFile(filePath string) (string, error)
}
