package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
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

func HttpGet(request_url string) map[string]string {

	resp, err := http.Get(request_url)

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

	err = json.Unmarshal(body, &bodyMap)

	if err != nil {

		panic(err)

	}

	return bodyMap
}

func Uid(prefix string) string {

	uid := uuid.Must(uuid.NewV4())

	return prefix + uid.String()

}

func JsonEncode(v interface{}) string {

	jsonStr, _ := json.Marshal(v)

	return string(jsonStr)

}

func JsonDecode(v string) interface{} {

	var bodyMap interface{}

	json.Unmarshal([]byte(v), &bodyMap)

	return bodyMap

}

func TimeFormat(t time.Time, formatType int64) (formatTime string) {

	if formatType == 0 {
		formatTime = t.Format("2006-01-02 15:04:05")
	} else {
		formatTime = t.Format("2006.01.02")
	}
	return formatTime
}

func TimeSpan(t time.Time) string {

	nowTime := time.Now().Unix()
	var timeStr string

	tTime := t.Unix()

	resTime := nowTime - tTime

	if resTime > 0 && resTime < 3600 {

		timeStr = string(resTime/60) + "分钟前"

	} else if resTime >= 3600 && resTime < 86400 {

		timeStr = strconv.FormatInt(resTime/3600, 10) + "小时前"

	} else if resTime >= 86400 && resTime < 604800 {

		timeStr = string(resTime/86400) + "天前"

	} else {

		timeStr = t.Format("01月02日")

	}

	return timeStr

}
