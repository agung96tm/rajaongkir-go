package rajaongkir_go

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func sendGetRequest(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
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
