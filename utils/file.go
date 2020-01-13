package utils

import "os"

/**
创建文件并接着写入
*/
func OpenFile(filepath, filename string) *os.File {

	_, err := os.Stat(filepath)

	switch {
	case os.IsNotExist(err):
		Mkdir(filepath)
	case os.IsPermission(err):
		panic(err)
	}

	file, err := os.OpenFile(filepath+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	return file

}

/**
创建文件夹
*/
func Mkdir(filepath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+filepath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
