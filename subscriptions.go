package twitchwh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Condition for subscription. Empty values will be omitted. Fill out the options applicable to your subscription type
type Condition struct {
	// broadcaster_user_id
	BroadcasterUserID string `json:"broadcaster_user_id,omitempty"`

	// moderator_user_id
	ModeratorUserID string `json:"moderator_user_id,omitempty"`

	// user_id
	UserID string `json:"user_id,omitempty"`

	// from_broadcaster_id
	FromBroadcasterUserID string `json:"from_broadcaster_user_id,omitempty"`

	// to_broadcaster_id
	ToBroadcasterUserID string `json:"to_broadcaster_user_id,omitempty"`

	// reward_id
	//
	// This should be int/string depending on subscription type
	RewardID any `json:"reward_id,omitempty"`

	// client_id
	ClientID string `json:"client_id,omitempty"`

	// extension_client_id
	ExtensionClientID string `json:"extension_client_id,omitempty"`

	// conduit_id
	ConduitID string `json:"conduit_id,omitempty"`

	// organization_id
	OrganizationID string `json:"organization_id,omitempty"`

	// category_id
	CategoryID string `json:"category_id,omitempty"`

	// campaign_id
	CampaignID string `json:"campaign_id,omitempty"`
}

type Subscription struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Version string `json:"version"`
	Cost    int    `json:"cost"`
	// PLEASE NOTE that this will DEFAULT all unused conditions. Check the Type and get the correct condition for that type.
	Condition Condition `json:"condition"`
	Transport struct {
		Method   string `json:"method"`
		Callback string `json:"callback"`
	} `json:"transport"`
	CreatedAt time.Time `json:"created_at"`
}

type transport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

type subscriptionRequest struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Condition Condition `json:"condition"`
	Transport transport `json:"transport"`
}

// AddSubscription attemps to create a new subscription based on the type, version, and condition.
// You can find all subscription types, versions, and conditions at: [EventSub subscription types].
// It will block until Twitch sends the verification request, or timeout after 10 seconds.
//
// !! AddSubscription should only be called AFTER [twitchwh.Client.Handler] is set up accordingly. !!
//
//	// Setup the HTTP event handler
//	http.HandleFunc("/eventsub", client.Handler)
//	go http.ListenAndServe(":8080", nil)
//
//	_ := client.AddSubscription("stream.online", "1", twitchwh.Condition{
//		BroadcasterUserID: "215185844",
//	})
//
// [EventSub subscription types]: https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/
func (c *Client) AddSubscription(Type string, version string, condition Condition) error {
	reqBody, err := json.Marshal(subscriptionRequest{
		Type:      Type,
		Version:   version,
		Condition: condition,
		Transport: transport{
			Method:   "webhook",
			Callback: c.webhookURL,
			Secret:   c.webhookSecret,
		},
	})
	if err != nil {
		return fmt.Errorf("Could not marshal request body: %w", err)
	}

	request, err := http.NewRequest("POST", helixURL+"/eventsub/subscriptions", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("Could not create request: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Client-ID", c.clientID)
	request.Header.Set("Authorization", "Bearer "+c.token)

	res, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("Could not send request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Could not read response body: %w", err)
	}

	if res.StatusCode != 202 {
		defer res.Body.Close()
		c.logger.Println(string(body))
		return fmt.Errorf("Twitch returned unhandled status code %d", res.StatusCode)
	}

	var responseBody struct {
		Data []Subscription `json:"data"`
	}

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return fmt.Errorf("Could not parse response body: %w", err)
	}

	// Returned body is an array that contains a single subscription
	subscription := responseBody.Data[0]

	// Await confirmation
	for {
		select {
		case id := <-c.verifiedSubscriptions:
			if id == subscription.ID {
				c.logger.Printf("Subscription created: %s", subscription.ID)
				return nil
			} else {
				// Verified subscription was not for this subscription
				c.logger.Println("Subscription confirmation did not match ID, ignoring...")
				c.verifiedSubscriptions <- id
				continue
			}
		case <-time.After(10 * time.Second):
			return fmt.Errorf("Never received confirmation of subscription: %s", subscription.ID)
		}
	}
}

// RemoveSubscription attempts to remove a subscription based on the ID.
func (c *Client) RemoveSubscription(id string) error {
	url := "/eventsub/subscriptions?id=" + id
	res, err := c.genericRequest("DELETE", url)
	if err != nil {
		return fmt.Errorf("Could not create request: %w", err)
	}
	if res.StatusCode == 204 {
		return nil
	}
	if res.StatusCode == 404 {
		return fmt.Errorf("Subscription not found")
	}
	return fmt.Errorf("Unhandled status code %d", res.StatusCode)
}

// RemoveSubscriptionByType attempts to remove a subscription based on the type and condition.
//
// Note: This will remove ALL subscriptions that match the provided type and condition.
func (c *Client) RemoveSubscriptionByType(Type string, condition Condition) error {
	subs, err := c.GetSubscriptionsByType(Type)
	if err != nil {
		return err
	}
	for _, sub := range subs {
        // Both of these conditions have unused fields, but since they are both defaulted and of the same type it should be fine
		if sub.Condition == condition {
			c.logger.Printf("Removing subscription %s", sub.ID)
			err := c.RemoveSubscription(sub.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Internal function to fetch subscriptions using the provided URL parameters.
// Used by wrapper functions.
// Automatically handles pagination.
func (c *Client) fetchSubscriptions(urlParams string) (subscriptions []Subscription, err error) {
	page := 1
	cursor := ""
	for {
		c.logger.Printf("Fetching page %d of subscriptions", page)
		page++

		var params string
		if cursor != "" {
			if urlParams == "" {
				params = urlParams + "?after=" + cursor
			} else {
				params = "&after=" + cursor
			}
		}
		res, err := c.genericRequest("GET", "/eventsub/subscriptions"+params)
		if err != nil {
			return nil, fmt.Errorf("Could make request: %w", err)
		}

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Helix returned unhandled status code: %d", res.StatusCode)
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Could not read response body: %w", err)
		}
		var responseStruct struct {
			Data       []Subscription `json:"data"`
			Pagination struct {
				Cursor string `json:"cursor"`
			} `json:"pagination"`
		}
		err = json.Unmarshal(body, &responseStruct)
		if err != nil {
			return nil, fmt.Errorf("Could not parse response body: %w", err)
		}

		subscriptions = append(subscriptions, responseStruct.Data...)

		if responseStruct.Pagination.Cursor == "" {
			// No more subscriptions to fetch
			break
		}
		cursor = responseStruct.Pagination.Cursor
	}
	return subscriptions, nil
}

// GetSubscriptions retrieves all subscriptions, including revoked ones.
// Automatically handles pagination.
//
// Returns subscriptions and an error (if any).
func (c *Client) GetSubscriptions() (subscriptions []Subscription, err error) {
	urlParams := ""
	return c.fetchSubscriptions(urlParams)
}

// Get all subscriptions that match the provided type (eg. "stream.online").
// Automatically handles pagination.
//
// Returns subscriptions and an error (if any).
func (c *Client) GetSubscriptionsByType(Type string) (subscriptions []Subscription, err error) {
	urlParams := "?type=" + Type
	return c.fetchSubscriptions(urlParams)
}

// Get all subscriptions with the provided status.
// For a list of all status types see: https://dev.twitch.tv/docs/api/reference/#get-eventsub-subscriptions .
// Automatically handles pagination.
//
// Returns subscriptions and an error (if any).
func (c *Client) GetSubscriptionsByStatus(status string) (subscriptions []Subscription, err error) {
	urlParams := "?status=" + status
	return c.fetchSubscriptions(urlParams)
}
