package translation

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

var _ HelloClient = &APIClient{}

type APIClient struct {
	endpoint string
}

// NewHelloClient creates instance of client with a given endpoint.
func NewHelloClient(endpoint string) *APIClient {
	return &APIClient{
		endpoint: endpoint,
	}
}

// Translate will call external client for translation.
func (c *APIClient) Translate(word, language string) (string, error) {
	req := map[string]string{
		"word":     word,
		"language": language,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return "", errors.New("unable to encode msg")
	}

	resp, err := http.Post(c.endpoint, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return "", errors.New("call to api failed")
	}
	if resp.StatusCode == http.StatusNotFound {
		return "", nil
	}
	if resp.StatusCode == http.StatusInternalServerError {
		return "", errors.New("error in api")
	}
	b, _ = io.ReadAll(resp.Body)
	defer func(r *http.Response) {
		_ = resp.Body.Close()
	}(resp)

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return "", errors.New("unable to decode message")
	}
	return m["translation"].(string), nil
}
