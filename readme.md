# Oak Weather Recorder
The neatest way to record data published over from the Oak Shield over the Particle network.

## How to build

```
go get github.com/jhitze/oak-weather
cd oak-weather
go build *.go 
```

## How to run
```
./OakWeatherRecorder
```

## In action
```
./OakWeatherRecorder 
Enter username: user@example.com
Enter password: *******************
Found Devices
----------------------------------------
0: oak-1 - Online->false
1: oak-weather - Online->true
----------------------------------------
Pick a number:
1
2016/03/30 23:15:44.027142 Device oak-weather picked.
2016/03/30 23:15:44.103949 Connected to the stream of device oak-weather (--sanitized, device id was here--)
2016/03/30 23:15:48.279596 Temperature: 19.79°C, Ambient: 11.33%, Pressure: 1005.64 mbar, Humidity: 54.10%
2016/03/30 23:15:58.673257 Temperature: 19.80°C, Ambient: 11.33%, Pressure: 1005.61 mbar, Humidity: 54.06%
```
