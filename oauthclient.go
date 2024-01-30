package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	oauthURL = "https://auth.us-east-2.propeldata.com/oauth2/token"
)

type PropelOAuthClient interface {
	OAuthToken(ctx context.Context, applicationID, applicationSecret string) (*OAuthToken, error)
}

type OauthClient struct {
	client *http.Client
}

type OAuthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

var _ PropelOAuthClient = (*OauthClient)(nil)

func NewOauthClient() PropelOAuthClient {
	client := newHttpClient(Options{
		Timeout: 2 * time.Second,
		Retries: 3,
		Delay:   5 * time.Millisecond,
	})

	return &OauthClient{client: client}
}

func (c *OauthClient) OAuthToken(ctx context.Context, clientID string, clientSecret string) (*OAuthToken, error) {
	payload := url.Values{}
	payload.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, oauthURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, clientSecret)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch access token; status=%d", resp.StatusCode)
	}

	var token *OAuthToken

	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return token, nil
}
