package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"weinrero/huego/mqttclient"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	keepAlive := make(chan os.Signal)
	signal.Notify(keepAlive, os.Interrupt, syscall.SIGTERM)

	mqttclient.Client()
	mqttclient.Subscribe("test/#", msgHandler)

	<-keepAlive
}

var msgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
