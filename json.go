package goweb_json

import (
	jsoniter "github.com/json-iterator/go"
	"log"
)

func GetJSON(value interface{}) []byte {
	if value == nil {
		return make([]byte, 0)
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Println(err)
	}
	return bytes
}
