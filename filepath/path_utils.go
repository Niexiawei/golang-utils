package filepath

import (
	"errors"
	"io/fs"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil {
		return f.IsDir(), nil
	}
	return false, err
}

func IsFile(path string) (bool, error) {
	is, err := IsDir(path)
	if err != nil {
		return false, err
	}
	return !is, nil
}
