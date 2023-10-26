package loop

import (
	"fmt"
	"time"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
	"github.com/jwoglom/macos-findmy-mqtt/internal/homeassistant"
	"github.com/jwoglom/macos-findmy-mqtt/internal/mqttconn"
)

func MainLoop(conn *mqttconn.MqttConn) error {
	homeassistant.UpdateConnectivity(conn)
	homeassistant.UpdateOnline(conn)

	fmt.Println("Published online status")

	for {
		items, err := fmipcore.ReadItems()
		if err != nil {
			return err
		}
		for _, item := range items {
			homeassistant.UpdateItem(conn, item)
		}

		fmt.Println("sleeping for 60 seconds")
		time.Sleep(60 * time.Second)
	}
}
