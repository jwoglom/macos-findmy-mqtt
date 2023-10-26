package loop

import (
	"fmt"
	"time"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
	"github.com/jwoglom/macos-findmy-mqtt/internal/homeassistant"
	"github.com/jwoglom/macos-findmy-mqtt/internal/mqttconn"
)

func MainLoop(conn *mqttconn.MqttConn, freq int) error {
	for {
		homeassistant.UpdateConnectivity(conn)
		homeassistant.UpdateOnline(conn)

		items, err := fmipcore.ReadItems()
		if err != nil {
			return err
		}
		for _, item := range items {
			homeassistant.UpdateItem(conn, item)
		}

		if freq == 0 {
			return nil
		}
		fmt.Printf("sleeping for %d seconds\n", freq)
		time.Sleep(time.Duration(freq) * time.Second)
	}
}
