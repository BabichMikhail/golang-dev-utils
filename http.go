package utils

import (
	"fmt"
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

func BuildUrl(rawUrl string, params map[string]interface{}) string {
	CheckTrue(IsUrl(rawUrl), "Invalid url")

	var query string
	if len(params) > 0 {
		query = "?"
		for key, value := range params {
			switch value.(type) {
			case []string:
				for _, item := range value.([]string) {
					query += fmt.Sprintf("%s[]=%s", key, item)
				}
				break
			case string:
				query += fmt.Sprintf("%s=%s", key, value.(string))
				break
			default:
				panic("Unrecognized type")
			}
		}
	}

	return rawUrl + query
}
