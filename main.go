package main

import (
	"flag"
	"fmt"
	"os"
    "os/signal"
    "syscall"

	"github.com/Rhodanthe1116/go-gin-boilerplate/config"
	"github.com/Rhodanthe1116/go-gin-boilerplate/db"
	"github.com/Rhodanthe1116/go-gin-boilerplate/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        db.Clear()
        os.Exit(1)
    }()
	server.Init()
}
