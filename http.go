package utils

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "path"
)

func GetHttpContents(rawUrl string) []byte {
    return CheckNoError(ioutil.ReadAll(CheckNoError(http.Get(rawUrl)).(*http.Response).Body)).([]byte)
}

func IsUrl(rawUrl string) bool {
    _, err := url.ParseRequestURI(rawUrl)
    return err == nil
}

func UrlJoin(parts ...string) string {
    finalUrl := ""
    for _, part := range parts {
        if finalUrl == "" {
            finalUrl = part
        } else {
            parsedUrl := CheckNoError(url.Parse(finalUrl)).(*url.URL)
            parsedUrl.Path = path.Join(parsedUrl.Path, part)
            finalUrl = parsedUrl.String()
        }
    }

    return finalUrl
}
