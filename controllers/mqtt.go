// controllers/mqtt.go
package controllers

import (
	"log"
	"os"

	"github.com/Miya784/update-status-service/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messageReceived = make(chan struct{})

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	MessageUpdateStatus(msg.Payload())
	messageReceived <- struct{}{}
	// log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func WaitForMessage() {
	<-messageReceived // Blocks until a message is received
}

func DefaultSubscribeHandler(client mqtt.Client) {
	topic := utils.Strip(os.Getenv("MQTT_TOPIC"))
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s", topic)
}

func Subscribe(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s", topic)
}

func Publish(client mqtt.Client) {
	topic := "topic/test"
	log.Printf("Before Publish to topic %s", topic)
	token := client.Publish(topic, 0, false, "Hello, world!")
	token.Wait()
	log.Printf("Publish to topic %s", topic)
}
