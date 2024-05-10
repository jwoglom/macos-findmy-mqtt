package homeassistant

import (
	"flag"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
)

var topic = flag.String("topic", "findmy-mqtt", "MQTT default topic name")

var BaseMqttTopic string
var ConnectivityTopic string
var StatusTopic string
var BaseDeviceTrackerTopic string

func init() {
	BaseMqttTopic = *topic
	ConnectivityTopic = "homeassistant/binary_sensor/" + BaseMqttTopic + "/connectivity/config"
	StatusTopic = BaseMqttTopic + "/status"
	BaseDeviceTrackerTopic = "homeassistant/device_tracker/"
}

func ItemConfigTopic(item fmipcore.FmItem) string {
	return BaseDeviceTrackerTopic + BaseMqttTopic + "_" + item.Identifier + "/config"
}

func ItemStateTopic(item fmipcore.FmItem) string {
	return BaseMqttTopic + "/" + item.Identifier + "/state"
}
func ItemAttributesTopic(item fmipcore.FmItem) string {
	return BaseMqttTopic + "/" + item.Identifier + "/attributes"
}
