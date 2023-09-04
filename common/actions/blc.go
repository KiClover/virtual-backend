package actions

import (
	"encoding/json"
	"go-admin/common/blcmodel"
	"io"
	"io/ioutil"
	"net/http"
)

func RecRequest(method string, url string, key string, body io.Reader) (blcmodel.Response, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("x-api-key", key)
	resp, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	respbody, err := ioutil.ReadAll(resp.Body)
	jsonString := string(respbody)
	var response blcmodel.Response
	json.Unmarshal([]byte(jsonString), &response)
	return response, err
}
