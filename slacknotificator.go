package slacknotificator 


import (
	"github.com/slack-go/slack"
)

type Slackapi struct {
	Client		*slack.Client
	ChanId 		*string
}

func GetClient(token string) *Slackapi {
	return &Slackapi{
		Client: slack.New(token),
	}
}

func (api *Slackapi) CreateDMChannel(users string) (error) {
	var chanId *slack.Channel
	var err error
	chanId,_,_,err = api.Client.OpenConversation(&slack.OpenConversationParameters{
		ReturnIM : false,
		Users: []string{users},
	})
	if err != nil {
		return err
	}

	api.ChanId = &chanId.GroupConversation.Conversation.ID

	return nil
}

func (api *Slackapi) SendMessage(msg string) error {
	_, _, err := api.Client.PostMessage(*api.ChanId,slack.MsgOptionText(msg, false))
	if err != nil {
		return err
	}

	return nil
}