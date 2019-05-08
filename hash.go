package utils

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
)

func Md5(key string) string {
	return base64.URLEncoding.EncodeToString(md5.New().Sum([]byte(key)))
}

func Sha512(key string) string {
	return base64.URLEncoding.EncodeToString(sha512.New().Sum([]byte(key)))
}
