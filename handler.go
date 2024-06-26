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

			eventHandlers := map[string]func(json.RawMessage){
				"automod.message.hold":                                   getHandler(c.OnAutomodMessageHold),
				"automod.message.update":                                 getHandler(c.OnAutomodMessageUpdate),
				"automod.settings.update":                                getHandler(c.OnAutomodSettingsUpdate),
				"automod.terms.update":                                   getHandler(c.OnAutomodTermsUpdate),
				"channel.update":                                         getHandler(c.OnChannelUpdate),
				"channel.follow":                                         getHandler(c.OnChannelFollow),
				"channel.ad_break.begin":                                 getHandler(c.OnChannelAdBreakBegin),
				"channel.chat.clear":                                     getHandler(c.OnChannelChatClear),
				"channel.chat.clear_user_messages":                       getHandler(c.OnChannelChatClearUserMessages),
				"channel.chat.message":                                   getHandler(c.OnChannelChatMessage),
				"channel.chat.message_delete":                            getHandler(c.OnChannelChatMessageDelete),
				"channel.chat.notification":                              getHandler(c.OnChannelChatNotification),
				"channel.chat_settings.update":                           getHandler(c.OnChannelChatSettingsUpdate),
				"channel.chat.user_message_hold":                         getHandler(c.OnChannelChatUserMessageHold),
				"channel.chat.user_message_update":                       getHandler(c.OnChannelChatUserMessageUpdate),
				"channel.subscribe":                                      getHandler(c.OnChannelSubscribe),
				"channel.subscription.end":                               getHandler(c.OnChannelSubscriptionEnd),
				"channel.subscription.gift":                              getHandler(c.OnChannelSubscriptionGift),
				"channel.subscription.message":                           getHandler(c.OnChannelSubscriptionMessage),
				"channel.cheer":                                          getHandler(c.OnChannelCheer),
				"channel.raid":                                           getHandler(c.OnChannelRaid),
				"channel.ban":                                            getHandler(c.OnChannelBan),
				"channel.unban":                                          getHandler(c.OnChannelUnban),
				"channel.unban_request.create":                           getHandler(c.OnChannelUnbanRequestCreate),
				"channel.unban_request.resolve":                          getHandler(c.OnChannelUnbanRequestResolve),
				"channel.moderate":                                       getHandler(c.OnChannelModerate),
				"channel.moderator.add":                                  getHandler(c.OnChannelModeratorAdd),
				"channel.moderator.remove":                               getHandler(c.OnChannelModeratorRemove),
				"channel.channel_points_automatic_reward_redemption.add": getHandler(c.OnChannelPointsAutomaticRewardRedemption),
				"channel.channel_points_custom_reward.add":               getHandler(c.OnChannelPointsCustomRewardAdd),
				"channel.channel_points_custom_reward.update":            getHandler(c.OnChannelPointsCustomRewardUpdate),
				"channel.channel_points_custom_reward.remove":            getHandler(c.OnChannelPointsCustomRewardRemove),
				"channel.channel_points_custom_reward_redemption.add":    getHandler(c.OnChannelPointsCustomRewardRedemptionAdd),
				"channel.channel_points_custom_reward_redemption.update": getHandler(c.OnChannelPointsCustomRewardRedemptionUpdate),
				"channel.poll.begin":                                     getHandler(c.OnChannelPollBegin),
				"channel.poll.progress":                                  getHandler(c.OnChannelPollProgress),
				"channel.poll.end":                                       getHandler(c.OnChannelPollEnd),
				"channel.prediction.begin":                               getHandler(c.OnChannelPredictionBegin),
				"channel.prediction.progress":                            getHandler(c.OnChannelPredictionProgress),
				"channel.prediction.lock":                                getHandler(c.OnChannelPredictionLock),
				"channel.prediction.end":                                 getHandler(c.OnChannelPredictionEnd),
				"channel.suspicious_user.message":                        getHandler(c.OnChannelSuspiciousUserMessage),
				"channel.suspicious_user.update":                         getHandler(c.OnChannelSuspiciousUserUpdate),
				"channel.vip.add":                                        getHandler(c.OnChannelVIPAdd),
				"channel.vip.remove":                                     getHandler(c.OnChannelVIPRemove),
				"channel.charity_campaign.donate":                        getHandler(c.OnChannelCharityCampaignDonate),
				"channel.charity_campaign.start":                         getHandler(c.OnChannelCharityCampaignStart),
				"channel.charity_campaign.progress":                      getHandler(c.OnChannelCharityCampaignProgress),
				"channel.charity_campaign.stop":                          getHandler(c.OnChannelCharityCampaignStop),
				"conduit.shard.disabled":                                 getHandler(c.OnConduitShardDisabled),
				"drop.entitlement.grant":                                 getHandler(c.OnDropEntitlementGrant),
				"extension.bits_transaction.create":                      getHandler(c.OnExtensionBitsTransactionCreate),
				"channel.goal.begin":                                     getHandler(c.OnChannelGoalBegin),
				"channel.goal.progress":                                  getHandler(c.OnChannelGoalProgress),
				"channel.goal.end":                                       getHandler(c.OnChannelGoalEnd),
				"channel.hype_train.begin":                               getHandler(c.OnChannelHypeTrainBegin),
				"channel.hype_train.progress":                            getHandler(c.OnChannelHypeTrainProgress),
				"channel.hype_train.end":                                 getHandler(c.OnChannelHypeTrainEnd),
				"channel.shield_mode.begin":                              getHandler(c.OnChannelShieldModeBegin),
				"channel.shield_mode.end":                                getHandler(c.OnChannelShieldModeEnd),
				"channel.shoutout.create":                                getHandler(c.OnChannelShoutoutCreate),
				"channel.shoutout.receive":                               getHandler(c.OnChannelShoutoutReceive),
				"stream.online":                                          getHandler(c.OnStreamOnline),
				"stream.offline":                                         getHandler(c.OnStreamOffline),
				"user.authorization.grant":                               getHandler(c.OnUserAuthorizationGrant),
				"user.authorization.revoke":                              getHandler(c.OnUserAuthorizationRevoke),
				"user.update":                                            getHandler(c.OnUserUpdate),
				"user.whisper.message":                                   getHandler(c.OnUserWhisperMessage),
			}

			if handler, ok := eventHandlers[payload.Subscription.Type]; ok {
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
