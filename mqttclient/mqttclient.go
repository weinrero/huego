package mqttclient

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var c mqtt.Client = nil

func Client() mqtt.Client {
	if c == nil {
		createClient()
	}
	return c
}

func Subscribe(topic string, handler mqtt.MessageHandler) {
	Client().Subscribe(topic, 0, handler).Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

func createClient() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://huegoopi.fritz.box:1883")
	opts.SetClientID("huego")
	// opts.SetUsername("user")
	// opts.SetPassword("password")
	opts.SetOrderMatters(false)
	opts.SetDefaultPublishHandler(func(mqtt.Client, mqtt.Message) {})
	opts.SetOnConnectHandler(connectHandler)
	opts.OnConnectionLost = connectLostHandler
	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Could not connect to broker: %v", token.Error())
		panic(token.Error())
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
	panic(err)
}
