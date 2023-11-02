package rajaongkir_go

import (
	"encoding/json"
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

func get[T interface{}](baseUrl string, endpoint string, headers map[string]string, values url.Values) (res T, err error) {
	url, err := url.JoinPath(baseUrl, endpoint)
	if err != nil {
		return res, err
	}

	if values != nil {
		if queries := values.Encode(); queries != "" {
			url = fmt.Sprintf("%s/%s", url, queries)
		}
	}

	cnt, err := sendGetRequest(url, headers)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	return
}
