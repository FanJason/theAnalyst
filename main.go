package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"os"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
	"github.com/joho/godotenv"
)

type Response struct {
	Status       string
	TotalResults int
	Articles     []Article
}

type Source struct {
	ID           int
	Name         string
}

type Article struct {
	Source       Source
	Author       string
	Title        string
	Description  string
	Url          string
	UrlToImage   string
	PublishedAt  string
	Content      string
}

type User struct {
	Username    string
	Password    string
}

type Comment struct {
	Body        string
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("failed to get env variable: %v", err)
	}
	return os.Getenv(key)
}

func getArticles() []Article {
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

var articleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"Author": &graphql.Field{
				Type: graphql.String,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Description": &graphql.Field{
				Type: graphql.String,
			},
			"Url": &graphql.Field{
				Type: graphql.String,
			},
			"UrlToImage": &graphql.Field{
				Type: graphql.String,
			},
			"PublishedAt": &graphql.Field{
				Type: graphql.String,
			},
			"Content": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func main() {
	articles := getArticles()

	fields := graphql.Fields {
		"article": &graphql.Field{
			Type: articleType,
			Description: "Get Article by Title",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				title, ok := p.Args["title"].(string)
				if ok {
					for _, article := range articles {
						if string(article.Title) == title {
							return article, nil
						}
					}
				}
				return nil, nil
			},
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(articleType),
			Description: "Get Article List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return articles, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{ Name: "RootQuery", Fields: fields }
	schemaConfig := graphql.SchemaConfig{ Query: graphql.NewObject(rootQuery) }
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	http.Handle("/graphql", h)
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":8080", nil)
}