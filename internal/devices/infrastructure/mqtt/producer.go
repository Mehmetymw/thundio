package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTProducer struct {
	client mqtt.Client
	topic  string
}
