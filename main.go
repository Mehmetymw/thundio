package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mehmetymw/thundio/configs"
	"github.com/mehmetymw/thundio/internal/devices/application"
	"github.com/mehmetymw/thundio/internal/devices/db"
	"github.com/mehmetymw/thundio/internal/devices/infrastructure/mqtt"
	"github.com/mehmetymw/thundio/internal/devices/infrastructure/repository"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Error creating config: %v\n", err)
	}

	mqttPublisher, err := mqtt.NewMQTTPublisher(config.MQTTBroker, config.DeviceTopic, config.MQTTClientID)
	if err != nil {
		log.Fatalf("Error creating MQTT publisher: %v\n", err)
	}

	db.RunMigrations(config.DatabaseUrl)
	dbConn, err := db.InitDB(config.DatabaseUrl)
	if err != nil {
		log.Fatalf("db cannot init: %v\n", err)
	}

	deviceRepo := repository.NewDeviceRepository(dbConn)
	deviceService := application.NewDeviceService(deviceRepo, mqttPublisher)

	mqttSubscriber, err := mqtt.NewMQTTSubscriber(config.MQTTBroker, config.DeviceTopic, config.MQTTClientID)
	if err != nil {
		log.Fatalf("Error creating MQTT subscriber: %v\n", err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := mqttPublisher.PublishMessage("Device registration started")
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		}
	}()

	go func() {
		mqttSubscriber.Subscribe()
	}()

	deviceName := "Sensor-" + fmt.Sprintf("%d", 1)
	deviceType := "Sensor"

	go func() {
		device, err := deviceService.RegisterDevice(deviceName, deviceType)
		if err != nil {
			log.Printf("Failed to register device: %v", err)
		} else {
			log.Printf("Device %s registered successfully with ID: %d", device.Name, device.ID)

			err := mqttPublisher.PublishMessage(fmt.Sprintf("Device registered: %s", device.Name))
			if err != nil {
				log.Printf("Failed to publish message after device registration: %v", err)
			}
		}
	}()

	<-stopChan
	log.Println("Shutting down gracefully...")
}
