package main

import (
	"flag"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var broker = flag.String("broker", "", "The broker URI. ex: tcp://10.10.1.1:1883")
var user = flag.String("user", "", "The User (optional)")
var password = flag.String("password", "", "The password (optional)")
var id = flag.String("id", "testgoid", "The ClientID (optional)")
var qos = flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")

func buildOpts() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetClientID(*id)

	return opts
}

func strPtr(s string) *string {
	return &s
}

func main() {
	flag.Parse()
	client := mqtt.NewClient(buildOpts())
	fmt.Printf("connecting to %s\n", *broker)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer client.Disconnect(1000)
}
