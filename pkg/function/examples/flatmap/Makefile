.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./dist/flatmap-example main.go

.PHONY: image
image: build
	docker build -t "quay.io/numaio/flatmap-example:go" --target flatmap .

clean:
	-rm -rf ./dist
