package main

import (
	"flag"
	"fmt"
	"os/exec"
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
var openApp = flag.Bool("openApp", true, "whether to open the Find My MacOS app if not detected running")

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

func findMyAppOpen() bool {
	cmd := exec.Command("sh", "-c", "ps aux | grep '[F]indMy.app'")
	if _, err := cmd.CombinedOutput(); err != nil {
		_, ok := err.(*exec.ExitError)
		return !ok
	}
	return true
}

func runMqtt() {

	if *openApp {
		if !findMyAppOpen() {
			fmt.Println("Opening FindMy.app")
			exec.Command("open", "/System/Applications/FindMy.app").Start()
		} else {
			fmt.Println("FindMy.app is already running")
		}
	}

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
