package config

import "flag"

func init() {
	var port int
	var secret bool

	flag.IntVar(&port, "port", 9000, "server port, default is [9000]")
	flag.BoolVar(&secret, "secret", false, "request with secretkey, default is [false]")
	flag.Parse()

	Server.Port = port
	Server.SecretRequest = secret
}
