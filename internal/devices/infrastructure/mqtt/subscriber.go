package mqtt

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

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
		log.Printf("Message received: %s", m.Payload())
	})
	token.Wait()

	if token.Error() != nil {
		log.Fatalf("Failed to subscribe: %v", token.Error())
	}

	select {}

}

func (mc *MQTTSubscriber) Close() {
	mc.client.Disconnect(250)
}
