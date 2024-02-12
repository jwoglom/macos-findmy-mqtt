package main

import (
	"flag"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jwoglom/macos-findmy-mqtt/internal/httpserver"
	"github.com/jwoglom/macos-findmy-mqtt/internal/loop"
	"github.com/jwoglom/macos-findmy-mqtt/internal/mqttconn"
)

var broker = flag.String("broker", "", "The broker URI. ex: tcp://10.10.1.1:1883")
var user = flag.String("user", "", "The User (optional)")
var password = flag.String("password", "", "The password (optional)")
var id = flag.String("id", "testgoid", "The ClientID (optional)")
var qos = flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
var freq = flag.Int("freq", 30, "The number of seconds between MQTT pushes")
var httpPort = flag.Int("httpPort", 8080, "http port for webserver")

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

func runMqtt() {

	if *broker == "" {
		fmt.Println("no broker provided, only running webserver")
		time.Sleep(365 * 24 * time.Hour)
		return
	}

	client := mqtt.NewClient(buildOpts())
	fmt.Printf("connecting to %s\n", *broker)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	conn := mqttconn.NewMqttConn(client, byte(*qos), false)

	err := loop.MainLoop(conn, *freq)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(1000)
}

func main() {
	flag.Parse()

	go httpserver.Run(*httpPort)

	runMqtt()
}
