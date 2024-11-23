package mqtt

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SensorData struct {
	DeviceName  string  `json:"deviceName"`
	DeviceType  string  `json:"deviceType"`
	SensorValue float64 `json:"sensorValue"`
	Timestamp   string  `json:"timestamp"`
}

type MQTTSubscriber struct {
	client mqtt.Client
	topic  string
}

func NewMQTTSubscriber(broker, topic, clientID string) (*MQTTSubscriber, error) {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	return &MQTTSubscriber{
		client: client,
		topic:  topic,
	}, nil
}

func (mp *MQTTSubscriber) Subscribe() {
	token := mp.client.Subscribe(mp.topic, 0, func(c mqtt.Client, m mqtt.Message) {
		var sensorData SensorData
		err := json.Unmarshal(m.Payload(), &sensorData)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return
		}
		log.Printf("Received message: Device Name: %s, Type: %s, Sensor Value: %f, Timestamp: %s",
			sensorData.DeviceName, sensorData.DeviceType, sensorData.SensorValue, sensorData.Timestamp)

	})
	token.Wait()

	if token.Error() != nil {
		log.Fatalf("Failed to subscribe: %v", token.Error())
	}
}

func (mc *MQTTSubscriber) Close() {
	mc.client.Disconnect(250)
}
