package apiCaller

import (
	"fmt"
	"sync"
	"time"
	"net/http"
)

type ApiResult struct {
	Url         string `json:"url"`
	Success     bool   `json:"success"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func FileJsonCreator(Result []ApiResult) {
	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile("result.json", file, 0644)
}

func FileTxtCreator(Result []ApiResult) {
	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile("result.txt", file, 0644)
}

func FileCsvCreator(Result []ApiResult) {
	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile("result.csv", file, 0644)
}

func FileXlsxCreator(Result []ApiResult) {
	file, _ := json.MarshalIndent(Result, "", " ")
	_ = ioutil.WriteFile("result.xlsx", file, 0644)
}

