[![Build Status](https://travis-ci.org/jhitze/oak_weather_recorder.svg?branch=master)](https://travis-ci.org/jhitze/oak_weather_recorder)

# Oak Weather Recorder
The neatest way to record data published over from the Oak Shield over the Particle network.

## Inspiration
https://github.com/who93/oak_weathershield

## How to build

```
go get github.com/jhitze/oak-weather
go get github.com/peterhellberg/sseclient
cd oak-weather
go build *.go 
```

## How to run
```
./OakWeatherRecorder
```

## How to test
```
go test
```

## In action

First time:

```
./OakWeatherRecorder 
2016/03/31 22:32:59.685253 Going to attempt to load data from oak_weather.json
2016/03/31 22:32:59.685519 Could not read settings file. reason: open oak_weather.json: no such file or directory
2016/03/31 22:32:59.685532 Reverting to asking for the settings.
Enter username: user@example.com
Enter password: *******************
Found Devices
----------------------------------------
0: oak-1 - Online->false
1: oak-weather - Online->true
----------------------------------------
Pick a number:
1
2016/03/31 22:33:15.804334 Device oak-weather picked.
2016/03/31 22:33:15.804353 Going to attempt to save settings to oak_weather.json
2016/03/31 22:33:15.804701 Saved successfully for next time!
2016/03/31 22:33:16.016650 Connected to the stream of device oak-weather (--sanitized, device id was here--)
2016/03/31 22:33:20.936188 Temperature: 20.57°C, Ambient: 7.03%, Pressure: 1001.15 mbar, Humidity: 54.84%
```

Second time (or any time there is a proper JSON settings file)

```
./OakWeatherRecorder 
2016/03/31 22:33:27.777095 Going to attempt to load data from oak_weather.json
2016/03/31 22:33:27.777414 Getting current information for device: --sanitized, device id was here--
2016/03/31 22:33:28.377079 Connected to the stream of device oak-weather (--sanitized, device id was here--)
2016/03/31 22:33:31.376787 Temperature: 20.51°C, Ambient: 5.08%, Pressure: 1001.21 mbar, Humidity: 54.97%
```
