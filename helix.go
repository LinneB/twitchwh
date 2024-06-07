package twitchwh

import "net/http"

const helixURL = "https://api.twitch.tv/helix"

// Interal generic request function that includes authorization headers.
// TODO: Should this return the request rather than the response?
func (c *Client) genericRequest(method string, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(method, helixURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Client-ID", c.clientID)

	return c.httpClient.Do(req)
}
