package apiCaller

import (
	"encoding/json"
	"io/ioutil"
)

type ApiResult struct {
	Url         string `json:"url"`
	Success     bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Response     map[string]interface{} `json:"response,omitempty"`
}

func FileJsonCreator(Result []ApiResult,name string) {

	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile(name+".json", file, 0644)
}

func FileTxtCreator(Result []ApiResult,name string) {
	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile(name+".json", file, 0644)
}

func FileCsvCreator(Result []ApiResult,name string) {
	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile(name+".json", file, 0644)
}

func FileXlsxCreator(Result []ApiResult,name string) {
	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile(name+".json", file, 0644)
}

