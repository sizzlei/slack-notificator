package slacknotificator 


import (
	"github.com/slack-go/slack"
	"encoding/json"
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
	_, _, err := api.Client.PostMessage(*api.ChanId,slack.MsgOptionText(msg, false),slack.MsgOptionAsUser(false))
	if err != nil {
		return err
	}

	return nil
}

func (api *Slackapi) SendAttachment(att slack.Attachment) error {
	_, _, err := api.Client.PostMessage(*api.ChanId,slack.MsgOptionText("", false),slack.MsgOptionAttachments(att),slack.MsgOptionAsUser(false))
	if err != nil {
		return err
	}

	return nil
}

func CreateAttachement(jsonString string) (slack.Attachment, error) {
	var r slack.Attachment
	err := json.Unmarshal([]byte(jsonString),&r)
	if err != nil {
		return r,err
	}

	return r, nil
} 

func (api *Slackapi) GetMemberId(email string) (*string, error){
	user, err := api.Client.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	memberId := user.ID

	return &memberId, nil
}