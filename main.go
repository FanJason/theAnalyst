package main

import (
	"net/http"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql-go-handler"
)

type Article struct {
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

func getArticles() []Article {
	article1 := Article{
		Author: "Jason",
		Title: "Test Article 1",
		Description: "Test description",
		Url: "url",
		UrlToImage: "image link",
		PublishedAt: "04/30/2020",
		Content: "hello world",
	}
	article2 := Article{
		Author: "Jason",
		Title: "Test Article 2",
		Description: "Test description",
		Url: "url",
		UrlToImage: "image link",
		PublishedAt: "04/30/2020",
		Content: "hello world - part 2",
	}
	var articles []Article
	articles = append(articles, article1)
	articles = append(articles, article2)
	return articles
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