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

func getAPIArticles(topic string) []Article{
	var apiKey = env.GetEnvVariable("NEWS")
	url := "http://newsapi.org/v2/everything?q=" + topic + "&sortBy=popularity&apiKey=" + apiKey
	var data Response
	request.Get(url, &data)
	return data.Articles
}

func getTitles(articles []Article) []string {
	var titles []string
	visited := make(map[string]bool)
	for j := 0;  j < 8; j++ {
		title := articles[j].Title;
		if !visited[title] {
			titles = append(titles, title)
			visited[title] = true;
		}
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
	for j := 0;  j < 8; j++ {
		article := articles[j]
		classification := response[j].Classifications[0]
		setSentimentFields(&article, classification.Tag_Name, classification.Confidence)
		result = append(result, article)
	}
	return result
}

func GetArticles(topic string) []Article {
	articles := getAPIArticles(topic)
	titles := getTitles(articles)
	response := sentiment.AnalyzeSentiment(getContent(titles))
	results := getArticlesWithSentiment(articles, response)
	return results
}
