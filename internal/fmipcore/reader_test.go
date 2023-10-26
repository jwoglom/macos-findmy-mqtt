package fmipcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadItems(t *testing.T) {
	bytes := []byte(`[
		{
		  "address" : {
			"locality" : "New York",
			"stateCode" : "NY",
			"subAdministrativeArea" : "Kings County",
			"streetName" : "Street St",
			"countryCode" : "US",
			"administrativeArea" : "New York",
			"areaOfInterest" : [
			  "Long Island"
			],
			"label" : "123 Street St",
			"formattedAddressLines" : [
			  "123 Street St",
			  "Brooklyn, NY  11111",
			  "United States"
			],
			"fullThroroughfare" : "123 Street St",
			"country" : "United States",
			"mapItemFullAddress" : "123 Street St, Brooklyn, NY  11111",
			"streetAddress" : "123"
		  },
		  "name" : "Case",
		  "groupIdentifier" : "4577AEC4-C5BF-44EF-BE0B-F5925604C796",
		  "productIdentifier" : "FA0AD98D-0239-4122-B0EE-1BCC01FC9893",
		  "productType" : {
			"productInformation" : {
			  "requiresAdditionalConnectionTime" : false,
			  "antennaPower" : 4,
			  "modelName" : "AirPods Pro (2nd generation)",
			  "defaultHeroIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8213-0\/online-infobox__2x.png",
			  "requiresAudioSafetyAlert" : false,
			  "defaultHeroIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8213-0\/online-infobox.png",
			  "defaultListIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8213-0\/online-sourcelist__2x.png",
			  "vendorIdentifier" : 76,
			  "productIdentifier" : 8213,
			  "defaultListIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8213-0\/online-sourcelist.png",
			  "appBundleIdentifier" : "",
			  "defaultHeroIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8213-0\/online-infobox__3x.png",
			  "defaultListIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8213-0\/online-sourcelist__3x.png",
			  "manufacturerName" : "Apple"
			},
			"type" : "hawkeye"
		  },
		  "safeLocations" : [
			{
			  "type" : 1,
			  "address" : {
				"countryCode" : "US",
				"formattedAddressLines" : [
				  "234 Street St",
				  "Brooklyn, NY  11111",
				  "United States"
				],
				"fullThroroughfare" : "234 Street St",
				"subAdministrativeArea" : "Kings County",
				"mapItemFullAddress" : "234 Street St, Brooklyn, NY  11111",
				"administrativeArea" : "New York",
				"streetAddress" : "234",
				"locality" : "New York",
				"areaOfInterest" : [
				  "Long Island"
				],
				"country" : "United States",
				"stateCode" : "NY",
				"label" : "234 Street St",
				"streetName" : "Street St"
			  },
			  "location" : {
				"altitude" : 0,
				"locationFinished" : true,
				"positionType" : "safeLocation",
				"latitude" : 40.1111111111111,
				"verticalAccuracy" : 80,
				"longitude" : -73.99999999999999,
				"isOld" : true,
				"floorLevel" : -1,
				"timeStamp" : 1698289517185,
				"horizontalAccuracy" : 80,
				"isInaccurate" : false
			  },
			  "approvalState" : 1,
			  "name" : "Home",
			  "identifier" : "DA1DBAE9-18BC-4221-9BFF-2B66C8F55128"
			}
		  ],
		  "owner" : "owner@localhost",
		  "location" : {
			"horizontalAccuracy" : 25.945370610329167,
			"latitude" : 40.666666666666666,
			"verticalAccuracy" : -1,
			"isInaccurate" : false,
			"altitude" : -1,
			"longitude" : -73.99999999999999,
			"isOld" : false,
			"positionType" : "ownedDeviceLocation",
			"timeStamp" : 1698288601000,
			"floorLevel" : -1,
			"locationFinished" : true
		  },
		  "isAppleAudioAccessory" : true,
		  "isFirmwareUpdateMandatory" : false,
		  "crowdSourcedLocation" : {
			"latitude" : 40.777777777777777,
			"longitude" : -73.888888888888888,
			"isInaccurate" : false,
			"horizontalAccuracy" : 25.945370610329167,
			"timeStamp" : 1698288601000,
			"locationFinished" : true,
			"floorLevel" : -1,
			"altitude" : -1,
			"verticalAccuracy" : -1,
			"isOld" : false,
			"positionType" : "ownedDeviceLocation"
		  },
		  "identifier" : "E72B636C-945E-4E1E-AC38-1E9544D03B0E",
		  "capabilities" : 878,
		  "lostModeMetadata" : null,
		  "partInfo" : {
			"type" : "case",
			"name" : "Case"
		  },
		  "batteryStatus" : 1,
		  "serialNumber" : "KKK7777777",
		  "role" : {
			"identifier" : 999,
			"emoji" : "😀",
			"name" : "Custom Name"
		  },
		  "systemVersion" : "51.11.1"
		}
	  ]`)
	items, err := readItems(bytes)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(items))

	i := items[0]
	assert.Equal(t, "Case", i.Name)
	assert.Equal(t, FmItemAddress{
		Label:              "123 Street St",
		MapItemFullAddress: "123 Street St, Brooklyn, NY  11111",
		FormattedAddressLines: []string{
			"123 Street St",
			"Brooklyn, NY  11111",
			"United States",
		},
	}, i.Address)
	assert.Equal(t, FmItemProductType{
		Type: "hawkeye",
		ProductInformation: FmItemProductTypeInfo{
			ModelName: "AirPods Pro (2nd generation)",
		},
	}, i.ProductType)
	assert.Equal(t, FmItemLocation{
		Latitude:     40.666666666666666,
		Longitude:    -73.99999999999999,
		Timestamp:    1698288601000,
		IsInaccurate: false,
		IsOld:        false,
	}, i.Location)
	assert.Equal(t, FmItemLocation{
		Latitude:     40.777777777777777,
		Longitude:    -73.888888888888888,
		Timestamp:    1698288601000,
		IsInaccurate: false,
		IsOld:        false,
	}, i.CrowdSourcedLocation)
	assert.Equal(t, "E72B636C-945E-4E1E-AC38-1E9544D03B0E", i.Identifier)
	assert.Equal(t, "4577AEC4-C5BF-44EF-BE0B-F5925604C796", i.GroupIdentifier)
	assert.Equal(t, "FA0AD98D-0239-4122-B0EE-1BCC01FC9893", i.ProductIdentifier)
	assert.Equal(t, "KKK7777777", i.SerialNumber)
	assert.Equal(t, "51.11.1", i.SystemVersion)
	assert.Equal(t, FmItemPartInfo{
		Type: "case",
		Name: "Case",
	}, i.PartInfo)
}
