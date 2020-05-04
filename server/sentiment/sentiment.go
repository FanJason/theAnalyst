package sentiment

import (
	// "github.com/FanJason/theAnalyst/server/env"
	// "github.com/FanJason/theAnalyst/server/request"
)

type Classification struct {
	Tag_Name           string
	Tag_ID             int
	Confidence         float64
}

type Response struct {
	Text               string
	External_ID        int
	Error              bool
	Classifications    []Classification
}

func AnalyzeSentiment(content string) []Response {
	// url := "https://api.monkeylearn.com/v3/classifiers/cl_pi3C7JiL/classify/"
	// request.Post(url, content, env.GetEnvVariable("SENTIMENT"), &data)
	// return data
	classification := Classification{
			Tag_Name: "Positive",
			Tag_ID: 123124,
			Confidence: 0.94,
		}

	var classifications []Classification
	classifications = append(classifications, classification)
	response1 := Response{
		Text: "test 1",
		External_ID: 1,
		Error: false,
		Classifications: classifications,
	}
	var data []Response
	data = append(data, response1)
	return data
}