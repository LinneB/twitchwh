# Usage

This file contains examples and methods for using TwitchWH.

- [Configuration](#configuration)
- [Events](#events)
- [Client methods](#client-methods)
    - [AddSubscription](#addsubscription)
    - [RemoveSubscription](#removesubscription)
    - [RemoveSubscriptionByType](#removesubscriptionbytype)
    - [GetSubscriptions](#getsubscriptions)
    - [GetSubscriptionByType](#getsubscriptionbytype)
    - [GetSubscriptionByStatus](#getsubscriptionbystatus)
    - [Handler](#handler)

## Configuration

The Client is configured with a ClientConfig struct. It contains the following fields:

| Field | Description |
| --- | --- |
| `ClientID` | The client ID of your Twitch application. |
| `ClientSecret` | The client secret generated for your Twitch application. **This is not your webhook secret** |
| `Token` | (Optional) If you have generated a token elsewhere in your project you can supply it here. If you have this assigned you can remove `ClientSecret` |
| `WebhookSecret` | The webhook secret used to verify events. This should be a random string between 10-100 characters. |
| `WebhookURL` | The full EventSub URL path, eg: https://mydomain.com/eventsub. |
| `Debug` | If true, the client will log output to stdout. |

## Events

Listening to event is as simple as setting one of the On<EventName> functions in the Client struct

```go
client, err := twitchwh.New(config)
client.OnStreamOnline = func(event twitchwh.StreamOnline) {
	log.Printf("%s just went live!", event.BroadcasterUserLogin)
}
```

Twitch may revoke your subscriptions for a variety of reasons. Whenever TwitchWH recieves a revocation message it fires the OnRevocation handler.

```go
client, err := twitchwh.New(config)
client.OnRevocation = func(sub twitchwh.Subscription) {
	log.Printf("A %s subscription was revoked!", sub.Type)
}
```

## Client methods

### AddSubscription

Adds a new EventSub subscription.

Condition is a struct that contains all the conditions. Use the ones you need for the subscription type.

```go
func (c *Client) AddSubscription(Type string, version string, condition Condition) error
```

```go
err := client.AddSubscription("stream.online", "1", twitchwh.Condition{
	BroadcasterUserID: "215185844",
})
// or
err := client.AddSubscription("channel.raid", "1", twitchwh.Condition{
	ToBroadcasterUserID: "215185844",
})
```

### RemoveSubscription

Removes an existing EventSub subscription by ID.

```go
func (c *Client) RemoveSubscription(id string) error
```

```go
subscriptions, _ := client.GetSubscriptions()
err := client.RemoveSubscription(subscriptions[0].ID) // Assume there is > 0 subscriptions
```

### RemoveSubscriptionByType

Removes an existing EventSub subscription by type and condition.

Similar to `AddSubscription`, but without the version. (And for removal, of course)

```go
func (c *Client) RemoveSubscriptionByType(Type string, condition Condition) error
```

```go
err := client.RemoveSubscriptionByType("stream.online", twitchwh.Condition{
	BroadcasterUserID: "215185844",
})
```

### GetSubscriptions

Get all subscriptions, including revoked ones.

This function will automatically handles pagination.

```go
func (c *Client) GetSubscriptions() (subscriptions []Subscription, err error)
```

```go
subscriptions, err := client.GetSubscriptions()
if err != nil {
	// Handle error
}
```

### GetSubscriptionByType

Identical to `GetSubscriptions`, but with a type. (eg. "stream.online")

```go
func (c *Client) GetSubscriptionsByType(Type string) (subscriptions []Subscription, err error)
```

```go
subscriptions, err := client.GetSubscriptionsByType("stream.online")
if err != nil {
	// Handle error
}
```

### GetSubscriptionByStatus

Identical to `GetSubscriptions`, but with a status. (eg. "enabled")

```go
func (c *Client) GetSubscriptionsByStatus(status string) (subscriptions []Subscription, err error)
```

```go
subscriptions, err := client.GetSubscriptionsByStatus("enabled")
if err != nil {
	// Handle error
}
```

### Handler

This HTTP handler is used to recieve and verify events.
It needs to be assigned to a HTTP server before you can use add subscriptions.

```go
func (c *Client) Handler(w http.ResponseWriter, r *http.Request)
```

```go
http.HandleFunc("/eventsub", client.Handler)
http.ListenAndServe(":8080", nil)
```
