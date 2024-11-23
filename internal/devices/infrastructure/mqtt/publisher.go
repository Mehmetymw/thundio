package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTPublisher struct {
	client mqtt.Client
	topic  string
}

func NewMQTTPublisher(broker, topic, clientID string) (*MQTTPublisher, error) {
	clientID = fmt.Sprintf("%s-%d", clientID, time.Now().UnixNano())

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	return &MQTTPublisher{
		client: client,
		topic:  topic,
	}, nil
}

func (mp *MQTTPublisher) PublishMessage(message string) error {
	token := mp.client.Publish(mp.topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish message: %v", token.Error())
	}

	log.Printf("Message published to topic %s: %s", mp.topic, message)
	return nil
}

func (mp *MQTTPublisher) Close() {
	mp.client.Disconnect(250)
}
