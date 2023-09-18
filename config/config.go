package config

import (
	"os"

	"github.com/joho/godotenv"
)

type dingRobot struct {
	Token  string
	Secret string
}

type server struct {
	Port          int
	SecretRequest bool
	SecretKey     string
}

var (
	DingRobot dingRobot
)

var (
	Server server
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

	Server.SecretKey = os.Getenv("WEBHOOK_SECRET_KEY")
}
