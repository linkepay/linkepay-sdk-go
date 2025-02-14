package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Endpoint represents an API endpoint to document
type RequestConfig struct {
	BaseURL string
	Method  string
	Path    string
	Params  url.Values
	Body    interface{}
	Headers map[string]string
}

// Request makes an HTTP request and returns the response body
func Request(rc RequestConfig) ([]byte, error) {
	url := rc.BaseURL + rc.Path
	if len(rc.Params) > 0 {
		url += "?" + rc.Params.Encode()
	}

	var reqBody []byte
	var err error
	if rc.Body != nil {
		reqBody, err = json.Marshal(rc.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %v", err)
		}
	}

	req, err := http.NewRequest(rc.Method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for k, v := range rc.Headers {
		req.Header.Set(k, v)
	}

	if rc.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	return body, nil
}
