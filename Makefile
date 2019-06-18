.PHONY: build

build:
	go build -o frontdesk ./cmd/frontdesk

test:
	go test $(FLAGS) ./...

clean:
	rm -f frontdesk