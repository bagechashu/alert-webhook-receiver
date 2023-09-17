package message

import (
	"bytes"
	"fmt"

	"github.com/bagechashu/alert-webhook-receiver/medium"
)

type Raw struct {
	Body []byte
}

func (raw Raw) ConvertToDingMarkdown() (markdown medium.DingTalkMarkdown, err error) {
	var buffer bytes.Buffer
	buffer.WriteString(string(raw.Body))

	markdown = medium.NewDingTalkMarkdown()
	markdown.SetTitle(fmt.Sprintln("Cloud Resource Alert"))
	markdown.SetText(buffer.String())

	return
}
