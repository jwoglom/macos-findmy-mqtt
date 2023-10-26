package homeassistant

import "github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"

const BaseMqttTopic = "findmy-mqtt"
const ConnectivityTopic = "homeassistant/binary_sensor/" + BaseMqttTopic + "/connectivity/config"
const StatusTopic = BaseMqttTopic + "/status"
const BaseDeviceTrackerTopic = "homeassistant/device_tracker/"

func ItemConfigTopic(item fmipcore.FmItem) string {
	return BaseDeviceTrackerTopic + BaseMqttTopic + "_" + item.Identifier + "/config"
}

func ItemStateTopic(item fmipcore.FmItem) string {
	return BaseMqttTopic + "/" + item.Identifier + "/state"
}
func ItemAttributesTopic(item fmipcore.FmItem) string {
	return BaseMqttTopic + "/" + item.Identifier + "/attributes"
}
