build:
	@go build

test:
	@go fmt
	@go test ./... -v

coverage:
	echo "Test Coverage script will be here"

clean:
	rm heimdall-plugin
