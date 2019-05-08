package utils

import (
	"io/ioutil"
)

func ReadAllFile(filePath string) []byte {
	return AssertNoError(ioutil.ReadFile(filePath)).([]byte)
}
