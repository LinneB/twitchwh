package twitchwh

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"
)

// List of request headers sent from Twitch
// See: https://dev.twitch.tv/docs/eventsub/handling-webhook-events/#list-of-request-headers
const twitchMessageID = "Twitch-Eventsub-Message-Id"
const twitchMessageTimestamp = "Twitch-Eventsub-Message-Timestamp"
const twitchMessageSignature = "Twitch-Eventsub-Message-Signature"
const messageType = "Twitch-Eventsub-Message-Type"

// Message types
const messageTypeNotification = "notification"
const messageTypeVerification = "webhook_callback_verification"
const messageTypeRevocation = "revocation"

type webhookPayload struct {
	Challenge    string          `json:"challenge"`
	Subscription Subscription    `json:"subscription"`
	Event        json.RawMessage `json:"event"`
}

// Handler is the HTTP handler for requests from Twitch.
// It is up to you to assign this handler to the correct path according to your setup
//
//	client, _ := twitchwh.New(twitchwh.ClientConfig{
//		// ...
//		WebhookURL:    "https://mydomain.com/eventsub",
//	})
//	http.HandleFunc("/eventsub", client.Handler)
//	http.ListenAndServe(":443", nil)
//
// This example assumes https://mydomain.com is pointing to the Go app.
func (c *Client) Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		c.logger.Printf("Could not read request body: %s", err)
		w.WriteHeader(500)
		return
	}

	hmacMessage := r.Header.Get(twitchMessageID) + r.Header.Get(twitchMessageTimestamp) + string(body)
	expectedSignature := "sha256=" + generateHmac(c.webhookSecret, hmacMessage)
	if verifyHmac(expectedSignature, r.Header.Get(twitchMessageSignature)) {
		c.logger.Println("Received valid signature")

		var payload webhookPayload
		err := json.Unmarshal(body, &payload)
		if err != nil {
			c.logger.Printf("Could not serialize webhook payload: %s", err)
			w.WriteHeader(500)
			return
		}

		message_type := r.Header.Get(messageType)
		if message_type == messageTypeNotification {
			c.logger.Printf("Received event for %s ", payload.Subscription.Type)
			if slices.Contains(c.handledEvents, r.Header.Get(twitchMessageID)) {
				c.logger.Println("Got request for handled event, ignoring...")
				w.WriteHeader(204)
				return
			} else {
				c.handledEvents = append(c.handledEvents, r.Header.Get(twitchMessageID))
			}

			if handler, ok := c.handlers[payload.Subscription.Type]; ok {
				go handler(payload.Event)
			} else {
				c.logger.Printf("No handler for event %s", payload.Subscription.Type)
			}

			w.WriteHeader(204)
			return
		}
		if message_type == messageTypeVerification {
			c.logger.Printf("Got challenge request for %s", payload.Subscription.ID)
			go func() {
				c.verifiedSubscriptions <- payload.Subscription.ID
			}()
			w.WriteHeader(200)
			w.Write([]byte(payload.Challenge))
			return
		}
		if message_type == messageTypeRevocation {
			// Subscription was revoked. This could be as simple as a user deactivating or Twitch not reaching the endpoint.
			c.logger.Printf("Twitch revoked subscription %s", payload.Subscription.ID)
			if c.OnRevocation != nil {
				c.OnRevocation(payload.Subscription)
			}
			w.WriteHeader(204)
			return
		}
	} else {
		w.WriteHeader(403)
	}
}
