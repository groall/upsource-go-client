package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Error represents an Upsource RPC error envelope
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Client is a minimal Upsource RPC client
type Client struct {
	baseURL    string
	httpClient *http.Client
	authHeader string
}

// Options configure the Client
type Options struct {
	// BaseURL is like https://your-upsource-host
	BaseURL string
	// Username and Password are used for Basic auth; if empty, requests are unauthenticated (guest)
	Username string
	Password string
	// HTTPClient allows providing a custom http.Client; if nil a default is used
	HTTPClient *http.Client
	// Timeout overrides HTTP client timeout; ignored if HTTPClient provided
	Timeout time.Duration
}

// New creates a new Client instance
func New(opts Options) (*Client, error) {
	if opts.BaseURL == "" {
		return nil, errors.New("BaseURL is required")
	}
	base := strings.TrimRight(opts.BaseURL, "/")
	hc := opts.HTTPClient
	if hc == nil {
		timeout := opts.Timeout
		if timeout <= 0 {
			timeout = 30 * time.Second
		}
		hc = &http.Client{Timeout: timeout}
	}
	var auth string
	if opts.Username != "" || opts.Password != "" {
		token := base64.StdEncoding.EncodeToString([]byte(opts.Username + ":" + opts.Password))
		auth = "Basic " + token
	}
	return &Client{baseURL: base, httpClient: hc, authHeader: auth}, nil
}

// Call performs an RPC request to /~rpc/{method}.
// params is marshaled as JSON. If params is nil, an empty object is sent.
// result must be a pointer to a value to receive the unwrapped "result".
func (c *Client) Call(ctx context.Context, method string, params any, result any) error {
	if method == "" {
		return errors.New("method is required")
	}

	// Marshal params as JSON
	if params == nil {
		params = struct{}{}
	}
	bodyBytes, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("marshal params: %w", err)
	}

	endpoint := c.baseURL + "/~rpc/" + url.PathEscape(method)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(string(bodyBytes)))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if c.authHeader != "" {
		req.Header.Set("Authorization", c.authHeader)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	// Decode into envelope with generic type that matches result
	// We must decode into a temporary map to keep generic flexibility
	var envelope struct {
		Result json.RawMessage `json:"result"`
		Error  *Error          `json:"error"`
	}
	if err := json.Unmarshal(respBody, &envelope); err != nil {
		return fmt.Errorf("decode envelope: %w; body=%s", err, string(respBody))
	}
	if envelope.Error != nil {
		return fmt.Errorf("upsource error %d: %s", envelope.Error.Code, envelope.Error.Message)
	}
	if result == nil {
		return nil
	}
	if len(envelope.Result) == 0 {
		// Some endpoints can return empty result
		return nil
	}
	if err := json.Unmarshal(envelope.Result, result); err != nil {
		return fmt.Errorf("decode result: %w; body=%s", err, string(respBody))
	}
	return nil
}
