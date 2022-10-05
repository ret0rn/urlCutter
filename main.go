package main

import (
	"github.com/ret0rn/urlCutter/service"
)

func main() {
	server, err := service.NewService()
	if err != nil {
		panic(err)
	}
	if err := server.Run(); err != nil {
		panic(err)
	}
}
