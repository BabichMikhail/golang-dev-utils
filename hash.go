package utils

import (
	"crypto/md5"
	"encoding/base64"
)

func Md5(key string) string {
	return base64.URLEncoding.EncodeToString(md5.New().Sum([]byte(key)))
}
