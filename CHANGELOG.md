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
