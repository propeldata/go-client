package client

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type Options struct {
	// Transport sets the http.RoundTripper to use.
	Transport http.RoundTripper

	// Timeout sets a per-request timeout, which applies to each attempt. Defaults to "no timeout". For enforcing a
	// per-operation timeout, which applies to all attempts, use context.
	Timeout time.Duration

	// Retries is the maximum number of retries for an operation. Defaults to 2 retries (or 3 attempts, total). Setting
	// this to 0 means 1 attempt, 0 retries.
	Retries int

	// Delay sets the initial delay between attempts. Defaults to 0 or "no delay".
	Delay time.Duration

	// MaxDelay sets the maximum delay between attempts. Setting this allows implementing capped exponential backoff.
	// Defaults to "no limit".
	MaxDelay time.Duration
}

// newHttpClient returns a new HTTP client implementing capped, exponential backoff. Typical retryable HTTP status codes
// will be retried, including 500 and 429.
func newHttpClient(opts Options) *http.Client {
	client := retryablehttp.NewClient()
	client.HTTPClient.Transport = opts.Transport
	client.HTTPClient.Timeout = opts.Timeout
	client.RetryMax = opts.Retries
	client.RetryWaitMin = opts.Delay
	client.RetryWaitMax = opts.MaxDelay

	if client.RetryWaitMax == 0 {
		client.RetryWaitMax = 5 * time.Second
	}

	client.ErrorHandler = retryablehttp.PassthroughErrorHandler

	return client.StandardClient()
}
