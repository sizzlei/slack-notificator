package slacknotificator 


import (
	"github.com/slack-go/slack"
)

func GetClient(token string) *slack.Client {
	return slack.New(token)
}

func CreateDMChannel(api *slack.Client, users string) (*string, error) {
	var chanId *slack.Channel
	var err error
	chanId,_,_,err = api.OpenConversation(&slack.OpenConversationParameters{
		ReturnIM : false,
		Users: []string{users},
	})
	if err != nil {
		return nil, err
	}

	return &chanId.GroupConversation.Conversation.ID, nil
}

func SendMessage(api *slack.Client, chanId string, msg string) error {
	_, _, err := api.PostMessage(chanId,slack.MsgOptionText(msg, false))
	if err != nil {
		return err
	}

	return nil
}