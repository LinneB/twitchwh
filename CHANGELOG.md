## v0.1.0

Released: 2025-01-24

This is a breaking release, mainly focusing on cleanup.

- Events are now generic `json.RawMessage`, and are parsed by the user.
- Helix requests that `401` will be retried after refreshing the token
- Remove `ExternalToken` option

## v0.0.5

Released: 2024-08-19

This is a tiny bugfix that changes the token validation interval from 1 minute to 1 hour.

## v0.0.4

Released: 2024-08-19

This is a **breaking** feature release. It adds support for external tokens using the `SetToken` method and `ExternalToken` config option.

- Added a new `ExternalToken` config option. If true, the client won't validate or generate new tokens.
- Properly validate and regenerate tokens.

## v0.0.3

Released: 2024-07-11

This is a non-breaking feature release that reworks error handling and adds a new handler.

- Added a new `OnRevocation` handler for whenever Twitch revokes a subscription.
- Added custom error types for common conditions. For example: UnauthorizedError for authentication failures.

## v0.0.2

Released: 2024-06-09

This is a **breaking** feature release. It completely changes the behaviour of the GetSubscriptions methods. It also adds a new method, RemoveSubscriptionByType.

- GetSubscriptions, GetSubscriptionsByType, and GetSubscriptionsByStatus now automatically handle pagination.
- Added RemoveSubscriptionByType for removing subscriptions by type and condition.

## v0.0.1

Released: 2024-06-07

Initial release
