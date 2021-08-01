package xhttp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClientWrapper struct {
	*http.Client
}

func NewHttpClientWrapper(client *http.Client) *HttpClientWrapper {
	return &HttpClientWrapper{Client: client}
}

func (h *HttpClientWrapper) HttpReqGetOk(url string, timeout time.Duration) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.TODO())
	time.AfterFunc(timeout, func() {
		cancel()
	})
	req = req.WithContext(ctx)

	response, err := h.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statuscode: %d, body: %v", response.StatusCode, body)

	} else {
		return body, nil
	}
}

func (h *HttpClientWrapper) HttpReqPostOk(url string, data []byte, timeout time.Duration) ([]byte, error) {
	return h.HttpReqOk(url, http.MethodPost, data, timeout)
}

func (h *HttpClientWrapper) HttpReqOk(url, method string, data []byte, timeout time.Duration) ([]byte, error) {
	body, status, err := h.HttpReq(url, method, data, timeout)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("status: %d, body: %v", status, body))

	} else {
		return body, nil
	}
}

func (h *HttpClientWrapper) HttpReqPost(url string, data []byte, timeout time.Duration) ([]byte, int, error) {
	return h.HttpReq(url, http.MethodPost, data, timeout)
}

func (h *HttpClientWrapper) HttpReq(url, method string, data []byte, timeout time.Duration) ([]byte, int, error) {
	reqest, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, 0, err
	}
	reqest.Header.Set("Connection", "Keep-Alive")
	ctx, cancel := context.WithCancel(context.TODO())
	time.AfterFunc(timeout, func() {
		cancel()
	})
	reqest = reqest.WithContext(ctx)

	response, err := h.Do(reqest)
	if err != nil {
		return nil, 0, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, response.StatusCode, nil
}

func (h *HttpClientWrapper) HttpReqWithHeadOk(url, method string, heads map[string]string, data []byte, timeout time.Duration) ([]byte, error) {
	body, status, err := h.HttpReqWithHead(url, method, heads, data, timeout)
	if err != nil {
		return nil, err
	}

	if status < 200 || status > 299 {
		return nil, errors.New(fmt.Sprintf("status: %d, body: %v", status, body))

	} else {
		return body, nil
	}
}

func (h *HttpClientWrapper) HttpReqWithHead(url, method string, heads map[string]string, data []byte, timeout time.Duration) ([]byte, int, error) {
	reqest, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, 0, err
	}
	ctx, cancel := context.WithCancel(context.TODO())
	time.AfterFunc(timeout, func() {
		cancel()
	})
	reqest = reqest.WithContext(ctx)

	for key, val := range heads {
		reqest.Header.Set(key, val)
	}

	response, err := h.Do(reqest)
	if err != nil {
		return nil, 0, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, response.StatusCode, nil
}
