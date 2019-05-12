default: clean build

clean:
	rm -f auto-tank

stop:
	./stop.sh

build:
	go build

build-mock:
    go build -tags mock

tank:
	sudo ./auto-tank

run: clean build tank
