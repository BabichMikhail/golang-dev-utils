package utils

import (
	"encoding/json"
)

func ToJson(data interface{}) []byte {
	return CheckNoError(json.Marshal(data)).([]byte)
}

func ToStruct(jsonStr []byte, data interface{}) {
	CheckNoError(json.Unmarshal([]byte(jsonStr), data))
}
