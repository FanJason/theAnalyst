package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"

	"github.com/FanJason/theAnalyst/server/articles"
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

type CacheEntry struct {
	date        string
	articles    []articles.Article
}

// key: topic, value: CacheEntry
var cachedArticles = make(map[string]CacheEntry)

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
			"TagName": &graphql.Field{
				Type: graphql.String,
			},
			"Confidence": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

func compareDates(currentDate string, cachedDate string) bool {
	currentDay, err1 := strconv.Atoi(currentDate[3:5])
	cachedDay, err2 := strconv.Atoi(cachedDate[3:5])
	if err1 != nil || err2 != nil {
		fmt.Printf("failed to parse date: %v, $v", err1, err2);
	}
	return currentDay - cachedDay < 1
}

func addToCache(topic string) {
	dt := time.Now()
	cacheEntry := CacheEntry{
		date: dt.Format("01-01-2020"),
		articles: articles.GetArticles(topic),
	}
	cachedArticles[topic] = cacheEntry
}

func isCached(topic string) bool {
	dt := time.Now()
	if len(cachedArticles[topic].articles) > 0 {
		return compareDates(dt.Format("01-01-2020"), cachedArticles[topic].date)
	}
	return false
}

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
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				title, ok := params.Args["title"].(string)
				articles := articles.GetArticles("finance")
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
			Args: graphql.FieldConfigArgument{
				"topic": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				topic, ok := params.Args["topic"].(string)
				if ok {
					if !isCached(topic) {
						addToCache(topic)
					}
					return cachedArticles[topic].articles, nil
				}
				return nil, nil
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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
        next.ServeHTTP(w,r)
    })
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

	http.Handle("/graphql", corsMiddleware(h))
	http.Handle("/graphiql", graphiqlHandler)
	http.ListenAndServe(":8080", nil)
}