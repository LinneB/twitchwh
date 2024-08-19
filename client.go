// Package twitchwh is a library for interacting with Twitch EventSub over the Webhook transport.
// It allows you to assign event handlers to specific events.
//
// To get started, create a new client using the New function. Then, assign an event handler using the On<EventType> fields.
// Finally, setup the HTTP handler for your application using the Handler function.
package twitchwh

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// ClientConfig is used to configure a new Client
type ClientConfig struct {
	// Client ID of your Twitch application
	ClientID string
	// Client Secret generated for your Twitch application. !! THIS IS NOT YOUR WEBHOOK SECRET !!
	ClientSecret string
	// If you have generated a token elsewhere in your project you can supply it here
	Token string
	// Webhook secret used to verify events. This should be a random string between 10-100 characters
	WebhookSecret string
	// Full EventSub URL path, eg: https://mydomain.com/eventsub
	WebhookURL string
	// If you have your own token logic, you should set this to true and update tokens using Client.SetToken to prevent duplicates.
	// If this is false twitchwh will use its own internal token generation and validation system.
	ExternalToken bool
	// Log output
	Debug bool
}

type Client struct {
	clientID      string
	clientSecret  string
	token         string
	webhookSecret string
	webhookURL    string
	externalToken bool
	debug         bool

	logger        *log.Logger
	httpClient    *http.Client
	handledEvents []string
	// Client.Handler sends verified IDs to this channel to be read in Client.AddSubscription
	verifiedSubscriptions chan string

	// Fired whenever a subscription is revoked.
	// Check Subscription.Status for the reason.
	OnRevocation func(Subscription)

	OnAutomodMessageHold                        func(AutomodMessageHold)
	OnAutomodMessageUpdate                      func(AutomodMessageUpdate)
	OnAutomodSettingsUpdate                     func(AutomodSettingsUpdate)
	OnAutomodTermsUpdate                        func(AutomodTermsUpdate)
	OnChannelUpdate                             func(ChannelUpdate)
	OnChannelFollow                             func(ChannelFollow)
	OnChannelAdBreakBegin                       func(ChannelAdBreakBegin)
	OnChannelChatClear                          func(ChannelChatClear)
	OnChannelChatClearUserMessages              func(ChannelChatClearUserMessages)
	OnChannelChatMessage                        func(ChannelChatMessage)
	OnChannelChatMessageDelete                  func(ChannelChatMessageDelete)
	OnChannelChatNotification                   func(ChannelChatNotification)
	OnChannelChatSettingsUpdate                 func(ChannelChatSettingsUpdate)
	OnChannelChatUserMessageHold                func(ChannelChatUserMessageHold)
	OnChannelChatUserMessageUpdate              func(ChannelChatUserMessageUpdate)
	OnChannelSubscribe                          func(ChannelSubscribe)
	OnChannelSubscriptionEnd                    func(ChannelSubscriptionEnd)
	OnChannelSubscriptionGift                   func(ChannelSubscriptionGift)
	OnChannelSubscriptionMessage                func(ChannelSubscriptionMessage)
	OnChannelCheer                              func(ChannelCheer)
	OnChannelRaid                               func(ChannelRaid)
	OnChannelBan                                func(ChannelBan)
	OnChannelUnban                              func(ChannelUnban)
	OnChannelUnbanRequestCreate                 func(ChannelUnbanRequestCreate)
	OnChannelUnbanRequestResolve                func(ChannelUnbanRequestResolve)
	OnChannelModerate                           func(ChannelModerate)
	OnChannelModeratorAdd                       func(ChannelModeratorAdd)
	OnChannelModeratorRemove                    func(ChannelModeratorRemove)
	OnChannelPointsAutomaticRewardRedemption    func(ChannelPointsAutomaticRewardRedemption)
	OnChannelPointsCustomRewardAdd              func(ChannelPointsCustomRewardAdd)
	OnChannelPointsCustomRewardUpdate           func(ChannelPointsCustomRewardUpdate)
	OnChannelPointsCustomRewardRemove           func(ChannelPointsCustomRewardRemove)
	OnChannelPointsCustomRewardRedemption       func(ChannelPointsCustomRewardRedemption)
	OnChannelPointsCustomRewardRedemptionAdd    func(ChannelPointsCustomRewardRedemptionAdd)
	OnChannelPointsCustomRewardRedemptionUpdate func(ChannelPointsCustomRewardRedemptionUpdate)
	OnChannelPollBegin                          func(ChannelPollBegin)
	OnChannelPollProgress                       func(ChannelPollProgress)
	OnChannelPollEnd                            func(ChannelPollEnd)
	OnChannelPredictionBegin                    func(ChannelPredictionBegin)
	OnChannelPredictionProgress                 func(ChannelPredictionProgress)
	OnChannelPredictionLock                     func(ChannelPredictionLock)
	OnChannelPredictionEnd                      func(ChannelPredictionEnd)
	OnChannelSuspiciousUserMessage              func(ChannelSuspiciousUserMessage)
	OnChannelSuspiciousUserUpdate               func(ChannelSuspiciousUserUpdate)
	OnChannelVIPAdd                             func(ChannelVIPAdd)
	OnChannelVIPRemove                          func(ChannelVIPRemove)
	OnChannelCharityCampaignDonate              func(ChannelCharityCampaignDonate)
	OnChannelCharityCampaignStart               func(ChannelCharityCampaignStart)
	OnChannelCharityCampaignProgress            func(ChannelCharityCampaignProgress)
	OnChannelCharityCampaignStop                func(ChannelCharityCampaignStop)
	OnConduitShardDisabled                      func(ConduitShardDisabled)
	OnDropEntitlementGrant                      func(DropEntitlementGrant)
	OnExtensionBitsTransactionCreate            func(ExtensionBitsTransactionCreate)
	OnChannelGoalBegin                          func(ChannelGoalBegin)
	OnChannelGoalProgress                       func(ChannelGoalProgress)
	OnChannelGoalEnd                            func(ChannelGoalEnd)
	OnChannelHypeTrainBegin                     func(ChannelHypeTrainBegin)
	OnChannelHypeTrainProgress                  func(ChannelHypeTrainProgress)
	OnChannelHypeTrainEnd                       func(ChannelHypeTrainEnd)
	OnChannelShieldModeBegin                    func(ChannelShieldModeBegin)
	OnChannelShieldModeEnd                      func(ChannelShieldModeEnd)
	OnChannelShoutoutCreate                     func(ChannelShoutoutCreate)
	OnChannelShoutoutReceive                    func(ChannelShoutoutReceive)
	OnStreamOnline                              func(StreamOnline)
	OnStreamOffline                             func(StreamOffline)
	OnUserAuthorizationGrant                    func(UserAuthorizationGrant)
	OnUserAuthorizationRevoke                   func(UserAuthorizationRevoke)
	OnUserUpdate                                func(UserUpdate)
	OnUserWhisperMessage                        func(UserWhisperMessage)
}

// Creates a new client
func New(config ClientConfig) (*Client, error) {
	c := &Client{
		clientID:              config.ClientID,
		clientSecret:          config.ClientSecret,
		token:                 config.Token,
		webhookSecret:         config.WebhookSecret,
		webhookURL:            config.WebhookURL,
		logger:                log.New(os.Stdout, "TwitchWH: ", log.Ltime|log.Lmicroseconds),
		externalToken:         config.ExternalToken,
		debug:                 config.Debug,
		httpClient:            &http.Client{},
		verifiedSubscriptions: make(chan string),
	}

	// Disable logging if debug is false
	if !c.debug {
		c.logger.SetOutput(io.Discard)
	}

	// Generate token if neccesary
	if !c.externalToken {
		c.logger.Println("Using twitchwh internal token store")
		token, err := c.generateToken(c.clientID, c.clientSecret)
		if err != nil {
			return nil, &UnauthorizedError{}
		}
		c.logger.Println("Generated token")
		c.token = token
		go func() {
			for {
				time.Sleep(1 * time.Hour)
				valid, err := c.validateToken(c.token)
				if err != nil {
					c.logger.Printf("Could not validate token: %s", err)
					continue
				}
				if !valid {
					c.logger.Println("Token invalid, generating a new one")
					token, err := c.generateToken(c.clientID, c.clientSecret)
					if err != nil {
						c.logger.Printf("Could not validate token: %s", err)
						continue
					}
					c.token = token
				}
			}
		}()
	} else {
		c.logger.Println("Using external token store")
	}

	return c, nil
}
