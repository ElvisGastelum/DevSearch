package main

import (
	"fmt"
	"log"
	"dev-search/slackbot"
)

func init() {
	fmt.Println("Desde el paquete principal")
}

func main() {
	slackbot.HelloWorld()
	var user slackbot.User
	user.Name = "Juanito"

	user2 := slackbot.User{ Name: "Panchito" }

	log.Println(user.Name)
	log.Println(user2.Name)
}