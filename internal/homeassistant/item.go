package homeassistant

import (
	"fmt"
	"math"
	"strings"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
	"github.com/jwoglom/macos-findmy-mqtt/internal/mqttconn"
)

func UpdateItem(conn *mqttconn.MqttConn, item fmipcore.FmItem) error {
	name := item.Name
	if item.IsAirTag() {
		name += " (AirTag)"
	}

	fmt.Printf("Found item %s (%s)\n", name, item.Identifier)

	var group *fmipcore.FmItemGroup
	if item.GroupIdentifier != "" {
		// find the base group object name
		group, _ = fmipcore.GetItemGroup(item.GroupIdentifier)
		if group == nil {
			fmt.Printf("unable to find group %s, so skipping object\n", item.GroupIdentifier)
			return nil
		}

		// prepend name of the group (e.g. Left Bud -> AirPods Left Bud)
		name = group.Name + " - " + name
	}

	config := map[string]string{
		"state_topic":           ItemStateTopic(item),
		"json_attributes_topic": ItemAttributesTopic(item),
		"name":                  name,
		"unique_id":             BaseMqttTopic + "_" + item.Identifier,
	}

	err := conn.PublishJson(ItemConfigTopic(item), config)
	if err != nil {
		return err
	}

	state := item.Address.Label

	if safeLocation, ok := itemInSafeLocation(item); ok {
		state = safeLocation.Name
	}

	err = conn.PublishString(ItemStateTopic(item), state)
	if err != nil {
		return err
	}

	attrs := map[string]string{
		"latitude":       fmt.Sprintf("%.8f", item.Location.Latitude),
		"longitude":      fmt.Sprintf("%.8f", item.Location.Longitude),
		"address_label":  item.Address.Label,
		"address_full":   strings.Join(item.Address.FormattedAddressLines, "\n"),
		"last_updated":   fmt.Sprintf("%d", item.Location.Timestamp),
		"serial_number":  item.SerialNumber,
		"system_version": item.SystemVersion,
	}

	if group != nil {
		attrs["group_name"] = group.Name
		attrs["pairing_state"] = group.PairingState(item.Identifier)
	}

	err = conn.PublishJson(ItemAttributesTopic(item), attrs)
	if err != nil {
		return err
	}

	return nil
}

const InSavedLocationThreshold = 0.0025

func itemInSafeLocation(item fmipcore.FmItem) (*fmipcore.FmItemSafeLocation, bool) {
	item_lat := item.Location.Latitude
	item_lng := item.Location.Longitude
	item_label := item.Address.Label

	for _, loc := range item.SafeLocations {
		safe_lat := loc.Location.Latitude
		safe_lng := loc.Location.Longitude
		safe_label := loc.Address.Label

		if math.Abs(item_lat-safe_lat) <= InSavedLocationThreshold && math.Abs(item_lng-safe_lng) <= InSavedLocationThreshold {
			return &loc, true
		}

		if item_label == safe_label {
			return &loc, true
		}

	}
	return nil, false
}
