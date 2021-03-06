package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Get(url string, data interface{}) {
	client := &http.Client {}
	request, err := http.NewRequest("GET", url, strings.NewReader("{}"))
  
	if err != nil {
	  fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(data)

	if err != nil {
		fmt.Printf("failed to parse api response: %v", err)
	}
}

func Post(url string, content string, apiKey string, data interface{}) {
	method := "POST"
	payload := strings.NewReader(content)
  
	client := &http.Client{}
	request, err := http.NewRequest(method, url, payload)
  
	if err != nil {
	  fmt.Printf("failed to create request: %v", err)
	}

	request.Header.Add("Authorization", "Token " + apiKey)
	request.Header.Add("Content-Type", "application/json")
  
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("failed to send POST request: %v", err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		fmt.Printf("failed to decode POST response: %v", err)
	}
}