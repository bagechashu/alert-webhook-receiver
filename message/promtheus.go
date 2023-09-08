package message

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/bagechashu/alert-webhook-receiver/medium/dingtalk"
)

// from prometheus model
type Alert struct {
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
}

type Notification struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`
}

type Prom struct {
	Body []byte
}

func (prom Prom) ConvertToDingMarkdown() (markdown dingtalk.DingTalkMarkdown) {

	var notification Notification
	err := json.Unmarshal(prom.Body, &notification)
	if err != nil {
		log.Printf("error: unmarshal prom notification data error: %s", err.Error())
		return
	}

	var buffer bytes.Buffer

	var alertsNum int = 0
	for _, alert := range notification.Alerts {
		if alert.Status == "resolved" {
			annotations := alert.Annotations
			buffer.WriteString(fmt.Sprintf("### <font color=\"#08d417\"> %s </font>\n", "恢复通知"))
			buffer.WriteString(fmt.Sprintf("##### %s\n", annotations["summary"]))
			buffer.WriteString(fmt.Sprintf("\n> 状态: %s\n", alert.Status))
			buffer.WriteString(fmt.Sprintf("\n> 级别: %s\n", alert.Labels["severity"]))
			buffer.WriteString(fmt.Sprintf("\n> 时间: %s\n", alert.StartsAt.Local().Format("2006-01-02 15:04:05")))
			buffer.WriteString(fmt.Sprintf("\n> 详情: %s\n", annotations["description"]))
		} else {
			annotations := alert.Annotations
			buffer.WriteString(fmt.Sprintf("### <font color=\"#FF0000\"> %s </font>\n", "告警通知"))
			buffer.WriteString(fmt.Sprintf("##### %s\n", annotations["summary"]))
			buffer.WriteString(fmt.Sprintf("\n> 状态: %s\n", alert.Status))
			buffer.WriteString(fmt.Sprintf("\n> 级别: %s\n", alert.Labels["severity"]))
			buffer.WriteString(fmt.Sprintf("\n> 时间: %s\n", alert.StartsAt.Local().Format("2006-01-02 15:04:05")))
			buffer.WriteString(fmt.Sprintf("\n> 详情: %s\n", annotations["description"]))
			alertsNum += 1
		}
	}

	buffer.WriteString(fmt.Sprintf("---\n##### [当前报警共有 %d 条 ]\n", alertsNum))

	markdown = dingtalk.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &dingtalk.Markdown{
			Title: fmt.Sprintf("您有%d条监控信息, 请及时查看", len(notification.Alerts)),
			Text:  buffer.String(),
		},
		At: &dingtalk.At{
			IsAtAll: false,
		},
	}

	return
}
