package rajaongkir_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func sendGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
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

func get[T interface{}](baseUrl string, endpoint string, values url.Values) (res T, err error) {
	url, err := url.JoinPath(baseUrl, endpoint)
	if err != nil {
		return res, err
	}

	if values != nil {
		if queries := values.Encode(); queries != "" {
			url = fmt.Sprintf("%s/%s", url, queries)
		}
	}

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}
	return
}
