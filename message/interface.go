package message

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// enhance standard JSON library
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Message interface {
	ConvertToMarkdown() (title, markdown string, err error)
}

// 根据 msgType 初始化报警源结构体
func CreateInterfaceMessage(msgType string, body []byte) (msg Message, err error) {
	switch msgType {
	case "raw":
		msg = Raw{Body: body}
	case "prom":
		msg = Prom{Body: body}
	case "huaweismn":
		msg = HuaweiSMN{Body: body}
	default:
		err = fmt.Errorf("msgtype type error: %s. Only support [raw/prom/huaweismn]", msgType)
		return
	}
	return
}
