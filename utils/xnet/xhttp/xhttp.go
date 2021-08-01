package xhttp

import (
	"net/http"
	"time"
)

var defaultClient = NewHttpClientWrapper(
	&http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 128,
			MaxConnsPerHost:     1024,
		},
		Timeout: 0,
	})

func HttpReqGetOk(url string, timeout time.Duration) ([]byte, error) {
	return defaultClient.HttpReqGetOk(url, timeout)
}

func HttpReqPostOk(url string, data []byte, timeout time.Duration) ([]byte, error) {
	return defaultClient.HttpReqPostOk(url, data, timeout)
}

func HttpReqOk(url, method string, data []byte, timeout time.Duration) ([]byte, error) {
	return defaultClient.HttpReqOk(url, method, data, timeout)
}

func HttpReqPost(url string, data []byte, timeout time.Duration) ([]byte, int, error) {
	return defaultClient.HttpReqPost(url, data, timeout)
}

func HttpReq(url, method string, data []byte, timeout time.Duration) ([]byte, int, error) {
	return defaultClient.HttpReq(url, method, data, timeout)
}

func HttpReqWithHeadOk(url, method string, heads map[string]string, data []byte, timeout time.Duration) ([]byte, error) {
	return defaultClient.HttpReqWithHeadOk(url, method, heads, data, timeout)
}

func HttpReqWithHead(url, method string, heads map[string]string, data []byte, timeout time.Duration) ([]byte, int, error) {
	return defaultClient.HttpReqWithHead(url, method, heads, data, timeout)
}
