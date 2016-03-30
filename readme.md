# Oak Weather Recorder
The neatest way to record data published over from the Oak Shield over the Particle network.

## How to build

```
go get github.com/jhitze/oak-weather
cd oak-weather
go build OakWeatherRecorder.go 
```

## How to run
```
./OakWeatherRecorder {device id} {access id}
```

## In action
```
./OakWeatherRecorder dOakDeviceNumbersAndLetters aAccessCodeForParticle
2016/03/29 22:51:32.817848 Connected to device dOakDeviceNumbersAndLetters
2016/03/29 22:51:43.363526 Temperature: 19.75°C, Ambient: 0.39%, Pressure: 1021.36mbar, Humidity: 52.11%
2016/03/29 22:51:54.875323 Temperature: 19.74°C, Ambient: 0.39%, Pressure: 1021.35mbar, Humidity: 52.08%
```