# macos-findmy-mqtt

Writes information on AirTags and other Find My devices, obtained from reading a MacOS device's local Find My cache, to a MQTT broker for parsing as Home Assistant device tracker entities.

## Setup

This is a standard go application which can be built via `go build`.

```bash
go build .
chmod +x ./macos-findmy-mqtt
```

## Run

Provide the address to a MQTT broker, with optional username+password. `freq` configures the polling interval at which new data will be read and sent. To run once as part of a cronjob, set `-freq=0`.

```bash
./macos-findmy-mqtt \
  -broker='tcp://homeassistant:1883' \
  -user='username' \
  -password='password' \
  -freq=60
```