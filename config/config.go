package config

import "os"

type dingRobot struct {
	Token  string
	Secret string
}

var (
	DingRobot dingRobot
)

func init() {
	DingRobot.Token = os.Getenv("DING_ROBOT_TOKEN")
	DingRobot.Secret = os.Getenv("DING_ROBOT_SECRET")
}
