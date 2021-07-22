package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Blog struct {
	Title     string `json:"title"`
	Published int64  `json:"publishedOn"`
	URL       string `json:"url"`
	Image     string `json:"image"`
	Summary   string `json:"summary"`
}

func GetBlogs() []Blog {
	response, err := http.Get("https://codekin.tech/api/blogs")
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var blogs []Blog

	err = json.Unmarshal(responseData, &blogs)

	if err != nil {
		return nil
	}

	return blogs
}
