package articles

import (
	"fmt"
	"encoding/json"
	"os"
	"net/http"

	"github.com/joho/godotenv"
)

type Response struct {
	Status       string
	TotalResults int
	Articles     []Article
}

type Article struct {
	Author       string
	Title        string
	Description  string
	Url          string
	UrlToImage   string
	PublishedAt  string
	Content      string
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("failed to get env variable: %v", err)
	}
	return os.Getenv(key)
}

func GetArticles() []Article {
	var apiKey = getEnvVariable("KEY")
	url := "http://newsapi.org/v2/everything?q=finance&from=2020-04-01&sortBy=publishedAt&apiKey=" + apiKey

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get api data, error: %v", err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var data Response
	err = decoder.Decode(&data)

	if err != nil {
		fmt.Printf("failed to parse api response: %v", err)
	}
	return data.Articles
}
