package utils

import (
    "io/ioutil"
    "net/http"
    "net/url"
)

func GetHttpContents(rawUrl string) string {
    return string(AssertNoError(ioutil.ReadAll(AssertNoError(http.Get(rawUrl)).(*http.Response).Body)).([]byte))
}

func IsUrl(rawUrl string) bool {
    _, err := url.ParseRequestURI(rawUrl)
    return err == nil
}
