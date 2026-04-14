package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"time"
)

const DefaultTimeout = 30 * time.Second

// fetchPublicIP queries an external echo service to discover the public
// outbound IP this process is using. This is the IP LinkePay's whitelist
// is matched against, so logging it makes "IP not whitelisted" errors
// debuggable without shelling into the host.
func fetchPublicIP() string {
	c := &http.Client{Timeout: 5 * time.Second}
	resp, err := c.Get("https://api.ipify.org")
	if err != nil {
		return fmt.Sprintf("(lookup failed: %v)", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("(read failed: %v)", err)
	}
	return string(b)
}

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

	var localAddr string
	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			if info.Conn != nil && info.Conn.LocalAddr() != nil {
				localAddr = info.Conn.LocalAddr().String()
			}
		},
	}
	ctx = httptrace.WithClientTrace(ctx, trace)

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
		publicIP := fetchPublicIP()
		fmt.Printf("[LinkePay Debug] Request failed status=%d url=%s\n", resp.StatusCode, url)
		fmt.Printf("[LinkePay Debug] Local socket addr: %s\n", localAddr)
		fmt.Printf("[LinkePay Debug] Public outbound IP: %s\n", publicIP)
		fmt.Printf("[LinkePay Debug] Response body: %s\n", string(body))
		return nil, fmt.Errorf("HTTP %d: %s (outbound_ip=%s local_addr=%s)", resp.StatusCode, string(body), publicIP, localAddr)
	}

	fmt.Printf("[LinkePay Debug] Response Status: %d (local_addr=%s)\n", resp.StatusCode, localAddr)
	fmt.Printf("[LinkePay Debug] Response Body (first 200 chars): %s\n", string(body[:min(200, len(body))]))

	return body, nil
}
