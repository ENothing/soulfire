package utils

import "encoding/json"

func JsonEncode(v interface{}) string {

	jsonStr, _ := json.Marshal(v)

	return string(jsonStr)

}

func JsonDecode(v string) map[string]interface{} {

	var bodyMap map[string]interface{}

	json.Unmarshal([]byte(v), &bodyMap)

	return bodyMap

}
