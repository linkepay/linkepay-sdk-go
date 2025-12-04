package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const DefaultTimeout = 30 * time.Second

// RequestConfig represents configuration for an HTTP request
type RequestConfig struct {
	BaseURL string
	Method  string
	Path    string
	Params  url.Values
	Body    interface{}
	Headers map[string]string
	Timeout int // Timeout in seconds (0 uses DefaultTimeout)
}

// Request makes an HTTP request and returns the response body
func Request(rc RequestConfig) ([]byte, error) {
	// Determine timeout
	timeout := DefaultTimeout
	if rc.Timeout > 0 {
		timeout = time.Duration(rc.Timeout) * time.Second
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

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

	req, err := http.NewRequestWithContext(ctx, rc.Method, url, bytes.NewBuffer(reqBody))
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
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("request timeout: request to %s exceeded %v", rc.Path, timeout)
		}
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	fmt.Printf("[LinkePay Debug] Response Status: %d\n", resp.StatusCode)
	fmt.Printf("[LinkePay Debug] Response Body (first 200 chars): %s\n", string(body[:min(200, len(body))]))

	return body, nil
}
