# Auto Tank

3D printed, raspberry pi powered [ProtoTank](https://www.thingiverse.com/thing:972768).  The pi takes input from a camera, accelerometer, and magnetometer and drives an OLED screen and motor controller.

## Goals:
* Direct control via a web server hosted on the PI
* LOGO interpreter
* Autonomous driving

## Running:
```
go build
sudo ./auto-tank
```

WiringPi requires sudo access to be able to run PWM fast enough.

Visit http://<raspberry pi ip>:8080 in a web browser.

## Installation:
Running on a Raspberry Pi B+ running Raspbian stretch

Install Go 1.11:
```
https://dl.google.com/go/go1.11.1.linux-armv6l.tar.gz
sudo tar -C /usr/local -xvf go1.11.1.linux-arm61.tar.gz
```

Add the following lines to ~/.profile:
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Check with `go version`

Checkout 

## Useful commands
* Enable GPIO and I2C (`sudo raspi-config`)
* Raspbian 9 comes with [Wiring Pi](http://wiringpi.com/download-and-install/) installed - check with `gpio -v` and `gpio readall`
* List i2c devices with `i2cdetect -y 1`
* Probe i2c devices with `i2cdetect 1` - i2c busses are listed under `ls /dev | grep i2c`.  There's two, but only bus `1` is enabled by default.  Bus `0` requires pullup resistors, since it uses GPIO pins.
* The `pinout` command lists useful info.
* `gpio readall` prints the pin numbers for the GPIO header.