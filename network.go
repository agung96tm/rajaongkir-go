package rajaongkir_go

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func sendGetRequest(url string, headers map[string]string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func sendPostRequest(url string, headers map[string]string, values url.Values) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(values.Encode()))

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func get(baseUrl string, endpoint string, headers map[string]string, values url.Values) (res []byte, err error) {
	urlWithPath, err := url.JoinPath(baseUrl, endpoint)
	if err != nil {
		return res, err
	}

	if values != nil {
		if queries := values.Encode(); queries != "" {
			urlWithPath = fmt.Sprintf("%s/?%s", urlWithPath, queries)
		}
	}

	cnt, err := sendGetRequest(urlWithPath, headers)
	if err != nil {
		return res, err
	}

	return cnt, nil
}

func post(baseUrl string, endpoint string, headers map[string]string, values url.Values) (res []byte, err error) {
	urlWithPath, err := url.JoinPath(baseUrl, endpoint)
	if err != nil {
		return res, err
	}

	cnt, err := sendPostRequest(urlWithPath, headers, values)
	if err != nil {
		return res, err
	}

	return cnt, nil
}
