package articles

import (
	"github.com/FanJason/theAnalyst/env"
	"github.com/FanJason/theAnalyst/request"
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

func GetArticles() []Article {
	var apiKey = env.GetEnvVariable("NEWS")
	url := "http://newsapi.org/v2/everything?q=finance&from=2020-04-01&sortBy=publishedAt&apiKey=" + apiKey
	var data Response
	request.Get(url, &data)
	return data.Articles
}
