package store

import (
	"io/ioutil"
	"os"
	"time"
)

const lifetime = time.Second * 60 * 30

type fs struct{}

func (fs) SaveWithExpire(bytes []byte, exp time.Duration, pattern string) (string, error) {
	f, err := ioutil.TempFile("./", pattern)
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

	return f.Name(), nil
}
