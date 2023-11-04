package httphandler

import (
	"errors"
	"os"
)

type File struct {
	directory string
}

func (f *File) Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f *File) Handle(path string) (string, error) {
	if f.Exists(path) {
		file, _ := Read(path)
		return string(file), nil
	}
	return "", errors.New("File not found")
}
