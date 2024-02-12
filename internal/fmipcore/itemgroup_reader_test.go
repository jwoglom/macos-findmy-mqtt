package fmipcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const AirTagsGroupInput = `[
	{
	  "identifier": "4577AEC4-C5BF-44EF-BE0B-F5925604C796",
	  "lostMetadata": null,
	  "name": "James’s AirPods Pro",
	  "itemPairingStateMap": {
		"GWDJJJJJJJJP": [
		  "p",
		  "C8F2F66A-A2A4-4C6B-B61A-F89735BE0609"
		],
		"GXXXXXXXXXXQ": [
		  "p",
		  "B0D3A871-BE54-4EE5-952B-4E5217A62B7F"
		],
		"KKK7777777": [
		  "p",
		  "E72B636C-945E-4E1E-AC38-1E9544D03B0E"
		]
	  },
	  "items": [
		{
			"productIdentifier" : "CA1E6D06-FB96-4FEA-84ED-A9B96C4C56E5",
			"groupIdentifier" : "4577AEC4-C5BF-44EF-BE0B-F5925604C796",
			"location" : {
			  "horizontalAccuracy" : 27.062871238639026,
			  "latitude" : 40.666666666666666,
			  "timeStamp" : 1698288601000,
			  "longitude" : -73.99999999999999,
			  "isOld" : false,
			  "locationFinished" : true,
			  "floorLevel" : -1,
			  "verticalAccuracy" : -1,
			  "isInaccurate" : false,
			  "altitude" : -1,
			  "positionType" : "ownedDeviceLocation"
			},
			"isAppleAudioAccessory" : true,
			"identifier" : "B0D3A871-BE54-4EE5-952B-4E5217A62B7F",
			"lostModeMetadata" : null,
			"serialNumber" : "GXXXXXXXXXXQ",
			"address" : {
			  "stateCode" : "NY",
			  "country" : "United States",
			  "administrativeArea" : "New York",
			  "locality" : "New York",
			  "streetAddress" : "123",
			  "streetName" : "Street St",
			  "label" : "123 Street St",
			  "countryCode" : "US",
			  "mapItemFullAddress" : "123 Street St, Brooklyn, NY  11111",
			  "areaOfInterest" : [
				"Long Island"
			  ],
			  "formattedAddressLines" : [
				"123 Street St",
				"Brooklyn, NY  11111",
				"United States"
			  ],
			  "fullThroroughfare" : "123 Street St",
			  "subAdministrativeArea" : "Kings County"
			},
			"safeLocations" : [
			  {
				"address" : {
				  "areaOfInterest" : [
					"Long Island"
				  ],
				  "formattedAddressLines" : [
					"234 Street St",
					"Brooklyn, NY  11111",
					"United States"
				  ],
				  "mapItemFullAddress" : "234 Street St, Brooklyn, NY  11111",
				  "subAdministrativeArea" : "Kings County",
				  "administrativeArea" : "New York",
				  "country" : "United States",
				  "stateCode" : "NY",
				  "streetAddress" : "234",
				  "locality" : "New York",
				  "fullThroroughfare" : "234 Street St",
				  "countryCode" : "US",
				  "streetName" : "Street St",
				  "label" : "234 Street St"
				},
				"location" : {
				  "floorLevel" : -1,
				  "timeStamp" : 1698289517185,
				  "isOld" : true,
				  "isInaccurate" : false,
				  "latitude" : 40.1111111111111,
				  "longitude" : -73.888888888888888,
				  "positionType" : "safeLocation",
				  "locationFinished" : true,
				  "verticalAccuracy" : 80,
				  "altitude" : 0,
				  "horizontalAccuracy" : 80
				},
				"identifier" : "ED8617E8-17E8-4519-9C91-4A2C420D17B5",
				"type" : 1,
				"name" : "Home",
				"approvalState" : 1
			  }
			],
			"crowdSourcedLocation" : {
			  "positionType" : "ownedDeviceLocation",
			  "altitude" : -1,
			  "latitude" : 40.77777777777,
			  "longitude" : -73.888888888888888,
			  "horizontalAccuracy" : 27.062871238639026,
			  "verticalAccuracy" : -1,
			  "isInaccurate" : false,
			  "isOld" : false,
			  "timeStamp" : 1698288601000,
			  "floorLevel" : -1,
			  "locationFinished" : true
			},
			"name" : "Left Bud",
			"role" : {
			  "emoji" : "😀",
			  "identifier" : 999,
			  "name" : "Custom Name"
			},
			"productType" : {
			  "productInformation" : {
				"defaultListIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist.png",
				"appBundleIdentifier" : "",
				"defaultListIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__2x.png",
				"productIdentifier" : 8212,
				"requiresAdditionalConnectionTime" : false,
				"defaultListIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__3x.png",
				"antennaPower" : 4,
				"vendorIdentifier" : 76,
				"defaultHeroIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__3x.png",
				"defaultHeroIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__2x.png",
				"defaultHeroIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox.png",
				"modelName" : "AirPods Pro (2nd generation)",
				"requiresAudioSafetyAlert" : true,
				"manufacturerName" : "Apple"
			  },
			  "type" : "hawkeye"
			},
			"systemVersion" : "80.23.144",
			"capabilities" : 814,
			"batteryStatus" : 0,
			"partInfo" : {
			  "type" : "leftBud",
			  "name" : "Left Bud",
			  "symbol" : "l.circle.fill"
			},
			"isFirmwareUpdateMandatory" : false,
			"owner" : "owner@localhost"
		},{
			"productType" : {
			  "type" : "hawkeye",
			  "productInformation" : {
				"manufacturerName" : "Apple",
				"defaultHeroIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox.png",
				"requiresAudioSafetyAlert" : true,
				"requiresAdditionalConnectionTime" : false,
				"appBundleIdentifier" : "",
				"productIdentifier" : 8212,
				"vendorIdentifier" : 76,
				"defaultHeroIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__3x.png",
				"defaultListIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist.png",
				"defaultHeroIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__2x.png",
				"modelName" : "AirPods Pro (2nd generation)",
				"defaultListIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__2x.png",
				"defaultListIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__3x.png",
				"antennaPower" : 4
			  }
			},
			"crowdSourcedLocation" : {
			  "isOld" : false,
			  "locationFinished" : true,
			  "positionType" : "ownedDeviceLocation",
			  "timeStamp" : 1698288601000,
			  "isInaccurate" : false,
			  "longitude" : -73.888888888888888,
			  "verticalAccuracy" : -1,
			  "altitude" : -1,
			  "latitude" : 40.777777777777,
			  "floorLevel" : -1,
			  "horizontalAccuracy" : 27.062871238639026
			},
			"role" : {
			  "emoji" : "😀",
			  "identifier" : 999,
			  "name" : "Custom Name"
			},
			"lostModeMetadata" : null,
			"batteryStatus" : 0,
			"owner" : "owner@localhost",
			"isAppleAudioAccessory" : true,
			"partInfo" : {
			  "type" : "rightBud",
			  "name" : "Right Bud",
			  "symbol" : "r.circle.fill"
			},
			"address" : {
			  "administrativeArea" : "New York",
			  "streetAddress" : "123",
			  "subAdministrativeArea" : "Kings County",
			  "stateCode" : "NY",
			  "country" : "United States",
			  "fullThroroughfare" : "123 Street St",
			  "label" : "123 Street St",
			  "countryCode" : "US",
			  "formattedAddressLines" : [
				"123 Street St",
				"Brooklyn, NY  11111",
				"United States"
			  ],
			  "mapItemFullAddress" : "123 Street St, Brooklyn, NY  11111",
			  "locality" : "New York",
			  "streetName" : "Street St",
			  "areaOfInterest" : [
				"Long Island"
			  ]
			},
			"groupIdentifier" : "4577AEC4-C5BF-44EF-BE0B-F5925604C796",
			"serialNumber" : "GWDJJJJJJJJP",
			"capabilities" : 814,
			"safeLocations" : [
			  {
				"name" : "Home",
				"address" : {
				  "mapItemFullAddress" : "123 Street St, Brooklyn, NY  11111",
				  "fullThroroughfare" : "123 Street St",
				  "countryCode" : "US",
				  "streetAddress" : "123",
				  "administrativeArea" : "New York",
				  "subAdministrativeArea" : "Kings County",
				  "streetName" : "Street St",
				  "formattedAddressLines" : [
					"123 Street St",
					"Brooklyn, NY  11111",
					"United States"
				  ],
				  "locality" : "New York",
				  "stateCode" : "NY",
				  "country" : "United States",
				  "label" : "123 Street St",
				  "areaOfInterest" : [
					"Long Island"
				  ]
				},
				"type" : 1,
				"approvalState" : 1,
				"location" : {
				  "floorLevel" : -1,
				  "timeStamp" : 1698289517185,
				  "isOld" : true,
				  "isInaccurate" : false,
				  "latitude" : 40.1111111111111,
				  "longitude" : -73.888888888888888,
				  "positionType" : "safeLocation",
				  "locationFinished" : true,
				  "verticalAccuracy" : 80,
				  "altitude" : 0,
				  "horizontalAccuracy" : 80
				},
				"identifier" : "ED8617E8-17E8-4519-9C91-4A2C420D17B5"
			  }
			],
			"systemVersion" : "80.23.144",
			"name" : "Right Bud",
			"productIdentifier" : "CA1E6D06-FB96-4FEA-84ED-A9B96C4C56E5",
			"identifier" : "C8F2F66A-A2A4-4C6B-B61A-F89735BE0609",
			"isFirmwareUpdateMandatory" : false,
			"location" : {
			  "isOld" : false,
			  "locationFinished" : true,
			  "verticalAccuracy" : -1,
			  "timeStamp" : 1698288601000,
			  "positionType" : "ownedDeviceLocation",
			  "latitude" : 40.666666666666666,
			  "isInaccurate" : false,
			  "altitude" : -1,
			  "floorLevel" : -1,
			  "longitude" : -73.888888888888888,
			  "horizontalAccuracy" : 27.062871238639026
			}
		},{
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
				  "longitude" : -73.888888888888888,
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
			  "longitude" : -73.888888888888888,
			  "isOld" : false,
			  "positionType" : "ownedDeviceLocation",
			  "timeStamp" : 1698288601000,
			  "floorLevel" : -1,
			  "locationFinished" : true
			},
			"isAppleAudioAccessory" : true,
			"isFirmwareUpdateMandatory" : false,
			"crowdSourcedLocation" : {
			  "latitude" : 40.77777777777,
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
	  ],
	  "groupedItemIdentifiers": [
		[
		  "C8F2F66A-A2A4-4C6B-B61A-F89735BE0609",
		  "B0D3A871-BE54-4EE5-952B-4E5217A62B7F",
		  "E72B636C-945E-4E1E-AC38-1E9544D03B0E"
		]
	  ],
	  "itemIdentifiers": [
		"E72B636C-945E-4E1E-AC38-1E9544D03B0E",
		"C8F2F66A-A2A4-4C6B-B61A-F89735BE0609",
		"B0D3A871-BE54-4EE5-952B-4E5217A62B7F"
	  ],
	  "capabilities": 798,
	  "state": 65,
	  "groupedItems": [
		[
			{
				"productIdentifier" : "CA1E6D06-FB96-4FEA-84ED-A9B96C4C56E5",
				"groupIdentifier" : "4577AEC4-C5BF-44EF-BE0B-F5925604C796",
				"location" : {
				  "horizontalAccuracy" : 27.062871238639026,
				  "latitude" : 40.666666666666666,
				  "timeStamp" : 1698288601000,
				  "longitude" : -73.99999999999999,
				  "isOld" : false,
				  "locationFinished" : true,
				  "floorLevel" : -1,
				  "verticalAccuracy" : -1,
				  "isInaccurate" : false,
				  "altitude" : -1,
				  "positionType" : "ownedDeviceLocation"
				},
				"isAppleAudioAccessory" : true,
				"identifier" : "B0D3A871-BE54-4EE5-952B-4E5217A62B7F",
				"lostModeMetadata" : null,
				"serialNumber" : "GXXXXXXXXXXQ",
				"address" : {
				  "stateCode" : "NY",
				  "country" : "United States",
				  "administrativeArea" : "New York",
				  "locality" : "New York",
				  "streetAddress" : "123",
				  "streetName" : "Street St",
				  "label" : "123 Street St",
				  "countryCode" : "US",
				  "mapItemFullAddress" : "123 Street St, Brooklyn, NY  11111",
				  "areaOfInterest" : [
					"Long Island"
				  ],
				  "formattedAddressLines" : [
					"123 Street St",
					"Brooklyn, NY  11111",
					"United States"
				  ],
				  "fullThroroughfare" : "123 Street St",
				  "subAdministrativeArea" : "Kings County"
				},
				"safeLocations" : [
				  {
					"address" : {
					  "areaOfInterest" : [
						"Long Island"
					  ],
					  "formattedAddressLines" : [
						"234 Street St",
						"Brooklyn, NY  11111",
						"United States"
					  ],
					  "mapItemFullAddress" : "234 Street St, Brooklyn, NY  11111",
					  "subAdministrativeArea" : "Kings County",
					  "administrativeArea" : "New York",
					  "country" : "United States",
					  "stateCode" : "NY",
					  "streetAddress" : "234",
					  "locality" : "New York",
					  "fullThroroughfare" : "234 Street St",
					  "countryCode" : "US",
					  "streetName" : "Street St",
					  "label" : "234 Street St"
					},
					"location" : {
					  "floorLevel" : -1,
					  "timeStamp" : 1698289517185,
					  "isOld" : true,
					  "isInaccurate" : false,
					  "latitude" : 40.1111111111111,
					  "longitude" : -73.888888888888888,
					  "positionType" : "safeLocation",
					  "locationFinished" : true,
					  "verticalAccuracy" : 80,
					  "altitude" : 0,
					  "horizontalAccuracy" : 80
					},
					"identifier" : "ED8617E8-17E8-4519-9C91-4A2C420D17B5",
					"type" : 1,
					"name" : "Home",
					"approvalState" : 1
				  }
				],
				"crowdSourcedLocation" : {
				  "positionType" : "ownedDeviceLocation",
				  "altitude" : -1,
				  "latitude" : 40.77777777777,
				  "longitude" : -73.888888888888888,
				  "horizontalAccuracy" : 27.062871238639026,
				  "verticalAccuracy" : -1,
				  "isInaccurate" : false,
				  "isOld" : false,
				  "timeStamp" : 1698288601000,
				  "floorLevel" : -1,
				  "locationFinished" : true
				},
				"name" : "Left Bud",
				"role" : {
				  "emoji" : "😀",
				  "identifier" : 999,
				  "name" : "Custom Name"
				},
				"productType" : {
				  "productInformation" : {
					"defaultListIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist.png",
					"appBundleIdentifier" : "",
					"defaultListIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__2x.png",
					"productIdentifier" : 8212,
					"requiresAdditionalConnectionTime" : false,
					"defaultListIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__3x.png",
					"antennaPower" : 4,
					"vendorIdentifier" : 76,
					"defaultHeroIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__3x.png",
					"defaultHeroIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__2x.png",
					"defaultHeroIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox.png",
					"modelName" : "AirPods Pro (2nd generation)",
					"requiresAudioSafetyAlert" : true,
					"manufacturerName" : "Apple"
				  },
				  "type" : "hawkeye"
				},
				"systemVersion" : "80.23.144",
				"capabilities" : 814,
				"batteryStatus" : 0,
				"partInfo" : {
				  "type" : "leftBud",
				  "name" : "Left Bud",
				  "symbol" : "l.circle.fill"
				},
				"isFirmwareUpdateMandatory" : false,
				"owner" : "owner@localhost"
			},{
				"productType" : {
				  "type" : "hawkeye",
				  "productInformation" : {
					"manufacturerName" : "Apple",
					"defaultHeroIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox.png",
					"requiresAudioSafetyAlert" : true,
					"requiresAdditionalConnectionTime" : false,
					"appBundleIdentifier" : "",
					"productIdentifier" : 8212,
					"vendorIdentifier" : 76,
					"defaultHeroIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__3x.png",
					"defaultListIcon" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist.png",
					"defaultHeroIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-infobox__2x.png",
					"modelName" : "AirPods Pro (2nd generation)",
					"defaultListIcon2x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__2x.png",
					"defaultListIcon3x" : "https:\/\/statici.icloud.com\/fmipmobile\/deviceImages-9.0\/Accessory\/AirPods_8212-0\/online-sourcelist__3x.png",
					"antennaPower" : 4
				  }
				},
				"crowdSourcedLocation" : {
				  "isOld" : false,
				  "locationFinished" : true,
				  "positionType" : "ownedDeviceLocation",
				  "timeStamp" : 1698288601000,
				  "isInaccurate" : false,
				  "longitude" : -73.888888888888888,
				  "verticalAccuracy" : -1,
				  "altitude" : -1,
				  "latitude" : 40.777777777777,
				  "floorLevel" : -1,
				  "horizontalAccuracy" : 27.062871238639026
				},
				"role" : {
				  "emoji" : "😀",
				  "identifier" : 999,
				  "name" : "Custom Name"
				},
				"lostModeMetadata" : null,
				"batteryStatus" : 0,
				"owner" : "owner@localhost",
				"isAppleAudioAccessory" : true,
				"partInfo" : {
				  "type" : "rightBud",
				  "name" : "Right Bud",
				  "symbol" : "r.circle.fill"
				},
				"address" : {
				  "administrativeArea" : "New York",
				  "streetAddress" : "123",
				  "subAdministrativeArea" : "Kings County",
				  "stateCode" : "NY",
				  "country" : "United States",
				  "fullThroroughfare" : "123 Street St",
				  "label" : "123 Street St",
				  "countryCode" : "US",
				  "formattedAddressLines" : [
					"123 Street St",
					"Brooklyn, NY  11111",
					"United States"
				  ],
				  "mapItemFullAddress" : "123 Street St, Brooklyn, NY  11111",
				  "locality" : "New York",
				  "streetName" : "Street St",
				  "areaOfInterest" : [
					"Long Island"
				  ]
				},
				"groupIdentifier" : "4577AEC4-C5BF-44EF-BE0B-F5925604C796",
				"serialNumber" : "GWDJJJJJJJJP",
				"capabilities" : 814,
				"safeLocations" : [
				  {
					"name" : "Home",
					"address" : {
					  "mapItemFullAddress" : "123 Street St, Brooklyn, NY  11111",
					  "fullThroroughfare" : "123 Street St",
					  "countryCode" : "US",
					  "streetAddress" : "123",
					  "administrativeArea" : "New York",
					  "subAdministrativeArea" : "Kings County",
					  "streetName" : "Street St",
					  "formattedAddressLines" : [
						"123 Street St",
						"Brooklyn, NY  11111",
						"United States"
					  ],
					  "locality" : "New York",
					  "stateCode" : "NY",
					  "country" : "United States",
					  "label" : "123 Street St",
					  "areaOfInterest" : [
						"Long Island"
					  ]
					},
					"type" : 1,
					"approvalState" : 1,
					"location" : {
					  "floorLevel" : -1,
					  "timeStamp" : 1698289517185,
					  "isOld" : true,
					  "isInaccurate" : false,
					  "latitude" : 40.1111111111111,
					  "longitude" : -73.888888888888888,
					  "positionType" : "safeLocation",
					  "locationFinished" : true,
					  "verticalAccuracy" : 80,
					  "altitude" : 0,
					  "horizontalAccuracy" : 80
					},
					"identifier" : "ED8617E8-17E8-4519-9C91-4A2C420D17B5"
				  }
				],
				"systemVersion" : "80.23.144",
				"name" : "Right Bud",
				"productIdentifier" : "CA1E6D06-FB96-4FEA-84ED-A9B96C4C56E5",
				"identifier" : "C8F2F66A-A2A4-4C6B-B61A-F89735BE0609",
				"isFirmwareUpdateMandatory" : false,
				"location" : {
				  "isOld" : false,
				  "locationFinished" : true,
				  "verticalAccuracy" : -1,
				  "timeStamp" : 1698288601000,
				  "positionType" : "ownedDeviceLocation",
				  "latitude" : 40.666666666666666,
				  "isInaccurate" : false,
				  "altitude" : -1,
				  "floorLevel" : -1,
				  "longitude" : -73.888888888888888,
				  "horizontalAccuracy" : 27.062871238639026
				}
			},{
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
					  "longitude" : -73.888888888888888,
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
				  "longitude" : -73.888888888888888,
				  "isOld" : false,
				  "positionType" : "ownedDeviceLocation",
				  "timeStamp" : 1698288601000,
				  "floorLevel" : -1,
				  "locationFinished" : true
				},
				"isAppleAudioAccessory" : true,
				"isFirmwareUpdateMandatory" : false,
				"crowdSourcedLocation" : {
				  "latitude" : 40.77777777777,
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
		]
	  ]
	}
  ]`

func TestReadItemGroup_AirTags(t *testing.T) {
	bytes := []byte(AirTagsGroupInput)
	groups, err := readItemGroups(bytes)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(groups))

	g := groups[0]
	assert.Equal(t, "James’s AirPods Pro", g.Name)
	assert.Equal(t, "4577AEC4-C5BF-44EF-BE0B-F5925604C796", g.Identifier)
	assert.Equal(t, "p", g.PairingState("C8F2F66A-A2A4-4C6B-B61A-F89735BE0609"))
	assert.Equal(t, "p", g.PairingState("B0D3A871-BE54-4EE5-952B-4E5217A62B7F"))
	assert.Equal(t, "p", g.PairingState("E72B636C-945E-4E1E-AC38-1E9544D03B0E"))
	assert.Equal(t, [][]string{
		{
			"C8F2F66A-A2A4-4C6B-B61A-F89735BE0609",
			"B0D3A871-BE54-4EE5-952B-4E5217A62B7F",
			"E72B636C-945E-4E1E-AC38-1E9544D03B0E",
		},
	}, g.GroupedItemIdentifiers)
	assert.Equal(t, []string{
		"E72B636C-945E-4E1E-AC38-1E9544D03B0E",
		"C8F2F66A-A2A4-4C6B-B61A-F89735BE0609",
		"B0D3A871-BE54-4EE5-952B-4E5217A62B7F",
	}, g.ItemIdentifiers)

	items := groups[0].Items

	// left bud
	i := items[0]
	assert.Equal(t, "Left Bud", i.Name)
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
		Latitude:     40.77777777777,
		Longitude:    -73.888888888888888,
		Timestamp:    1698288601000,
		IsInaccurate: false,
		IsOld:        false,
	}, i.CrowdSourcedLocation)
	assert.Equal(t, "B0D3A871-BE54-4EE5-952B-4E5217A62B7F", i.Identifier)
	assert.Equal(t, "4577AEC4-C5BF-44EF-BE0B-F5925604C796", i.GroupIdentifier)
	assert.Equal(t, "CA1E6D06-FB96-4FEA-84ED-A9B96C4C56E5", i.ProductIdentifier)
	assert.Equal(t, "GXXXXXXXXXXQ", i.SerialNumber)
	assert.Equal(t, "80.23.144", i.SystemVersion)
	assert.Equal(t, FmItemPartInfo{
		Type: "leftBud",
		Name: "Left Bud",
	}, i.PartInfo)

	// right bud
	i = items[1]
	assert.Equal(t, "Right Bud", i.Name)
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
		Longitude:    -73.888888888888888,
		Timestamp:    1698288601000,
		IsInaccurate: false,
		IsOld:        false,
	}, i.Location)
	assert.Equal(t, FmItemLocation{
		Latitude:     40.777777777777,
		Longitude:    -73.888888888888888,
		Timestamp:    1698288601000,
		IsInaccurate: false,
		IsOld:        false,
	}, i.CrowdSourcedLocation)
	assert.Equal(t, "C8F2F66A-A2A4-4C6B-B61A-F89735BE0609", i.Identifier)
	assert.Equal(t, "4577AEC4-C5BF-44EF-BE0B-F5925604C796", i.GroupIdentifier)
	assert.Equal(t, "CA1E6D06-FB96-4FEA-84ED-A9B96C4C56E5", i.ProductIdentifier)
	assert.Equal(t, "GWDJJJJJJJJP", i.SerialNumber)
	assert.Equal(t, "80.23.144", i.SystemVersion)
	assert.Equal(t, FmItemPartInfo{
		Type: "rightBud",
		Name: "Right Bud",
	}, i.PartInfo)

	// case
	i = items[2]
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
		Longitude:    -73.888888888888888,
		Timestamp:    1698288601000,
		IsInaccurate: false,
		IsOld:        false,
	}, i.Location)
	assert.Equal(t, FmItemLocation{
		Latitude:     40.77777777777,
		Longitude:    -73.888888888888888,
		Timestamp:    1698288601000,
		IsInaccurate: false,
		IsOld:        false,
	}, i.CrowdSourcedLocation)
	assert.Equal(t, 1, len(i.SafeLocations))
	assert.Equal(t, FmItemSafeLocation{
		Location: FmItemLocation{
			Latitude:     40.1111111111111,
			Longitude:    -73.888888888888888,
			Timestamp:    1698289517185,
			IsInaccurate: false,
			IsOld:        true,
		},
		Address: FmItemAddress{
			Label:              "234 Street St",
			MapItemFullAddress: "234 Street St, Brooklyn, NY  11111",
			FormattedAddressLines: []string{
				"234 Street St",
				"Brooklyn, NY  11111",
				"United States",
			},
		},
		Name: "Home",
	}, i.SafeLocations[0])
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
