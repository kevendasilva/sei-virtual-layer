package broker

import (
	"fmt"
	"log"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func onConnect(client MQTT.Client) {
	log.Println("Connected to MQTT broker!")
}

func onConnectionLost(client MQTT.Client, err error) {
	log.Printf("Connection lost: %v\n", err)
}

func CreateMQTTClient() MQTT.Client {
	err := godotenv.Load()
	env := os.Getenv("ENVIRONMENT")

	if env == "development" {
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	brokerHost := os.Getenv("MQTT_BROKER_ADDRESS")
	brokerPort := os.Getenv("MQTT_BROKER_PORT")
	brokerUsername := os.Getenv("MQTT_USERNAME")
	brokerPassword := os.Getenv("MQTT_PASSWORD")

	// Configurações do cliente MQTT
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", brokerHost, brokerPort))
	opts.SetClientID("go-MQTT-example")
	opts.SetUsername(brokerUsername)
	opts.SetPassword(brokerPassword)

	log.Printf("Connecting to %s:%s\n", brokerHost, brokerPort)

	opts.OnConnect = onConnect
	opts.OnConnectionLost = onConnectionLost

	// Criação do cliente MQTT
	client := MQTT.NewClient(opts)
	return client
}
