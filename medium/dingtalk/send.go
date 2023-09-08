package dingtalk

import (
	"bytes"
	"fmt"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func sign(timestamp int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", timestamp, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	hmacCode := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(hmacCode)
}

func robotURL() string {
	token := os.Getenv("DING_ROBOT_TOKEN")
	secret := os.Getenv("DING_ROBOT_SECRET")
	if token == "" || secret == "" {
		log.Println("error: env DING_ROBOT_TOKEN or DING_ROBOT_SECRET not found")
		return ""
	}
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	sign := sign(timestamp, secret)

	// log.Printf("info: ding webhook url: [ https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s ]\n", token, timestamp, sign)

	return fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", token, timestamp, sign)
}

func Send(markdown DingTalkMarkdown) (err error) {
	var webhook = robotURL()
	if err != nil || webhook == "" {
		return
	}

	reqbody, err := json.Marshal(markdown)
	if err != nil {
		return
	}
	log.Printf("info: dingtalk post data: %s", reqbody)

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(reqbody))
	if err != nil {
		log.Println("error: alarm request create error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error: alarm request send error:", err)
		return
	}
	defer resp.Body.Close()

	/* 处理 dingtalk response
	https://open.dingtalk.com/document/robots/custom-robot-access#title-7ur-3ok-s1a
	{
	"errcode": 400102,
	"errmsg": "description:机器人已经停用或者未启用;solution:请让企业管理员前往开放平台后台启用对应机器人 :https://open-dev.dingtalk.com/#/"
	}
	{
	"errcode": 0,
	"errmsg": "ok"
	}
	*/

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error: response body reading error:", err)
		return
	}

	log.Printf("info: dingtalk response: %s", body)

	var jsonResp map[string]interface{}
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		log.Println("error: json unmarshaling error:", err)
		return
	}

	errCode := jsonResp["errcode"].(float64)
	errMsg := jsonResp["errmsg"].(string)
	if errCode != 0 || errMsg != "ok" {
		log.Printf("error: dingtalk request error: %f - %s\n", errCode, errMsg)
		return
	}
	return
}
