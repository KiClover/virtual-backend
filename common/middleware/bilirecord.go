package middleware

import (
	"io"
	"net/http"
)

func RecRequest(method string, url string, key string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("x-api-key", key)
	resp, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	respbody, err := io.ReadAll(resp.Body)
	return respbody, err
}
