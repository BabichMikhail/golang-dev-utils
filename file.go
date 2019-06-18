package utils

import (
	"io/ioutil"
)

func GetFileContents(filePath string) []byte {
	return CheckNoError(ioutil.ReadFile(filePath)).([]byte)
}
