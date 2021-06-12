package server

import "github.com/Rhodanthe1116/go-gin-boilerplate/config"

func Init(service string) {
	config := config.GetConfig()
	r := NewRouter(service)
	r.Run(config.GetString("server.port"))
}
