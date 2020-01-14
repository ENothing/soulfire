package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type Header struct {
	Key   string
	Value string
}

func HttpGet(uri string, header []Header) ([]byte, error) {

	req, _ := http.NewRequest("GET", uri, nil)
	if reflect.DeepEqual(header, []Header{}) == false {
		for _, head := range header {
			req.Header.Set(head.Key, head.Value)
		}
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("http get error")
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

//HTTPPost post 请求
func HttpPost(uri string, data interface{}, header []Header) ([]byte, error) {

	byte2, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", uri, strings.NewReader(string(byte2)))
	if reflect.DeepEqual(header, []Header{}) == false {
		for _, head := range header {
			req.Header.Set(head.Key, head.Value)
		}
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("http post error")

	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
