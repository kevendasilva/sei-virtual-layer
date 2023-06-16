package main

import (
	"fmt"
	"log"
	"os"
	"time"

	broker "orchestrator/broker"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	client := broker.CreateMQTTClient()

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	for i := 1; i < 3; i++ {
		topic := fmt.Sprintf("camera/%d", i)

		token := client.Subscribe(topic, 1, handlerMessage)

		if token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}

		log.Printf("Subscribed to topic: %s\n", topic)
	}

	for {
		time.Sleep(10 * time.Millisecond)
	}
}

func handlerMessage(client MQTT.Client, msg MQTT.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	defaultPlate := os.Getenv("DEFAULT_PLATE")

	topic := fmt.Sprintf("%s/result", msg.Topic())

	if token := client.Publish(topic, 1, false, defaultPlate); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Printf("Published message: %s to topic: %s\n", defaultPlate, topic)
}
