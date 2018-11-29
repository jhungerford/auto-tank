default: clean build

clean:
	rm -f auto-tank

stop:
	./stop.sh

build:
	go build

tank:
	sudo ./auto-tank

run: clean build tank

