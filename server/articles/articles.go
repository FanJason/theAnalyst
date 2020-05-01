package articles

import (
	"github.com/FanJason/theAnalyst/server/env"
	"github.com/FanJason/theAnalyst/server/request"
	"github.com/FanJason/theAnalyst/server/sentiment"
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

func GetArticles() []Article {
	var apiKey = env.GetEnvVariable("NEWS")
	url := "http://newsapi.org/v2/everything?q=finance&from=2020-04-01&sortBy=popularity&apiKey=" + apiKey
	var data Response
	request.Get(url, &data)
	var result []Article
	for j := 0;  j < 10; j++ {
		currentArticle := data.Articles[j]
		classification := sentiment.AnalyzeSentiment(currentArticle.Title)
		setSentimentFields(&currentArticle, classification.Tag_Name, classification.Confidence)
		result = append(result, currentArticle)
	}
	return result
}
