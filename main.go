package main

import (
	"log"
	"net/http"

	"github.com/FanJason/theAnalyst/articles"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql-go-handler"
)

type User struct {
	Username    string
	Password    string
}

type Comment struct {
	Body        string
}

var sourceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Source",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var articleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"Source": &graphql.Field{
				Type: sourceType,
			},
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

func getFields() graphql.Fields {
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
				articles := articles.GetArticles(3)
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
				articles := articles.GetArticles(3)
				return articles, nil
			},
		},
	}
	return fields
}

func getSchema() graphql.Schema {
	rootQuery := graphql.ObjectConfig{ Name: "RootQuery", Fields: getFields() }
	schemaConfig := graphql.SchemaConfig{ Query: graphql.NewObject(rootQuery) }
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}

func main() {
	schema := getSchema()

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