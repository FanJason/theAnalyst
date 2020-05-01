package sentiment

import (
	"github.com/FanJason/theAnalyst/server/env"
	"github.com/FanJason/theAnalyst/server/request"
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

func AnalyzeSentiment(content string) Classification {
	url := "https://api.monkeylearn.com/v3/classifiers/cl_pi3C7JiL/classify/"
	var data []Response
	request.Post(url, content, env.GetEnvVariable("SENTIMENT"), &data)
	return data[0].Classifications[0]
}