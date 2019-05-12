.PHONY: clean build build-mock mock tank run

default: clean build

clean:
	rm -f auto-tank

stop:
	./stop.sh

build:
	go build

build-mock:
	go build -tags mock

mock: clean build-mock

tank:
	sudo ./auto-tank

run: clean build tank
