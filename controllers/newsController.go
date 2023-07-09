package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ramnath-1998/poorhouse/models"
)

type SourcesArgs struct {
	Language string
	Country  string
}

func GetNewsFromApi(url string) (*models.NewsResponseModel, error) {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", os.Getenv("NEWS_API_KEY"))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	var newsResponse models.NewsResponseModel
	err = json.NewDecoder(res.Body).Decode(&newsResponse)
	if err != nil {

		return nil, err

	}

	return &newsResponse, nil
}
