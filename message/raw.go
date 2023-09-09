package message

import (
	"bytes"
	"fmt"

	"github.com/bagechashu/alert-webhook-receiver/medium/dingtalk"
)

type Raw struct {
	Body []byte
}

func (raw Raw) ConvertToDingMarkdown() (markdown dingtalk.DingTalkMarkdown, err error) {
	var buffer bytes.Buffer
	buffer.WriteString(string(raw.Body))

	markdown = dingtalk.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &dingtalk.Markdown{
			Title: fmt.Sprintln("Cloud Resource Alert"),
			Text:  buffer.String(),
		},
		At: &dingtalk.At{
			IsAtAll: false,
		},
	}
	return
}
