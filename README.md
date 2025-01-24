# TwitchWH

TwitchWH is a Twitch Webhook EventSub library for Go.

Full documentation: https://pkg.go.dev/github.com/LinneB/twitchwh

Table of contents:

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Contributing](#contributing)
- [Supported Events](#supported-events)

## Installation

```bash
go get github.com/LinneB/twitchwh
```

## Basic Usage

This is a basic example of how to use TwitchWH. It creates a new client, adds an event listener, and subscribes to an event.

For the full documentation, see [pkg.go.dev](https://pkg.go.dev/github.com/LinneB/twitchwh).

```go
package main

import (
	"log"
	"net/http"

	"github.com/LinneB/twitchwh"
)

func main() {
	// Create a new Client
	client, err := twitchwh.New(twitchwh.ClientConfig{
		ClientID:      "client id",
		ClientSecret:  "super secret client secret",
		WebhookSecret: "random string between 10 and 100 characters",
		// This example assumes you have a domain that points to this app on port 8080
		WebhookURL:    "https://yourdomain.com/eventsub",
		// Enable log output
		Debug: true,
	})
	if err != nil {
		log.Panic(err)
	}

	// Set up an event handler
	client.On("stream.online", func(event json.RawMessage) {
		var eventBody struct {
			BroadcasterUserLogin string `json:"broadcaster_user_login"`
		}
		json.Unmarshal(event, &eventBody)
		log.Printf("%s went live!", eventBody.BroadcasterUserLogin)
	})

	// Setup the HTTP event handler
	// This needs to be done before you subscribe to any events, since AddSubscription will wait until Twitch sends the challenge request
	http.HandleFunc("/eventsub", client.Handler)
	go http.ListenAndServe(":8080", nil)

	// Add a subscription for LinneB going live
	// Note that this will throw an error if the subscription already exists
	err = client.AddSubscription("stream.online", "1", twitchwh.Condition{
		BroadcasterUserID: "215185844",
	})
	if err != nil {
		log.Panic(err)
	}

	// Wait forever
	select {}
}
```

## Contributing

Contributions are welcome. If you find any issues or have any suggestions, please open an issue or a pull request.

Questions and feature requests are also welcome, just open an issue.

## Supported Events

TwitchWH should theoretically support all current and future EventSub events, as long as the Condition struct has the required fields. If you find an event that is not supported, don't hesitate to open an issue.
