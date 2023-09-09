package message

import (
	"github.com/bagechashu/alert-webhook-receiver/medium/dingtalk"
	jsoniter "github.com/json-iterator/go"
)

// enhance standard JSON library
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Message interface {
	ConvertToDingMarkdown() (markdown dingtalk.DingTalkMarkdown, err error)
}
