package message

import (
	"bytes"
	"fmt"
	"log"

	"github.com/bagechashu/alert-webhook-receiver/medium/dingtalk"
)

// https://support.huaweicloud.com/usermanual-smn/smn_ug_a9002.html
type SmnReqBody struct {
	Type         string `json:"type"`
	Timestamp    string `json:"timestamp"`     // 消息第一次发送的时间戳
	SubscribeUrl string `json:"subscribe_url"` // 确认订阅消息
	Subject      string `json:"subject"`       // 推送消息标题
	Message      string `json:"message"`       // 消息详情
}

type HuaweiSMN struct {
	Body []byte
}

func (smn HuaweiSMN) ConvertToDingMarkdown() (markdown dingtalk.DingTalkMarkdown, err error) {
	var smnReqBody SmnReqBody
	err = json.Unmarshal(smn.Body, &smnReqBody)
	if err != nil {
		log.Printf("error: unmarshal huaweiSMN data error: %s", err.Error())
		return
	}

	var buffer bytes.Buffer

	if smnReqBody.Type == "SubscriptionConfirmation" {
		buffer.WriteString(fmt.Sprintf("### <font color=\"#08d417\"> %s </font>\n", "订阅确认"))
		buffer.WriteString(fmt.Sprintf("\n##### StartAt: %s\n", smnReqBody.Timestamp))
		buffer.WriteString(fmt.Sprintf("\n##### SubscribeUrl: %s\n", smnReqBody.SubscribeUrl))
		buffer.WriteString(fmt.Sprintf("\n##### Message: \n> %s\n", smnReqBody.Message))
	} else {
		buffer.WriteString(fmt.Sprintf("### <font color=\"#FF0000\"> %s </font> %s\n", "云资源报警", smnReqBody.Type))
		buffer.WriteString(fmt.Sprintf("\n##### StartAt: %s\n", smnReqBody.Timestamp))
		buffer.WriteString(fmt.Sprintf("\n##### Subject: %s\n", smnReqBody.Subject))
		var msg map[string]interface{}
		ok := json.Unmarshal([]byte(smnReqBody.Message), &msg)
		if ok != nil {
			// log.Printf("error: unmarshal huaweiSMN message data error: %s", ok)
			buffer.WriteString(fmt.Sprintf("\n##### Message: \n> %s\n", smnReqBody.Message))
		} else {
			buffer.WriteString(fmt.Sprintln("\n##### Message:"))
			for k, v := range msg {
				buffer.WriteString(fmt.Sprintf("\n> %s: %s\n", k, v))
			}
		}
	}

	markdown = dingtalk.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &dingtalk.Markdown{
			Title: fmt.Sprintln("云资源报警, 请及时查看."),
			Text:  buffer.String(),
		},
		At: &dingtalk.At{
			IsAtAll: false,
		},
	}

	return
}
