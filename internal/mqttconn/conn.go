package mqttconn

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttConn struct {
	client   mqtt.Client
	qos      byte
	retained bool
}

func NewMqttConn(client mqtt.Client, qos byte, retained bool) *MqttConn {
	return &MqttConn{
		client:   client,
		qos:      qos,
		retained: retained,
	}
}

func (c *MqttConn) PublishString(topic, value string) error {
	fmt.Printf("publish: %s = %s\n", topic, value)
	token := c.client.Publish(topic, c.qos, c.retained, value)
	if !token.Wait() {
		return fmt.Errorf("unable to publish message topic=%s value=%#v", topic, value)
	}
	return nil
}

func (c *MqttConn) PublishJson(topic string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.PublishString(topic, string(bytes))
}
