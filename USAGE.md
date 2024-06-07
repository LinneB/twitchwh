# Usage

This file contains examples and methods for using TwitchWH.

- [Configuration](#configuration)
- [Client methods](#client-methods)
    - [AddSubscription](#addsubscription)
    - [RemoveSubscription](#removesubscription)
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
subscriptions, _, _ := client.GetSubscriptions("")
err := client.RemoveSubscription(subscriptions[0].ID) // Assume there is > 0 subscriptions
```

### GetSubscriptions

Get all subscriptions, including revoked ones.

Twitch may return a cursor to get the next page of subscriptions. If the returned pagination is not empty, use it as the parameter for the next function call.

```go
func (c *Client) GetSubscriptions(cursor string) (subscriptions []Subscription, pagination string, err error)
```

```go
subscriptions, cursor, _ := client.GetSubscriptions("")
if cursor != "" {
	nextSubscriptions, nextCursor, _ := client.GetSubscriptions(cursor)
}
```

### GetSubscriptionByType

Identical to `GetSubscriptions`, but with a type. (eg. "stream.online")

```go
func (c *Client) GetSubscriptionsByType(Type string, cursor string) (subscriptions []Subscription, pagination string, err error)
```

```go
subscriptions, cursor, _ := client.GetSubscriptionsByType("stream.online", "")
if cursor != "" {
	nextSubscriptions, nextCursor, _ := client.GetSubscriptionsByType("stream.online", cursor)
}
```

### GetSubscriptionByStatus

Identical to `GetSubscriptions`, but with a status. (eg. "enabled")

```go
func (c *Client) GetSubscriptionsByStatus(status string, cursor string) (subscriptions []Subscription, pagination string, err error)
```

```go
subscriptions, cursor, _ := client.GetSubscriptionsByStatus("enabled", "")
if cursor != "" {
	nextSubscriptions, nextCursor, _ := client.GetSubscriptionsByStatus("enabled", cursor)
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