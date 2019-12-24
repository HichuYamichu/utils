package store

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const lifetime = time.Second * 60 * 30

func init() {
	os.Mkdir("tmp", os.ModePerm)
}

type FS struct{}

func (FS) SaveTemp(bytes []byte, pattern string) (string, error) {
	f, err := ioutil.TempFile("tmp", pattern)
	if err != nil {
		return "", err
	}

	go func() {
		<-time.After(lifetime)
		os.Remove(f.Name())
	}()

	if _, err := f.Write(bytes); err != nil {
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}

	return filepath.Base(f.Name()), nil
}
