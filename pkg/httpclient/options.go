package httpclient

import "time"

// Option -.
type Option func(*Client)

// Timeout -.
func Timeout(t time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = t
	}
}
