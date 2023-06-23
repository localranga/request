package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func FetchDataFromAPI(api main.API) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", api.URL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range api.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
