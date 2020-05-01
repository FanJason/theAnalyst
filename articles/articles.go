package articles

import (
	"fmt"

	"github.com/FanJason/theAnalyst/env"
	"github.com/FanJason/theAnalyst/request"
	"github.com/FanJason/theAnalyst/sentiment"
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
	TagName      string
	Confidence   float64
}

func setSentimentFields(currentArticle * Article, tagName string, confidence float64) {
	currentArticle.TagName = tagName
	currentArticle.Confidence = confidence
}

func GetArticles(i int) []Article {
	var apiKey = env.GetEnvVariable("NEWS")
	url := "http://newsapi.org/v2/everything?q=finance&from=2020-04-01&sortBy=publishedAt&apiKey=" + apiKey
	var data Response
	request.Get(url, &data)

	for j := 0;  j < i; j++ {	
		currentArticle := data.Articles[j]
		result := sentiment.AnalyzeSentiment(currentArticle.Title)
		setSentimentFields(&currentArticle, result.Tag_Name, result.Confidence)
		fmt.Println(currentArticle.TagName)
		fmt.Println(currentArticle.Confidence)
	}
	return data.Articles
}
