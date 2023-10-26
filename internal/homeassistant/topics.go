package homeassistant

const BaseMqttTopic = "findmy-mqtt"
const ConnectivityTopic = "homeassistant/binary_sensor/" + BaseMqttTopic + "/connectivity/config"
const StatusTopic = BaseMqttTopic + "/status"
