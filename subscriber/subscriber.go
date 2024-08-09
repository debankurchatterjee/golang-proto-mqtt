package main

import (
	"fmt"
	v1alpha1 "golang-proto-mqtt/gen/go"
	"log"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_subscriber")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to broker: %v", token.Error())
	}
	// Define a message handler
	var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		// Deserialize the message
		data := &v1alpha1.SensorData{}
		if err := proto.Unmarshal(msg.Payload(), data); err != nil {
			log.Printf("Error deserializing message: %v", err)
			return
		}
		// Print the received message
		fmt.Printf("Received message: %s\n", data.String())
	}
	// Subscribe to the topic
	if token := client.Subscribe("sensors/data", 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic: %v", token.Error())
	}
	fmt.Println("Subscribed to topic")
	select {} // Keep the program running
}
