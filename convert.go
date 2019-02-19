package utils

import (
	"encoding/json"
)

func ToJson(data interface{}) []byte {
	return AssertNoError(json.Marshal(data)).([]byte)
}

func ToStruct(jsonStr []byte, data interface{}) {
	AssertNoError(json.Unmarshal([]byte(jsonStr), data))
}
