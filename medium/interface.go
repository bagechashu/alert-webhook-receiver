package medium

import (
	"fmt"

	"github.com/bagechashu/alert-webhook-receiver/config"
	"github.com/bagechashu/alert-webhook-receiver/message"
)

type Medium interface {
	Send() (err error)
}

// 根据 msgMedium 初始化媒体结构体
func CreateInterfaceMedium(msgMedium string, msg message.Message) (med Medium, err error) {
	switch msgMedium {
	case "ding":
		var title, text string
		title, text, err = msg.ConvertToMarkdown()
		if err != nil {
			return
		}
		// new markdown for dingtalk message body.
		markdown := DingTalkMarkdown{
			MsgType: "markdown",
			At: &At{
				IsAtAll:   false,
				AtMobiles: []string{},
			},
			Markdown: &Markdown{
				Title: title,
				Text:  text,
			},
		}

		// set msgMedium to dingtalk
		med = DingRobot{
			Token:   config.DingRobot.Token,
			Secret:  config.DingRobot.Secret,
			ReqBody: markdown,
		}
	default:
		err = fmt.Errorf("message medium error: %s. Only support [ding]", msgMedium)
		return
	}
	return
}
