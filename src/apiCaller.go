package apiCaller

import (
	"net/http"
	"encoding/json"
)


func makeRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		apiResult := ApiResult{Url: url, Success: false, ErrorMessage: err.Error()}
		return json.Marshal(apiResult)
	}
	defer res.Body.Close()

	result := ApiResult{Url: url, Success: res.StatusCode == 200}
	if !result.Success {
		return json.Marshal(result)
	}

	var responseJson map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&responseJson)
	if err != nil {
		apiResult := ApiResult{Url: url, Success: false, ErrorMessage: err.Error()}
		return json.Marshal(apiResult)
	}
	result.Response = responseJson
	return json.Marshal(result)
}

