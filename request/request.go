package request

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func Get(url string, data interface{}) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get api data, error: %v", err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(data)

	if err != nil {
		fmt.Printf("failed to parse api response: %v", err)
	}
}