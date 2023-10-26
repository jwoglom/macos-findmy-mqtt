package homeassistant

import (
	"github.com/jwoglom/macos-findmy-mqtt/internal/mqttconn"
)

func UpdateConnectivity(conn *mqttconn.MqttConn) error {
	return conn.PublishJson(ConnectivityTopic, map[string]string{
		"name":    "Find My (macos-findmy-mqtt)",
		"uniq_id": BaseMqttTopic + "_connectivity",
		"stat_t":  StatusTopic,
		"dev_cla": "connectivity",
		"pl_on":   "online",
		"pl_off":  "offline",
	})
}

func UpdateOnline(conn *mqttconn.MqttConn) error {
	return conn.PublishString(StatusTopic, "online")
}
