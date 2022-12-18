package main

import (
	"testing"
	"FastApiCaller/events/src"
	"encoding/json"
)

type ApiResult struct {
	Url         string `json:"url"`
	Success     bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Response     map[string]interface{} `json:"response,omitempty"`
}


func TestMakeRequest(t *testing.T) {
	response, err := apiCaller.MakeRequest("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		t.Error("Error making request:", err)
	}
	
	var apiResult ApiResult
	err = json.Unmarshal(response, &apiResult)
	if err != nil {
		t.Error("Error unmarshalling JSON:", err)
	}

}

func TestFileJsonCreator(t *testing.T) {
	var apiResult ApiResult
	apiResult.Url = "https://jsonplaceholder.typicode.com/todos/1"
	apiResult.Success = true
	apiResult.Response = map[string]interface{}{"userId": 1, "id": 1, "title": "delectus aut autem", "completed": false}
	
	var apiResultList []ApiResult
	apiResultList = append(apiResultList, apiResult)
	
	apiCaller.FileJsonCreator(apiResultList)
}
