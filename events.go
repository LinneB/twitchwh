package twitchwh

import (
	"encoding/json"
	"time"
)

// Generic event handler
func getHandler[T any](handler func(T)) func(json.RawMessage) {
	return func(raw json.RawMessage) {
		var event T
		err := json.Unmarshal(raw, &event)
		if err != nil {
			// We don't have access to the Client logger so we can't log this error
			// TODO: Fix this
			return
		}
		if handler != nil {
			handler(event)
		}
	}
}

// automod.message.hold
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#automodmessagehold
type AutomodMessageHold struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	UserID               string    `json:"user_id"`
	UserName             string    `json:"user_name"`
	UserLogin            string    `json:"user_login"`
	MessageID            string    `json:"message_id"`
	Message              string    `json:"message"`
	Level                int       `json:"level"`
	Category             string    `json:"category"`
	HeldAt               time.Time `json:"held_at"`
	Fragments            struct {
		Emotes []struct {
			Text  string `json:"text"`
			ID    string `json:"id"`
			SetID string `json:"set-id"`
		} `json:"emotes"`
		Cheermotes []struct {
			Text   string `json:"text"`
			Amount int    `json:"amount"`
			Prefix string `json:"prefix"`
			Tier   int    `json:"tier"`
		} `json:"cheermotes"`
	} `json:"fragments"`
}

// automod.message.update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#automodmessageupdate
type AutomodMessageUpdate struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	UserID               string    `json:"user_id"`
	UserName             string    `json:"user_name"`
	UserLogin            string    `json:"user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	MessageID            string    `json:"message_id"`
	Message              string    `json:"message"`
	Level                int       `json:"level"`
	Category             string    `json:"category"`
	Status               string    `json:"status"`
	HeldAt               time.Time `json:"held_at"`
	Fragments            struct {
		Emotes []struct {
			Text  string `json:"text"`
			ID    string `json:"id"`
			SetID string `json:"set-id"`
		} `json:"emotes"`
		Cheermotes []struct {
			Text   string `json:"text"`
			Amount int    `json:"amount"`
			Prefix string `json:"prefix"`
			Tier   int    `json:"tier"`
		} `json:"cheermotes"`
	} `json:"fragments"`
}

// automod.settings.update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#automodsettingsupdate
type AutomodSettingsUpdate struct {
	Data []struct {
		BroadcasterUserID       string `json:"broadcaster_user_id"`
		BroadcasterUserName     string `json:"broadcaster_user_name"`
		BroadcasterUserLogin    string `json:"broadcaster_user_login"`
		ModeratorUserID         string `json:"moderator_user_id"`
		ModeratorUserName       string `json:"moderator_user_name"`
		ModeratorUserLogin      string `json:"moderator_user_login"`
		OverallLevel            int    `json:"overall_level"`
		Disability              int    `json:"disability"`
		Aggression              int    `json:"aggression"`
		SexualitySexOrGender    int    `json:"sexuality_sex_or_gender"`
		Misogyny                int    `json:"misogyny"`
		Bullying                int    `json:"bullying"`
		Swearing                int    `json:"swearing"`
		RaceEthnicityOrReligion int    `json:"race_ethnicity_or_religion"`
		SexBasedTerms           int    `json:"sex_based_terms"`
	} `json:"data"`
}

// automod.terms.update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#automodtermsupdate
type AutomodTermsUpdate struct {
	BroadcasterUserID    string   `json:"broadcaster_user_id"`
	BroadcasterUserName  string   `json:"broadcaster_user_name"`
	BroadcasterUserLogin string   `json:"broadcaster_user_login"`
	ModeratorUserID      string   `json:"moderator_user_id"`
	ModeratorUserLogin   string   `json:"moderator_user_login"`
	ModeratorUserName    string   `json:"moderator_user_name"`
	Action               string   `json:"action"`
	FromAutomod          bool     `json:"from_automod"`
	Terms                []string `json:"terms"`
}

// channel.update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelupdate
type ChannelUpdate struct {
	BroadcasterUserID           string   `json:"broadcaster_user_id"`
	BroadcasterUserLogin        string   `json:"broadcaster_user_login"`
	BroadcasterUserName         string   `json:"broadcaster_user_name"`
	Title                       string   `json:"title"`
	Language                    string   `json:"language"`
	CategoryID                  string   `json:"category_id"`
	CategoryName                string   `json:"category_name"`
	ContentClassificationLabels []string `json:"content_classification_labels"`
}

// channel.follow
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelfollow
type ChannelFollow struct {
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	FollowedAt           time.Time `json:"followed_at"`
}

// channel.ad_break.begin
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelad_breakbegin
type ChannelAdBreakBegin struct {
	DurationSeconds      string    `json:"duration_seconds"`
	StartedAt            time.Time `json:"started_at"`
	IsAutomatic          string    `json:"is_automatic"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	RequesterUserID      string    `json:"requester_user_id"`
	RequesterUserLogin   string    `json:"requester_user_login"`
	RequesterUserName    string    `json:"requester_user_name"`
}

// channel.chat.clear
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchatclear
type ChannelChatClear struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
}

// channel.chat.clear_user_messages
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchatclear_user_messages
type ChannelChatClearUserMessages struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	TargetUserID         string `json:"target_user_id"`
	TargetUserName       string `json:"target_user_name"`
	TargetUserLogin      string `json:"target_user_login"`
}

// channel.chat.message
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchatmessage
type ChannelChatMessage struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	ChatterUserID        string `json:"chatter_user_id"`
	ChatterUserName      string `json:"chatter_user_name"`
	ChatterUserLogin     string `json:"chatter_user_login"`
	MessageID            string `json:"message_id"`
	Message              struct {
		Text      string `json:"text"`
		Fragments []struct {
			Type      string `json:"type"`
			Text      string `json:"text"`
			Cheermote struct {
				Prefix string `json:"prefix"`
				Bits   int    `json:"bits"`
				Tier   int    `json:"tier"`
			} `json:"cheermote"`
			Emote struct {
				ID         string   `json:"id"`
				EmoteSetID string   `json:"emote_set_id"`
				OwnerID    string   `json:"owner_id"`
				Format     []string `json:"format"`
			} `json:"emote"`
			Mention struct {
				UserID    string `json:"user_id"`
				UserName  string `json:"user_name"`
				UserLogin string `json:"user_login"`
			} `json:"mention"`
		} `json:"fragments"`
	} `json:"message"`
	MessageType string `json:"message_type"`
	Badges      []struct {
		SetID string `json:"set_id"`
		ID    string `json:"id"`
		Info  string `json:"info"`
	} `json:"badges"`
	Cheer struct {
		Bits int `json:"bits"`
	} `json:"cheer"`
	Color string `json:"color"`
	Reply struct {
		ParentMessageID   string `json:"parent_message_id"`
		ParentMessageBody string `json:"parent_message_body"`
		ParentUserID      string `json:"parent_user_id"`
		ParentUserName    string `json:"parent_user_name"`
		ParentUserLogin   string `json:"parent_user_login"`
		ThreadMessageID   string `json:"thread_message_id"`
		ThreadUserID      string `json:"thread_user_id"`
		ThreadUserName    string `json:"thread_user_name"`
		ThreadUserLogin   string `json:"thread_user_login"`
	} `json:"reply"`
	ChannelPointsCustomRewardID string `json:"channel_points_custom_reward_id"`
	ChannelPointsAnimationID    string `json:"channel_points_animation_id"`
}

// channel.chat.message_delete
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchatmessage_delete
type ChannelChatMessageDelete struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	TargetUserID         string `json:"target_user_id"`
	TargetUserName       string `json:"target_user_name"`
	TargetUserLogin      string `json:"target_user_login"`
	MessageID            string `json:"message_id"`
}

// channel.chat.notification
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchatnotification
type ChannelChatNotification struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	ChatterUserID        string `json:"chatter_user_id"`
	ChatterUserName      string `json:"chatter_user_name"`
	ChatterUserLogin     string `json:"chatter_user_login"`
	ChatterIsAnonymous   bool   `json:"chatter_is_anonymous"`
	Color                string `json:"color"`
	Badges               []struct {
		SetID string `json:"set_id"`
		ID    string `json:"id"`
		Info  string `json:"info"`
	} `json:"badges"`
	SystemMessage string `json:"system_message"`
	MessageID     string `json:"message_id"`
	Message       struct {
		Text      string `json:"text"`
		Fragments []struct {
			Type      string `json:"type"`
			Text      string `json:"text"`
			Cheermote struct {
				Prefix string `json:"prefix"`
				Bits   int    `json:"bits"`
				Tier   int    `json:"tier"`
			} `json:"cheermote"`
			Emote struct {
				ID         string   `json:"id"`
				EmoteSetID string   `json:"emote_set_id"`
				OwnerID    string   `json:"owner_id"`
				Format     []string `json:"format"`
			} `json:"emote"`
			Mention struct {
				UserID    string `json:"user_id"`
				Username  string `json:"user_name"`
				UserLogin string `json:"user_login"`
			}
		} `json:"fragments"`
	} `json:"message"`
	NoticeType string `json:"notice_type"`
	Sub        struct {
		SubTier        string `json:"sub_tier"`
		IsPrime        bool   `json:"is_prime"`
		DurationMonths int    `json:"duration_months"`
	}
	Resub struct {
		CumulativeMonths  int    `json:"cumulative_months"`
		DurationMonths    int    `json:"duration_months"`
		StreakMonths      int    `json:"streak_months"`
		SubTier           string `json:"sub_tier"`
		IsPrime           bool   `json:"is_prime"`
		IsGift            bool   `json:"is_gift"`
		GifterIsAnonymous bool   `json:"gifter_is_anonymous"`
		GifterUserID      string `json:"gifter_user_id"`
		GifterUserName    string `json:"gifter_user_name"`
		GifterUserLogin   string `json:"gifter_user_login"`
	} `json:"resub"`
	SubGift struct {
		DurationMonths     int    `json:"duration_months"`
		CumulativeTotal    int    `json:"cumulative_total"`
		RecipientUserID    string `json:"recipient_user_id"`
		RecipientUserName  string `json:"recipient_user_name"`
		RecipientUserLogin string `json:"recipient_user_login"`
		SubTier            string `json:"sub_tier"`
		CommunityGiftID    string `json:"community_gift_id"`
	} `json:"sub_gift"`
	CommunitySubGift struct {
		ID              string `json:"id"`
		Total           int    `json:"total"`
		SubTier         string `json:"sub_tier"`
		CumulativeTotal int    `json:"cumulative_total"`
	} `json:"community_sub_gift"`
	GiftPaidUpgrade struct {
		GifterIsAnonymous bool   `json:"gifter_is_anonymous"`
		GifterUserID      string `json:"gifter_user_id"`
		GifterUserName    string `json:"gifter_user_name"`
		GifterUserLogin   string `json:"gifter_user_login"`
	} `json:"gift_paid_upgrade"`
	PrimePaidUpgrade struct {
		SubTier string `json:"sub_tier"`
	} `json:"prime_paid_upgrade"`
	Raid struct {
		UserID          string `json:"user_id"`
		UserName        string `json:"user_name"`
		UserLogin       string `json:"user_login"`
		ViewerCount     int    `json:"viewer_count"`
		ProfileImageURL string `json:"profile_image_url"`
	} `json:"raid"`
	Unraid       struct{} `json:"unraid"`
	PayItForward struct {
		GifterIsAnonymous bool   `json:"gifter_is_anonymous"`
		GifterUserID      string `json:"gifter_user_id"`
		GifterUserName    string `json:"gifter_user_name"`
		GifterUserLogin   string `json:"gifter_user_login"`
	} `json:"pay_it_forward"`
	Announcement struct {
		Color string `json:"color"`
	} `json:"announcement"`
	CharityDonation struct {
		CharityName string `json:"charity_name"`
		Amount      struct {
			Value        int    `json:"value"`
			DecimalPlace int    `json:"decimal_place"`
			Currency     string `json:"currency"`
		} `json:"amount"`
	} `json:"charity_donation"`
	BitsBadgeTier struct {
		Tier int `json:"tier"`
	} `json:"bits_badge_tier"`
}

// channel.chat_settings.update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchat_settingsupdate
type ChannelChatSettingsUpdate struct {
	BroadcasterUserID           string `json:"broadcaster_user_id"`
	BroadcasterUserLogin        string `json:"broadcaster_user_login"`
	BroadcasterUserName         string `json:"broadcaster_user_name"`
	EmoteMode                   bool   `json:"emote_mod"`
	FollowerMode                bool   `json:"follower_mode"`
	FollowerModeDurationMinutes int    `json:"follower_mode_duration_minutes"`
	SlowMode                    bool   `json:"slow_mode"`
	SlowModeWaitTimeSeconds     int    `json:"slow_mode_wait_time_seconds"`
	SubscriberMode              bool   `json:"subscriber_mode"`
	UniqueChatMode              bool   `json:"unique_chat_mode"`
}

// channel.chat.user_message_hold
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchatuser_message_hold
type ChannelChatUserMessageHold struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	MessageID            string `json:"message_id"`
	Message              struct {
		Text      string `json:"text"`
		Fragments []struct {
			Type  string `json:"type"`
			Text  string `json:"text"`
			Emote struct {
				ID         string `json:"id"`
				EmoteSetID string `json:"emote_set_id"`
			} `json:"emote"`
			Cheermote struct {
				Prefix string `json:"prefix"`
				Bits   int    `json:"bits"`
				Tier   int    `json:"tier"`
			} `json:"cheermote"`
		}
	} `json:"message"`
}

// channel.chat.user_message_update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchatuser_message_update
type ChannelChatUserMessageUpdate struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	Status               string `json:"status"`
	MessageID            string `json:"message_id"`
	Message              struct {
		Text      string `json:"text"`
		Fragments []struct {
			Type  string `json:"type"`
			Text  string `json:"text"`
			Emote struct {
				ID         string `json:"id"`
				EmoteSetID string `json:"emote_set_id"`
			} `json:"emote"`
			Cheermote struct {
				Prefix string `json:"prefix"`
				Bits   int    `json:"bits"`
				Tier   int    `json:"tier"`
			} `json:"cheermote"`
		}
	} `json:"message"`
}

// channel.subscribe
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelsubscribe
type ChannelSubscribe struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

// channel.subscription.end
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelsubscriptionend
type ChannelSubscriptionEnd struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

// channel.subscription.gift
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelsubscriptiongift
type ChannelSubscriptionGift struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Total                int    `json:"total"`
	Tier                 string `json:"tier"`
	CumulativeTotal      int    `json:"cumulative_total"`
	IsAnonymous          bool   `json:"is_anonymous"`
}

// channel.subscription.message
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelsubscriptionmessage
type ChannelSubscriptionMessage struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	Message              struct {
		Text   string `json:"text"`
		Emotes []struct {
			Begin int    `json:"begin"`
			End   int    `json:"end"`
			ID    string `json:"id"`
		}
	} `json:"message"`
	CumulativeMonths int `json:"cumulative_months"`
	StreakMonths     int `json:"streak_months"`
	DurationMonths   int `json:"duration_months"`
}

// channel.cheer
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelcheer
type ChannelCheer struct {
	IsAnonymous          bool   `json:"is_anonymous"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Message              string `json:"message"`
	Bits                 int    `json:"bits"`
}

// channel.raid
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelraid
type ChannelRaid struct {
	FromBroadcasterUserID    string `json:"from_broadcaster_user_id"`
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
	ToBroadcasterUserID      string `json:"to_broadcaster_user_id"`
	ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`
	ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`
	Viewers                  int    `json:"viewers"`
}

// channel.ban
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelban
type ChannelBan struct {
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	Reason               string    `json:"reason"`
	BannedAt             time.Time `json:"banned_at"`
	EndsAt               time.Time `json:"ends_at"`
	IsPermanent          bool      `json:"is_permanent"`
}

// channel.unban
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelunban
type ChannelUnban struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ModeratorUserID      string `json:"moderator_user_id"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	ModeratorUserName    string `json:"moderator_user_name"`
}

// channel.unban_request.create
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelunban_requestcreate
type ChannelUnbanRequestCreate struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	Text                 string    `json:"text"`
	CreatedAt            time.Time `json:"created_at"`
}

// channel.unban_request.resolve
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelunban_requestresolve
type ChannelUnbanRequestResolve struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ModeratorUserID      string `json:"moderator_user_id"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	ModeratorUserName    string `json:"moderator_user_name"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	ResolutionText       string `json:"resolution_text"`
	Status               string `json:"status"`
}

// channel.moderate
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelmoderate
type ChannelModerate struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ModeratorUserID      string `json:"moderator_user_id"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	ModeratorUserName    string `json:"moderator_user_name"`
	Action               string `json:"action"`
	Followers            struct {
		FollowDurationMinutes int `json:"follow_duration_minutes"`
	} `json:"followers"`
	Slow struct {
		WaitTimeSeconds int `json:"wait_time_seconds"`
	} `json:"slow"`
	Vip struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"vip"`
	UnVip struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"unvip"`
	Mod struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"mod"`
	UnMod struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"unmod"`
	Ban struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
		Reason    string `json:"reason"`
	} `json:"ban"`
	UnBan struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"unban"`
	Timeout struct {
		UserID    string    `json:"user_id"`
		UserLogin string    `json:"user_login"`
		UserName  string    `json:"user_name"`
		Reason    string    `json:"reason"`
		ExpiresAt time.Time `json:"expires_at"`
	} `json:"timeout"`
	UnTimeout struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"untimeout"`
	Raid struct {
		UserID      string `json:"user_id"`
		UserLogin   string `json:"user_login"`
		UserName    string `json:"user_name"`
		ViewerCount int    `json:"viewer_count"`
	} `json:"raid"`
	UnRaid struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"unraid"`
	Delete struct {
		UserID      string `json:"user_id"`
		UserLogin   string `json:"user_login"`
		UserName    string `json:"user_name"`
		MessageID   string `json:"message_id"`
		MessageBody string `json:"message_body"`
	} `json:"delete"`
	AutomodTerms struct {
		Action      string   `json:"action"`
		List        string   `json:"list"`
		Terms       []string `json:"terms"`
		FromAutomod bool     `json:"from_automod"`
	} `json:"automod_terms"`
	UnbanRequest struct {
		IsApproved       bool   `json:"is_approved"`
		UserID           string `json:"user_id"`
		UserLogin        string `json:"user_login"`
		UserName         string `json:"user_name"`
		ModeratorMessage string `json:"moderator_message"`
	}
}

// channel.moderator.add
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelmoderatoradd
type ChannelModeratorAdd struct {
	UserID               string `json:"user_id"`
	UserName             string `json:"user_name"`
	UserLogin            string `json:"user_login"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
}

// channel.moderator.remove
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelmoderatorremove
type ChannelModeratorRemove struct {
	UserID               string `json:"user_id"`
	UserName             string `json:"user_name"`
	UserLogin            string `json:"user_login"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
}

// channel.channel_points_automatic_reward_redemption.add
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchannel_points_automatic_reward_redemptionadd
type ChannelPointsAutomaticRewardRedemption struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	UserID               string `json:"user_id"`
	UserName             string `json:"user_name"`
	UserLogin            string `json:"user_login"`
	ID                   string `json:"id"`
	Reward               struct {
		Type          string `json:"type"`
		Cost          int    `json:"cost"`
		UnlockedEmote struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"unlocked_emote"`
	} `json:"reward"`
	Message struct {
		Text   string `json:"text"`
		Emotes []struct {
			ID    string `json:"id"`
			Begin int    `json:"begin"`
			End   int    `json:"end"`
		}
	} `json:"message"`
	UserInput  string    `json:"user_input"`
	RedeemedAt time.Time `json:"redeemed_at"`
}

// Generic Channel Points event type used by [ChannelPointsCustomRewardAdd] [ChannelPointsCustomRewardUpdate] [ChannelPointsCustomRewardRemove]
type ChannelPointsCustomReward struct {
	ID                                string `json:"id"`
	BroadcasterUserID                 string `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string `json:"broadcaster_user_login"`
	BroadcasterUserName               string `json:"broadcaster_user_name"`
	IsEnabled                         bool   `json:"is_enabled"`
	IsPaused                          bool   `json:"is_paused"`
	IsInStock                         bool   `json:"is_in_stock"`
	Title                             string `json:"title"`
	Cost                              int    `json:"cost"`
	Prompt                            string `json:"prompt"`
	IsUserInputRequired               bool   `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool   `json:"should_redemptions_skip_request_queue"`
	MaxPerStream                      struct {
		IsEnabled bool `json:"is_enabled"`
		Value     int  `json:"value"`
	} `json:"max_per_stream"`
	MaxPerUserPerStream struct {
		IsEnabled bool `json:"is_enabled"`
		Value     int  `json:"value"`
	} `json:"max_per_user_per_stream"`
	BackgroundColor string `json:"background_color"`
	Image           struct {
		Url1x string `json:"url_1x"`
		Url2x string `json:"url_2x"`
		Url4x string `json:"url_4x"`
	} `json:"image"`
	DefaultImage struct {
		Url1x string `json:"url_1x"`
		Url2x string `json:"url_2x"`
		Url4x string `json:"url_4x"`
	} `json:"default_image"`
	GlobalCooldown struct {
		IsEnabled bool `json:"is_enabled"`
		Seconds   int  `json:"seconds"`
	} `json:"global_cooldown"`
	CooldownExpiresAt                time.Time `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream int       `json:"redemptions_redeemed_current_stream"`
}

// channel.channel_points_custom_reward.add
// Alias of [ChannelPointsCustomReward]
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchannel_points_custom_rewardadd
type ChannelPointsCustomRewardAdd ChannelPointsCustomReward

// channel.channel_points_custom_reward.update
// Alias of [ChannelPointsCustomReward]
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchannel_points_custom_rewardupdate
type ChannelPointsCustomRewardUpdate ChannelPointsCustomReward

// channel.channel_points_custom_reward.remove
// Alias of [ChannelPointsCustomReward]
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchannel_points_custom_rewardremove
type ChannelPointsCustomRewardRemove ChannelPointsCustomReward

// Generic Channel Points Redemption event used by [ChannelPointsCustomRewardRedemptionAdd] [ChannelPointsCustomRewardRedemptionUpdate]
type ChannelPointsCustomRewardRedemption struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	UserInput            string `json:"user_input"`
	Status               string `json:"status"`
	Reward               struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Cost   int    `json:"cost"`
		Prompt string `json:"prompt"`
	} `json:"reward"`
	RedeemedAt time.Time `json:"redeemed_at"`
}

// channel.channel_points_custom_reward_redemption.add
// Alias of [ChannelPointsCustomRewardRedemption]
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchannel_points_custom_reward_redemptionadd
type ChannelPointsCustomRewardRedemptionAdd ChannelPointsCustomRewardRedemption

// channel.channel_points_custom_reward_redemption.update
// Alias of [ChannelPointsCustomRewardRedemption]
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelchannel_points_custom_reward_redemptionupdate
type ChannelPointsCustomRewardRedemptionUpdate ChannelPointsCustomRewardRedemption

// Generic Channel Poll event type used by [ChannelPollBegin] and [ChannelPollProgress]
type ChannelPoll struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Choices              []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		// According to Twitch this is unused and will always be 0
		BitsVotes          int `json:"bits_votes"`
		ChannelPointsVotes int `json:"channel_points_votes"`
		Votes              int `json:"votes"`
	} `json:"choices"`
	// According to Twitch this is unused and should be ignored
	BitsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"bits_voting"`
	ChannelPointsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"channel_points_voting"`
	StartedAt time.Time `json:"started_at"`
	EndsAt    time.Time `json:"ends_at"`
}

// channel.poll.begin
// Alias of [ChannelPoll]
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelpollbegin
type ChannelPollBegin ChannelPoll

// channel.poll.progress
// Alias of [ChannelPoll]
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelpollprogress
type ChannelPollProgress ChannelPoll

// channel.poll.end
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelpollend
type ChannelPollEnd struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Choices              []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		// According to Twitch this is unused and will always be 0
		BitsVotes          int `json:"bits_votes"`
		ChannelPointsVotes int `json:"channel_points_votes"`
		Votes              int `json:"votes"`
	} `json:"choices"`
	// According to Twitch this is unused and should be ignored
	BitsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"bits_voting"`
	ChannelPointsVoting struct {
		IsEnabled     bool `json:"is_enabled"`
		AmountPerVote int  `json:"amount_per_vote"`
	} `json:"channel_points_voting"`
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
	EndsAt    time.Time `json:"ends_at"`
}

// channel.prediction.begin
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelpredictionbegin
type ChannelPredictionBegin struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Outcomes             []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Color string `json:"color"`
	}
	StartedAt time.Time `json:"started_at"`
	LocksAt   time.Time `json:"locks_at"`
}

// channel.prediction.progress
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelpredictionprogress
type ChannelPredictionProgress struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Outcomes             []struct {
		ID            string `json:"id"`
		Title         string `json:"title"`
		Color         string `json:"color"`
		Users         int    `json:"users"`
		ChannelPoints int    `json:"channel_points"`
		TopPredictors []struct {
			UserID            string `json:"user_id"`
			UserLogin         string `json:"user_login"`
			UserName          string `json:"user_name"`
			ChannelPointsUsed int    `json:"channel_points_used"`
		}
	}
	StartedAt time.Time `json:"started_at"`
	LocksAt   time.Time `json:"locks_at"`
}

// channel.prediction.lock
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelpredictionlock
type ChannelPredictionLock struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Outcomes             []struct {
		ID            string `json:"id"`
		Title         string `json:"title"`
		Color         string `json:"color"`
		Users         int    `json:"users"`
		ChannelPoints int    `json:"channel_points"`
		TopPredictors []struct {
			UserID            string `json:"user_id"`
			UserLogin         string `json:"user_login"`
			UserName          string `json:"user_name"`
			ChannelPointsUsed int    `json:"channel_points_used"`
		}
	}
	StartedAt time.Time `json:"started_at"`
	LocksAt   time.Time `json:"locks_at"`
}

// channel.prediction.end
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelpredictionend
type ChannelPredictionEnd struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	WinningOutcomeID     string `json:"winning_outcome_id"`
	Outcomes             []struct {
		ID            string `json:"id"`
		Title         string `json:"title"`
		Color         string `json:"color"`
		Users         int    `json:"users"`
		ChannelPoints int    `json:"channel_points"`
		TopPredictors []struct {
			UserID            string `json:"user_id"`
			UserLogin         string `json:"user_login"`
			UserName          string `json:"user_name"`
			ChannelPointsWon  int    `json:"channel_points_won"`
			ChannelPointsUsed int    `json:"channel_points_used"`
		}
	}
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
	LocksAt   time.Time `json:"locks_at"`
}

// channel.suspicious_user.message
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelsuspicious_usermessage
type ChannelSuspiciousUserMessage struct {
	BroadcasterUserID     string   `json:"broadcaster_user_id"`
	BroadcasterUserName   string   `json:"broadcaster_user_name"`
	BroadcasterUserLogin  string   `json:"broadcaster_user_login"`
	UserID                string   `json:"user_id"`
	UserName              string   `json:"user_name"`
	UserLogin             string   `json:"user_login"`
	LowTrustStatus        string   `json:"low_trust_status"`
	SharedBanChannelIDs   []string `json:"shared_ban_channel_ids"`
	Types                 []string `json:"types"`
	BanEvasionEvauluation string   `json:"ban_evasion_evaluation"`
	Message               struct {
		MessageID string `json:"message_id"`
		Text      string `json:"text"`
		Fragments []struct {
			Type      string `json:"type"`
			Text      string `json:"text"`
			Cheermote struct {
				Prefix string `json:"prefix"`
				Bits   int    `json:"bits"`
				Tier   int    `json:"tier"`
			} `json:"cheermote"`
			Emote struct {
				ID         string `json:"id"`
				EmoteSetID string `json:"emote_set_id"`
			} `json:"emote"`
		} `json:"fragments"`
	} `json:"message"`
}

// channel.suspicious_user.update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelsuspicious_userupdate
type ChannelSuspiciousUserUpdate struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	ModeratorUserID      string `json:"moderator_user_id"`
	ModeratorUserName    string `json:"moderator_user_name"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	UserID               string `json:"user_id"`
	UserName             string `json:"user_name"`
	UserLogin            string `json:"user_login"`
	LowTrustStatus       string `json:"low_trust_status"`
}

// channel.vip.add
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelvipadd
type ChannelVIPAdd struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

// channel.vip.remove
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelvipremove
type ChannelVIPRemove struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

// channel.charity_campaign.donate
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelcharity_campaigndonate
type ChannelCharityCampaignDonate struct {
	ID                   string `json:"id"`
	CampaignID           string `json:"campaign_id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	CharityName          string `json:"charity_name"`
	CharityDescription   string `json:"charity_description"`
	CharityLogo          string `json:"charity_logo"`
	CharityWebsite       string `json:"charity_website"`
	Amount               struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	} `json:"amount"`
}

// channel.charity_campaign.start
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelcharity_campaignstart
type ChannelCharityCampaignStart struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	CharityName          string `json:"charity_name"`
	CharityDescription   string `json:"charity_description"`
	CharityLogo          string `json:"charity_logo"`
	CharityWebsite       string `json:"charity_website"`
	CurrentAmount        struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	} `json:"current_amount"`
	TargetAmount struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	} `json:"target_amount"`
	StartedAt time.Time `json:"started_at"`
}

// channel.charity_campaign.progress
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelcharity_campaignprogress
type ChannelCharityCampaignProgress struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	CharityName          string `json:"charity_name"`
	CharityDescription   string `json:"charity_description"`
	CharityLogo          string `json:"charity_logo"`
	CharityWebsite       string `json:"charity_website"`
	CurrentAmount        struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	} `json:"current_amount"`
	TargetAmount struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	} `json:"target_amount"`
}

// channel.charity_campaign.stop
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelcharity_campaignstop
type ChannelCharityCampaignStop struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	CharityName          string `json:"charity_name"`
	CharityDescription   string `json:"charity_description"`
	CharityLogo          string `json:"charity_logo"`
	CharityWebsite       string `json:"charity_website"`
	CurrentAmount        struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	} `json:"current_amount"`
	TargetAmount struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	} `json:"target_amount"`
	StoppedAt time.Time `json:"stopped_at"`
}

// conduit.shard.disabled
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#conduitsharddisabled
type ConduitShardDisabled struct {
	ConduitID string `json:"conduit_id"`
	ShardID   string `json:"shard_id"`
	Status    string `json:"status"`
	Transport struct {
		Method         string    `json:"method"`
		Callback       string    `json:"callback"`
		SessionID      string    `json:"session_id"`
		ConnectedAt    time.Time `json:"connected_at"`
		DisconnectedAt time.Time `json:"disconnected_at"`
	}
}

// drop.entitlement.grant
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#dropentitlementgrant
type DropEntitlementGrant []struct {
	ID   string `json:"id"`
	Data struct {
		OrganizationID string    `json:"organization_id"`
		CategoryID     string    `json:"category_id"`
		CategoryName   string    `json:"category_name"`
		UserID         string    `json:"user_id"`
		UserName       string    `json:"user_name"`
		UserLogin      string    `json:"user_login"`
		EntitlementID  string    `json:"entitlement_id"`
		BenefitID      string    `json:"benefit_id"`
		CreatedAt      time.Time `json:"created_at"`
	} `json:"data"`
}

// extension.bits_transaction.create
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#extensionbits_transactioncreate
type ExtensionBitsTransactionCreate struct {
	ID                   string `json:"id"`
	ExtensionClientID    string `json:"extension_client_id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	UserName             string `json:"user_name"`
	UserLogin            string `json:"user_login"`
	UserID               string `json:"user_id"`
	Product              struct {
		Name          string `json:"name"`
		SKU           string `json:"sku"`
		Bits          int    `json:"bits"`
		InDevelopment bool   `json:"in_development"`
	} `json:"product"`
}

// channel.goal.begin
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelgoalbegin
type ChannelGoalBegin struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	Type                 string    `json:"type"`
	Decscription         string    `json:"description"`
	CurrentAmount        int       `json:"current_amount"`
	TargetAmount         int       `json:"target_amount"`
	StartedAt            time.Time `json:"started_at"`
}

// channel.goal.progress
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelgoalprogress
type ChannelGoalProgress struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	Type                 string    `json:"type"`
	Decscription         string    `json:"description"`
	CurrentAmount        int       `json:"current_amount"`
	TargetAmount         int       `json:"target_amount"`
	StartedAt            time.Time `json:"started_at"`
}

// channel.goal.end
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelgoalend
type ChannelGoalEnd struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	Type                 string    `json:"type"`
	IsAchieved           bool      `json:"is_achieved"`
	Decscription         string    `json:"description"`
	CurrentAmount        int       `json:"current_amount"`
	TargetAmount         int       `json:"target_amount"`
	StartedAt            time.Time `json:"started_at"`
	EndedAt              time.Time `json:"ended_at"`
}

// channel.hype_train.begin
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelhype_trainbegin
type ChannelHypeTrainBegin struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	Total                int    `json:"total"`
	Progress             int    `json:"progress"`
	Goal                 int    `json:"goal"`
	TopContributions     []struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
		Type      string `json:"type"`
		Total     int    `json:"total"`
	} `json:"top_contributions"`
	LastContribution struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
		Type      string `json:"type"`
		Total     int    `json:"total"`
	} `json:"last_contribution"`
	Level     int       `json:"level"`
	StartedAt time.Time `json:"started_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// channel.hype_train.progress
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelhype_trainprogress
type ChannelHypeTrainProgress struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	Level                int    `json:"level"`
	Total                int    `json:"total"`
	Progress             int    `json:"progress"`
	Goal                 int    `json:"goal"`
	TopContributions     []struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
		Type      string `json:"type"`
		Total     int    `json:"total"`
	} `json:"top_contributions"`
	LastContribution struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
		Type      string `json:"type"`
		Total     int    `json:"total"`
	} `json:"last_contribution"`
	StartedAt time.Time `json:"started_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// channel.hype_train.end
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelhype_trainend
type ChannelHypeTrainEnd struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	Level                int    `json:"level"`
	Total                int    `json:"total"`
	TopContributions     []struct {
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
		Type      string `json:"type"`
		Total     int    `json:"total"`
	} `json:"top_contributions"`
	StartedAt       time.Time `json:"started_at"`
	EndedAt         time.Time `json:"ended_at"`
	CooldownEndstAt time.Time `json:"cooldown_ends_at"`
}

// channel.shield_mode.begin
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelshield_modebegin
type ChannelShieldModeBegin struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	StartedAt            time.Time `json:"started_at"`
}

// channel.shield_mode.end
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelshield_modeend
type ChannelShieldModeEnd struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	EndedAt              time.Time `json:"ended_at"`
}

// channel.shoutout.create
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelshoutoutcreate
type ChannelShoutoutCreate struct {
	BroadcasterUserID      string    `json:"broadcaster_user_id"`
	BroadcasterUserName    string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin   string    `json:"broadcaster_user_login"`
	ModeratorUserID        string    `json:"moderator_user_id"`
	ModeratorUserName      string    `json:"moderator_user_name"`
	ModeratorUserLogin     string    `json:"moderator_user_login"`
	ToBroadcasterUserID    string    `json:"to_broadcaster_user_id"`
	ToBroadcasterUserName  string    `json:"to_broadcaster_user_name"`
	ToBroadcasterUserLogin string    `json:"to_broadcaster_user_login"`
	StartedAt              time.Time `json:"started_at"`
	ViewerCount            int       `json:"viewer_count"`
	CooldownEndstAt        time.Time `json:"cooldown_ends_at"`
	TargetCooldownEndstAt  time.Time `json:"target_cooldown_ends_at"`
}

// channel.shoutout.receive
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#channelshoutoutreceive
type ChannelShoutoutReceive struct {
	BroadcasterUserID        string    `json:"broadcaster_user_id"`
	BroadcasterUserName      string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin     string    `json:"broadcaster_user_login"`
	FromBroadcasterUserID    string    `json:"from_broadcaster_user_id"`
	FromBroadcasterUserName  string    `json:"from_broadcaster_user_name"`
	FromBroadcasterUserLogin string    `json:"from_broadcaster_user_login"`
	ViewerCount              int       `json:"viewer_count"`
	StartedAt                time.Time `json:"started_at"`
}

// stream.online
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#streamonline
type StreamOnline struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	Type                 string    `json:"live"`
	StartedAt            time.Time `json:"started_at"`
}

// stream.offline
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#streamoffline
type StreamOffline struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

// user.authorization.grant
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#userauthorizationgrant
type UserAuthorizationGrant struct {
	ClientID  string `json:"client_id"`
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

// user.authorization.revoke
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#userauthorizationrevoke
type UserAuthorizationRevoke struct {
	ClientID  string `json:"client_id"`
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

// user.update
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#userupdate
type UserUpdate struct {
	UserID        string `json:"user_id"`
	UserLogin     string `json:"user_login"`
	UserName      string `json:"user_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Description   string `json:"description"`
}

// user.whisper.message
// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/#userwhispermessage
type UserWhisperMessage struct {
	FromUserID    string `json:"from_user_id"`
	FromUserLogin string `json:"from_user_login"`
	FromUserName  string `json:"from_user_name"`
	ToUserID      string `json:"to_user_id"`
	ToUserLogin   string `json:"to_user_login"`
	ToUserName    string `json:"to_user_name"`
	WhiperID      string `json:"whisper_id"`
	Whisper       struct {
		Text string `json:"text"`
	}
}
