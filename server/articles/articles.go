package articles

import (
	"fmt"
	"encoding/json"

	"github.com/FanJason/theAnalyst/server/env"
	"github.com/FanJason/theAnalyst/server/request"
	"github.com/FanJason/theAnalyst/server/sentiment"
)

type Response struct {
	Status       string
	TotalResults int
	Articles     []Article
}

type Payload struct {
	Data         []string `json:"data"`
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

func getAPIArticles() []Article{
	var apiKey = env.GetEnvVariable("NEWS")
	url := "http://newsapi.org/v2/everything?q=finance&sortBy=popularity&apiKey=" + apiKey
	var data Response
	request.Get(url, &data)
	return data.Articles
}

func getTitles(articles []Article) []string {
	var titles []string
	for j := 0;  j < 10; j++ {
		titles = append(titles, articles[j].Title)
	}
	return titles
}

func getContent(titles []string) string {
	payload := &Payload{
		Data: titles,
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("failed to parse titles: %v", err)
	}
	return string(bytes)
}

func setSentimentFields(currentArticle * Article, tagName string, confidence float64) {
	currentArticle.TagName = tagName
	currentArticle.Confidence = confidence
}

func getArticlesWithSentiment(articles []Article, response []sentiment.Response) []Article {
	var result []Article
	for j := 0;  j < 10; j++ {
		article := articles[j]
		classification := response[j].Classifications[0]
		setSentimentFields(&article, classification.Tag_Name, classification.Confidence)
		result = append(result, article)
	}
	return result
}

func GetArticles() []Article {
	articles := getAPIArticles()
	titles := getTitles(articles)
	response := sentiment.AnalyzeSentiment(getContent(titles))
	results := getArticlesWithSentiment(articles, response)
	return results
}
