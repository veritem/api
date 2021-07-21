package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Blog struct {
	Title     string `json:"title"`
	Published int64  `json:"publishedOn"`
	Url       string `json:"url"`
	Image     string `json:"image"`
	Summary   string `json:"summary"`
}

func GetBlogs() []Blog {
	response, err := http.Get("https://codekin.tech/api/blogs")
	if err != nil {
		return nil
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var blogs []Blog

	json.Unmarshal(responseData, &blogs)

	return blogs
}
