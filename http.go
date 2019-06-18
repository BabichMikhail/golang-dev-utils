package utils

import (
    "io/ioutil"
    "net/http"
    "net/url"
)

func GetHttpContents(rawUrl string) []byte {
    return CheckNoError(ioutil.ReadAll(CheckNoError(http.Get(rawUrl)).(*http.Response).Body)).([]byte)
}

func IsUrl(rawUrl string) bool {
    _, err := url.ParseRequestURI(rawUrl)
    return err == nil
}
