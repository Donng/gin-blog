package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(file multipart.File) (int, error) {
	content, err := ioutil.ReadAll(file)

	return len(content), err
}

func GetExt(file string) string {
	return path.Ext(file)
}

func CheckExist(file string) bool {
	_, err := os.Stat(file)

	return os.IsNotExist(err)
}

func CheckPermission(file string) bool {
    _, err := os.Stat(file)

    return os.IsPermission(err)
}

func IsNotExistMkDir(filePath string) error {
	if exist := CheckExist(filePath); exist == false {
		if err := MkDir(filePath); err != nil {
			return err
		}
	}

	return nil
}

func MkDir(filePath string) error {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Open(filename string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}