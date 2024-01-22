package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type WebhookClient struct {
	client *http.Client
}

func NewWebhookClient() *WebhookClient {
	client := newHttpClient(Options{
		Timeout: 10 * time.Second,
		Retries: 3,
		Delay:   5 * time.Millisecond,
	})

	return &WebhookClient{client: client}
}

type PostEventsInput struct {
	WebhookURL   string
	AuthUsername string
	AuthPassword string
	Events       []map[string]any
}

type PostEventResponse struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

func (c *WebhookClient) PostEvents(ctx context.Context, input *PostEventsInput) ([]error, error) {
	return c.publishEvents(ctx, http.MethodPost, input)
}

func (c *WebhookClient) DeleteEvents(ctx context.Context, input *PostEventsInput) ([]error, error) {
	return c.publishEvents(ctx, http.MethodDelete, input)
}

func (c *WebhookClient) publishEvents(ctx context.Context, method string, input *PostEventsInput) ([]error, error) {
	body, err := json.Marshal(input.Events)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal events to JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, input.WebhookURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(input.AuthUsername, input.AuthPassword)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to publish events to %s: %w", input.WebhookURL, err)
	}

	var eventsResponse []PostEventResponse
	if err := json.NewDecoder(resp.Body).Decode(&eventsResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response from %s: %w", input.WebhookURL, err)
	}

	eventErrors := make([]error, len(input.Events))
	errorCount := 0

	for i, response := range eventsResponse {
		if response.StatusCode != http.StatusOK {
			eventErrors[i] = errors.New(response.StatusMessage)
			errorCount++
		}
	}

	if errorCount > 0 {
		return eventErrors, nil
	}

	return nil, nil
}
