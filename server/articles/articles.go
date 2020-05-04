package articles

import (
	"fmt"
	"encoding/json"

	// "github.com/FanJason/theAnalyst/server/env"
	// "github.com/FanJason/theAnalyst/server/request"
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
	// var apiKey = env.GetEnvVariable("NEWS")
	// url := "http://newsapi.org/v2/everything?q=finance&sortBy=popularity&apiKey=" + apiKey
	// var data Response
	// request.Get(url, &data)
	// return data.Articles
	article1 := Article{
        Author: "Lisa Rowan on Two Cents, shared by Lisa Rowan to Lifehacker",
        Confidence: 0.94,
        Content: "Mere weeks ago, in the Before Times, I got an email from a reader asking whether it could be worth using your tax return to buy a mobile phone.\r\nMy wife and I are on the same phone plan and our iPhone 7s are getting a bit long in the tooth, this reader wrote.… [+3882 chars]",
        Description: "Mere weeks ago, in the Before Times, I got an email from a reader asking whether it could be worth using your tax return to buy a mobile phone.Read more...",
        PublishedAt: "2020-04-22T20:30:00Z",
        TagName: "Positive",
        Title: "When Does It Make Sense to Buy a Mobile Phone Outright?",
        Url: "https://twocents.lifehacker.com/when-does-it-make-sense-to-buy-a-mobile-phone-outright-1843006230",
        UrlToImage: "https://i.kinja-img.com/gawker-media/image/upload/c_fill,f_auto,fl_progressive,g_center,h_675,pg_1,q_80,w_1200/icdcua1egvnwqyymfuyg.jpg",
	}
	article2 := Article{
        Author: "Test author",
        Confidence: 0.98,
        Content: "Test content",
        Description: "article description",
        PublishedAt: "2020-04-22T20:30:00Z",
        TagName: "Neutral",
        Title: "Article title",
        Url: "url",
        UrlToImage: "image",
	}
	var articles []Article
	articles = append(articles, article1)
	articles = append(articles, article2)
	return articles
}

func getTitles(articles []Article) []string {
	var titles []string
	for j := 0;  j < 2; j++ {
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
	for j := 0;  j < 2; j++ {
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
