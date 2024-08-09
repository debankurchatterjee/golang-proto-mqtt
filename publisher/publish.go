package main

import (
	"encoding/json"
	"fmt"
	v1alpha1 "golang-proto-mqtt/gen/go"
	"log"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_mqtt_client")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to broker: %v", token.Error())
	}

	// Create a SensorData message
	data := &v1alpha1.SensorData{
		Id:          "sensor-1",
		Temperature: 23.5,
		Humidity:    60.0,
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Println(marshal)
	// Serialize the message to binary format
	dataBytes, err := proto.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing message: %v", err)
	}
	// Publish the message to the topic
	token := client.Publish("sensors/data", 0, true, dataBytes)
	fmt.Println("Message published")
	token.Wait()
	time.Sleep(2 * time.Second)
	client.Disconnect(250)
}
