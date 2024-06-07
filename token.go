package twitchwh

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const oauthURL = "https://id.twitch.tv/oauth2"
const tokenURL = oauthURL + "/token"
const validateURL = oauthURL + "/validate"

func (c *Client) generateToken(clientID string, secret string) (token string, err error) {
	values := url.Values{
		"client_id":     {clientID},
		"client_secret": {secret},
		"grant_type":    {"client_credentials"},
	}

	res, err := c.httpClient.PostForm(tokenURL, values)
	if err != nil {
		return "", fmt.Errorf("Could not send request: %w", err)
	}

	if res.StatusCode == 401 {
		return "", fmt.Errorf("ClientID or Client secret invalid")
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("Token generation failed with status: %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Could not read response body: %w", err)
	}

	var jsonBody struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return "", fmt.Errorf("Could not serialize token response body: %w", err)
	}

	return jsonBody.AccessToken, nil
}

func (c *Client) validateToken(token string) (bool, error) {
	req, err := http.NewRequest("GET", validateURL, nil)
	if err != nil {
		return false, fmt.Errorf("Could not create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("Could not send request: %w", err)
	}

	if res.StatusCode == 200 {
		return true, nil
	} else {
		return false, nil
	}
}
