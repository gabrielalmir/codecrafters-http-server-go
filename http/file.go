package httphandler

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	directory string
	matcher   string
}

func (f *File) Exists() bool {
	file := filepath.Join(f.directory, f.matcher)
	_, err := os.Stat(file)

	return !os.IsNotExist(err)
}

func (f *File) Read() ([]byte, error) {
	return os.ReadFile(filepath.Join(f.directory, f.matcher))
}

func (f *File) Handle() (string, error) {
	if f.Exists() {
		content, err := f.Read()
		if err != nil {
			log.Fatalln("Error reading file: ", err.Error())
			panic(err)
		}
		return string(content), nil
	}
	return "", errors.New("File not found")
}
