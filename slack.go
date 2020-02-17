package main

//import github.com/nlopes/slack
import (
	"fmt"
	"reflect"
	"time"

	"github.com/nlopes/slack"
)

var (
	appid        = "AU2F4FN3E"
	userid       = "DRN4CMQ77"
	url          = "*********"
	access_token = "*********"
	channelID    = "CTT1B5F33"
)

func SlackTest() {
	api := slack.New(access_token, slack.OptionDebug(true))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}

	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, channel := range channels {
		fmt.Println(channel.Name)

		// channel is of type conversation & groupConversation
		// see all available methods in `conversation.go`
	}

	user, err := api.GetUserInfo("URQA7QJLE")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)

	///
	//webhook
	var receivedPayload slack.WebhookMessage
	payload := &slack.WebhookMessage{
		Text: "golang slack test",
		Attachments: []slack.Attachment{
			{
				Text: "웹훅 잘 되넹 :: " + time.Now().String(),
			},
		},
	}

	err = slack.PostWebhook(url, payload)

	if err != nil {
		// t.Errorf("Expected not to receive error: %s", err)
	}

	if !reflect.DeepEqual(payload, &receivedPayload) {
		// t.Errorf("Payload did not match\nwant: %#v\n got: %#v", payload, receivedPayload)
	}

	///
	//send message
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage(channelID, slack.MsgOptionText("Hello, World", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
