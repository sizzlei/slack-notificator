# slack-notificator
## Overview
Slack-go 패키지를 이용하여 간단한 세팅으로 메세지를 보낼 수 있도록 작성된 패키지 입니다. 

## Installation
패키지 설치 구문은 아래와 같습니다.
```
go get github.com/sizzlei/slack-notificator
```


## Example
```go
package main 

import (
	slack "github.com/sizzlei/slack-notificator"
)


func main() {
	token := "xoxb-{user token}"
	api := slack.GetClient(token)

	err := api.CreateDMChannel({user_member_id})
	if err != nil {
		panic(err)
	}

	data := `
	{
		"blocks": [
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "test"
				}
			}
		]
	}
	`

	att, _ := slack.CreateAttachement(data)

	err = api.SendAttachment(att)
	if err != nil {
		panic(err)
	}
}
```