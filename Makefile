BROKER_BINARY=brokerApp

## build_broker: builds the broker binary as a linux executable
build:
	@echo "Building broker binary..."
	mkdir -p .build
	cd ./.build && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ../cmd/broker
	@echo "Done!"

clean:
	@echo "Cleaning broker binary..."
	rm -rf ./.build
	@echo "Done!"