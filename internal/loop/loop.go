package loop

import (
	"github.com/jwoglom/macos-findmy-mqtt/internal/homeassistant"
	"github.com/jwoglom/macos-findmy-mqtt/internal/mqttconn"
)

func MainLoop(conn *mqttconn.MqttConn) {
	homeassistant.UpdateConnectivity(conn)
	homeassistant.UpdateOnline(conn)
}
