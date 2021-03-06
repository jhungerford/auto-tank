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

Checkout this repo:
```
mkdir -p ~/go/src/github.com/jhungerford
cd ~/go/src/github.com/jhungerford
```

Make targets:

| Target     | Description                                                |
|------------|------------------------------------------------------------|
| make clean | Removes binaries that `make build` produces.               |
| make stop  | Stops the motors.                                          |
| make build | Compiles the binary.                                       |
| make tank  | Runs the tank with the latest binary - doesn't rebuild it. |
| make run   | Creates a clean build of the binary and runs it.           |


## Useful commands
* Enable GPIO and I2C (`sudo raspi-config`)
* Raspbian 9 comes with [Wiring Pi](http://wiringpi.com/download-and-install/) installed - check with `gpio -v` and `gpio readall`
* List i2c devices with `i2cdetect -y 1`
* Probe i2c devices with `i2cdetect 1` - i2c busses are listed under `ls /dev | grep i2c`.  There's two, but only bus `1` is enabled by default.  Bus `0` requires pullup resistors, since it uses GPIO pins.
* The `pinout` command lists useful info.
* `gpio readall` prints the pin numbers for the GPIO header.
* The pin numbering scheme depends on how wiring pi is initialized - currently using the wPi numbers.
* RaspberryPi hard locks when controlling a motor from pins 2, 3, and 24 with a fully charged battery.

```
$ gpio readall
 +-----+-----+---------+------+---+---Pi 3+--+---+------+---------+-----+-----+
 | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
 +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
 |     |     |    3.3v |      |   |  1 || 2  |   |      | 5v      |     |     |
 |   2 |   8 |   SDA.1 | ALT0 | 1 |  3 || 4  |   |      | 5v      |     |     |
 |   3 |   9 |   SCL.1 | ALT0 | 1 |  5 || 6  |   |      | 0v      |     |     |
 |   4 |   7 | GPIO. 7 |  OUT | 1 |  7 || 8  | 0 | IN   | TxD     | 15  | 14  |
 |     |     |      0v |      |   |  9 || 10 | 1 | IN   | RxD     | 16  | 15  |
 |  17 |   0 | GPIO. 0 |   IN | 0 | 11 || 12 | 0 | OUT  | GPIO. 1 | 1   | 18  |
 |  27 |   2 | GPIO. 2 |  OUT | 1 | 13 || 14 |   |      | 0v      |     |     |
 |  22 |   3 | GPIO. 3 |  OUT | 0 | 15 || 16 | 0 | OUT  | GPIO. 4 | 4   | 23  |
 |     |     |    3.3v |      |   | 17 || 18 | 1 | OUT  | GPIO. 5 | 5   | 24  |
 |  10 |  12 |    MOSI |   IN | 0 | 19 || 20 |   |      | 0v      |     |     |
 |   9 |  13 |    MISO |   IN | 0 | 21 || 22 | 0 | IN   | GPIO. 6 | 6   | 25  |
 |  11 |  14 |    SCLK |   IN | 0 | 23 || 24 | 0 | OUT  | CE0     | 10  | 8   |
 |     |     |      0v |      |   | 25 || 26 | 1 | IN   | CE1     | 11  | 7   |
 |   0 |  30 |   SDA.0 |   IN | 1 | 27 || 28 | 1 | IN   | SCL.0   | 31  | 1   |
 |   5 |  21 | GPIO.21 |   IN | 1 | 29 || 30 |   |      | 0v      |     |     |
 |   6 |  22 | GPIO.22 |   IN | 1 | 31 || 32 | 0 | IN   | GPIO.26 | 26  | 12  |
 |  13 |  23 | GPIO.23 |   IN | 0 | 33 || 34 |   |      | 0v      |     |     |
 |  19 |  24 | GPIO.24 |  OUT | 0 | 35 || 36 | 0 | OUT  | GPIO.27 | 27  | 16  |
 |  26 |  25 | GPIO.25 |   IN | 0 | 37 || 38 | 0 | IN   | GPIO.28 | 28  | 20  |
 |     |     |      0v |      |   | 39 || 40 | 0 | IN   | GPIO.29 | 29  | 21  |
 +-----+-----+---------+------+---+----++----+---+------+---------+-----+-----+
 | BCM | wPi |   Name  | Mode | V | Physical | V | Mode | Name    | wPi | BCM |
 +-----+-----+---------+------+---+---Pi 3+--+---+------+---------+-----+-----+
```