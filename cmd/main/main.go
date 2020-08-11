package main

import (
	"log"

	"github.com/elvisgastelum/devsearchbot"
)

var (
	bot devsearchbot.Bot = devsearchbot.NewDevSearchBot()
)

func main() {
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
