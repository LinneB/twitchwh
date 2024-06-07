# TwitchWH

TwitchWH is a Twitch Webhook EventSub library for Go.

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

For more examples and methods, see [USAGE.md](USAGE.md).

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
	client.OnStreamOnline = func(event twitchwh.StreamOnline) {
		log.Printf("%s went live!", event.BroadcasterUserLogin)
	}

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

TwitchWH supports all non-beta events as of `2024-06-07`. For a complete list of events, see [EventSub subscription types](https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/).

Please note that most of these events are not tested, and are only modeled based on the documentation.
Double check the documentation and make sure the struct fields match.

Full list of supported events:
- `automod.message.hold`
- `automod.message.update`
- `automod.settings.update`
- `automod.terms.update`
- `channel.update`
- `channel.follow`
- `channel.ad_break.begin`
- `channel.chat.clear`
- `channel.chat.clear_user_messages`
- `channel.chat.message`
- `channel.chat.message_delete`
- `channel.chat.notification`
- `channel.chat_settings.update`
- `channel.chat.user_message_hold`
- `channel.chat.user_message_update`
- `channel.subscribe`
- `channel.subscription.end`
- `channel.subscription.gift`
- `channel.subscription.message`
- `channel.cheer`
- `channel.raid`
- `channel.ban`
- `channel.unban`
- `channel.unban_request.create`
- `channel.unban_request.resolve`
- `channel.moderate`
- `channel.moderator.add`
- `channel.moderator.remove`
- `channel.channel_points_automatic_reward_redemption.add`
- `channel.channel_points_custom_reward.add`
- `channel.channel_points_custom_reward.update`
- `channel.channel_points_custom_reward.remove`
- `channel.channel_points_custom_reward_redemption.add`
- `channel.channel_points_custom_reward_redemption.update`
- `channel.poll.begin`
- `channel.poll.progress`
- `channel.poll.end`
- `channel.prediction.begin`
- `channel.prediction.progress`
- `channel.prediction.lock`
- `channel.prediction.end`
- `channel.suspicious_user.message`
- `channel.suspicious_user.update`
- `channel.vip.add`
- `channel.vip.remove`
- `channel.charity_campaign.donate`
- `channel.charity_campaign.start`
- `channel.charity_campaign.progress`
- `channel.charity_campaign.stop`
- `conduit.shard.disabled`
- `drop.entitlement.grant`
- `extension.bits_transaction.create`
- `channel.goal.begin`
- `channel.goal.progress`
- `channel.goal.end`
- `channel.hype_train.begin`
- `channel.hype_train.progress`
- `channel.hype_train.end`
- `channel.shield_mode.begin`
- `channel.shield_mode.end`
- `channel.shoutout.create`
- `channel.shoutout.receive`
- `stream.online`
- `stream.offline`
- `user.authorization.grant`
- `user.authorization.revoke`
- `user.update`
- `user.whisper.message`
