package message

import (
	"bytes"
	"fmt"
)

type Raw struct {
	Body []byte
}

func (raw Raw) ConvertToMarkdown() (title, markdown string, err error) {
	var buffer bytes.Buffer
	buffer.WriteString(string(raw.Body))

	title = fmt.Sprintln("Cloud Resource Alert")
	markdown = buffer.String()

	return
}
