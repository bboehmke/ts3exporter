package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TS3HttpClient struct {
	baseUrl string
	apiKey string

	client *http.Client
}

func NewClient(baseUrl string, apiKey string) *TS3HttpClient {
	return &TS3HttpClient{
		baseUrl: baseUrl,
		apiKey:  apiKey,
		client:  new(http.Client),
	}
}

func (c *TS3HttpClient) request(path string, data interface{}) error {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.baseUrl, path), nil)
	req.Header.Add("x-api-key", c.apiKey)
	response, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("error Code: %d\n%s", response.StatusCode, string(body))
	}

	var tsResp struct {
		Body   interface{} `json:"body"`
		Status struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"status"`
	}

	err = json.NewDecoder(response.Body).Decode(&tsResp)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if tsResp.Status.Code != 0 {
		return errors.New(tsResp.Status.Message)
	}

	if data == nil {
		return nil
	}

	jsonBody, err := json.Marshal(tsResp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	err = json.Unmarshal(jsonBody, &data)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}
	return nil
}
