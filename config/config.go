package config

import (
	"os"

	"github.com/joho/godotenv"
)

type dingRobot struct {
	Token  string
	Secret string
}

var (
	DingRobot dingRobot
)

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
		godotenv.Load()
	}
	godotenv.Load(".env." + env)

	DingRobot.Token = os.Getenv("DING_ROBOT_TOKEN")
	DingRobot.Secret = os.Getenv("DING_ROBOT_SECRET")
}
