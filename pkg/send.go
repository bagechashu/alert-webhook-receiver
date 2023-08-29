package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bagechashu/alert-webhook-receiver/model"
)

func Send(notification model.Notification) (err error) {
	markdown, webhook, err := transformToMarkdown(notification)
	if err != nil || webhook == "" {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	// log.Printf("DingTalk Post data: %s", data)

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error create alarm request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error send alarm request:", err)
		return
	}
	defer resp.Body.Close()

	// dingtalk response
	// https://open.dingtalk.com/document/robots/custom-robot-access#title-7ur-3ok-s1a
	// {
	// "errcode": 400102,
	// "errmsg": "description:机器人已经停用或者未启用;solution:请让企业管理员前往开放平台后台启用对应机器人 :https://open-dev.dingtalk.com/#/"
	// }
	// {
	// "errcode": 0,
	// "errmsg": "ok"
	// }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	log.Printf("DingTalk Response: %s", body)

	var jsonResp map[string]interface{}
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}

	errCode := jsonResp["errcode"].(float64)
	errMsg := jsonResp["errmsg"].(string)
	if errCode != 0 || errMsg != "ok" {
		log.Printf("Error: %f - %s\n", errCode, errMsg)
		return
	}

	return
}
