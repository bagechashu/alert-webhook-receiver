package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"
)

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
		log.Print("env DING_ROBOT_TOKEN or DING_ROBOT_SECRET not found")
		return ""
	}
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	sign := sign(timestamp, secret)

	// log.Printf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s \n", token, timestamp, sign)

	return fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s", token, timestamp, sign)
}
