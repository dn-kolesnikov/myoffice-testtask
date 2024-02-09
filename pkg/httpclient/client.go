package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	_defaultTimeout = 5 * time.Second
)

// Client -.
type Client struct {
	client *http.Client
}

// New -.
func New(opts ...Option) *Client {
	c := &Client{
		client: &http.Client{
			Timeout: _defaultTimeout,
		},
	}
	// Custom options -.
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Get -.
func (c Client) Get(URL string) (r Response, err error) {
	startTimer := time.Now()
	resp, err := c.client.Get(URL)
	if err != nil {
		return r, fmt.Errorf("failed to get response from %s: %w", URL, err)
	}
	stopTimer := time.Since(startTimer)

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = fmt.Errorf("failed to close response: %w", cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return r, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, fmt.Errorf("failed to read response body: %w", err)
	}

	return Response{
		handleTime:    stopTimer,
		contentLength: len(body),
	}, nil
}

// ValidateURL -.
func (c Client) ValidateURL(URL string) error {
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return fmt.Errorf("failed to parse url: %w", err)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("invalid scheme: %s", u.Scheme)
	}

	return nil
}

// Response -.
type Response struct {
	handleTime    time.Duration
	contentLength int
}

// ContentLength -.
func (r Response) ContentLength() int {
	return r.contentLength
}

// HandleTime -.
func (r Response) HandleTime() time.Duration {
	return r.handleTime
}
