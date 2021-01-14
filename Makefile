.PHONY: build release

COMMIT := $(shell git rev-parse --short HEAD)

build: 
	CGO_ENABLED=0 go build \
	-ldflags "-w -s -extldflags "-static"" \
	-trimpath \
	-o vc cmd/main.go

release:
	mkdir -p release
	GOOS=windows GOARCH=amd64 make build
	mv vc release/vc-win-amd64-$(COMMIT).exe
	GOOS=darwin GOARCH=amd64 make build
	mv vc release/vc-darwin-amd64-$(COMMIT)
	GOOS=linux GOARCH=amd64 make build
	mv vc release/vc-linux-amd64-$(COMMIT)
	GOOS=linux GOARCH=arm64 make build
	mv vc release/vc-linux-arm64-$(COMMIT)
