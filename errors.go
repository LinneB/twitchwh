package twitchwh

import "fmt"

// Helix returned an authorization error. This usually means the token, Client-ID, or client secret are invalid.
type UnauthorizedError struct {
	Body []byte
}

func (e *UnauthorizedError) Error() string {
	return "Helix returned 401 Unauthorized"
}

// Helix returned an unexpected HTTP status code that is not handled by TwitchWH.
type UnhandledStatusError struct {
	Status int
	Body   []byte
}

func (e *UnhandledStatusError) Error() string {
	return fmt.Sprint(e.Status)
}

// Attempted to add a subscription with a type and condition that already exists.
//
// The Condition and Type fields are included in the error.
type DuplicateSubscriptionError struct {
	Condition Condition
	Type      string
}

func (e *DuplicateSubscriptionError) Error() string {
	return "Duplicate subscription"
}

// Could not find a subscription with the specified parameters.
type SubscriptionNotFoundError struct{}

func (e *SubscriptionNotFoundError) Error() string {
	return "Could not find subscription"
}

// Returned whenever AddSubscription times out waiting for verification confirmation.
type VerificationTimeoutError struct {
	Subscription Subscription
}

func (e *VerificationTimeoutError) Error() string {
	return "Subscription was not verified within timeout duration"
}

// Returned for misc errors, like network or serialization errors for example.
type InternalError struct {
	message string
	// Original error that caused this error, if any.
	OriginalError error
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%s: %s", e.message, e.OriginalError)
}
