package slack

import (
	slackSdk "github.com/slack-go/slack"
	"log"
	"os"
)

type Message struct {
	Title   string
	Content string
}

func getSlackClient() *slackSdk.Client {
	return slackSdk.New(os.Getenv("SLACK_TOKEN"))
}

func SendMessage(message Message) {
	attachment := slackSdk.Attachment{
		Pretext: message.Title,
		Text:    message.Content,
	}

	_, _, err := getSlackClient().PostMessage(
		os.Getenv("SLACK_CHANNEL_NAME"),
		slackSdk.MsgOptionAttachments(attachment),
	)
	if err != nil {
		log.Printf("[SendMessage Error] %s\n", err)
		return
	}
}
