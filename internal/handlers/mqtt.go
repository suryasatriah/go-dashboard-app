package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/surysatriah/go-dashboard-app/internal/handlers/controller"
	"github.com/surysatriah/go-dashboard-app/internal/model"
	"github.com/surysatriah/go-dashboard-app/pkg"
)

// MQTT Credentials
var (
	broker   string = pkg.GetDotEnvVariable("MQTT_BROKER")
	port, _         = strconv.Atoi(pkg.GetDotEnvVariable("MQTT_PORT"))
	clientID string = pkg.GetDotEnvVariable("MQTT_CLIENT_ID")
	username string = pkg.GetDotEnvVariable("MQTT_USERNAME")
	password string = pkg.GetDotEnvVariable("MQTT_PASSWORD")
	topic    string = "go-mqtt/sample"
)

var messageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := model.Payload{}
	err := json.Unmarshal(msg.Payload(), &payload)

	if err != nil {
		log.Printf("Failed to parse MQTT payload: %v", err)
		return
	}

	controller.InsertPayload(payload)
}

var connectedHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("MQTT connection lost: %v", err)
}

func ConnectMqtt() {
	options := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port)).
		SetClientID(clientID).SetUsername(username).SetPassword(password)

	options.SetDefaultPublishHandler(messageHandler)
	options.SetOnConnectHandler(connectedHandler)
	options.SetConnectionLostHandler(connectionLostHandler)

	client := mqtt.NewClient(options)

	// Connect to the MQTT Broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to the MQTT topic
	subscribeMqtt(client, topic)

}

func subscribeMqtt(client mqtt.Client, topic string) {

	// Subscribe to the MQTT topic
	token := client.Subscribe(topic, 1, messageHandler)
	token.Wait()

	// Listen for SIGINT and SIGTERM signals to gracefully exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Waiting for SIGINT or SIGTERM signal...")
	<-sigCh

	// Unsubscribe from the MQTT topic and disconnect from the MQTT broker
	client.Unsubscribe(topic)
	client.Disconnect(250)

}
