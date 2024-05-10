package homeassistant

import (
	"flag"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
)

var topic = flag.String("topic", "findmy-mqtt", "MQTT default topic name")

func BaseMqttTopic() string {
	return *topic
}

func ConnectivityTopic() string {
	return "homeassistant/binary_sensor/" + BaseMqttTopic() + "/connectivity/config"
}

func StatusTopic() string {
	return BaseMqttTopic() + "/status"
}

func BaseDeviceTrackerTopic() string {
	return "homeassistant/device_tracker/"
}

func ItemConfigTopic(item fmipcore.FmItem) string {
	return BaseDeviceTrackerTopic() + BaseMqttTopic() + "_" + item.Identifier + "/config"
}

func ItemStateTopic(item fmipcore.FmItem) string {
	return BaseMqttTopic() + "/" + item.Identifier + "/state"
}
func ItemAttributesTopic(item fmipcore.FmItem) string {
	return BaseMqttTopic() + "/" + item.Identifier + "/attributes"
}
