package main

import (
	"github.com/Miya784/update-status-service/controllers"
	"github.com/Miya784/update-status-service/initials"
	_ "github.com/lib/pq"
)

func main() {
	initials.InitDB()
	client := initials.InitialMqttClient(controllers.MessagePubHandler)
	controllers.DefaultSubscribeHandler(client)
	for {
		controllers.WaitForMessage()
	}
}
