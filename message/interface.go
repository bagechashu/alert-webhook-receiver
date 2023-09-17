package message

import (
	"github.com/bagechashu/alert-webhook-receiver/medium"
	jsoniter "github.com/json-iterator/go"
)

// enhance standard JSON library
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Message interface {
	ConvertToDingMarkdown() (markdown medium.DingTalkMarkdown, err error)
}
