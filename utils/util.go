package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
)

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

/**
md5加密
*/
func Md5(password string) (md5str string) {

	data := []byte(password)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)

}

func HttpGet(request_url string) (map[string]string){

	resp, err :=   http.Get(request_url)

	if err != nil {
		// handle error
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		panic(err)
	}

	bodyMap := make(map[string]string)

	err = json.Unmarshal(body,&bodyMap)

	if err != nil {

		panic(err)

	}

	return bodyMap
}

func Uid(prefix string,num int64) string  {



	uid,_ := uuid.NewV4()


	err := uid.UnmarshalText(uid.Bytes())

	fmt.Printf("Successfully parsed: %s\n", err)
	fmt.Printf("Successfully parsed: %s\n", uid)
	return "123"
}
