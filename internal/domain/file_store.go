package domain

type FileStore interface {
	ReadFile() (string, error)
}
