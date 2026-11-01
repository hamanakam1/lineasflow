PI_CC ?= aarch64-linux-gnu-gcc

build-mac:
	GOARCH=arm64 GOOS=darwin go build -o bin/lineasflowd-mac ./cmd/lineasflowd

build-pi:
	CGO_ENABLED=1 CC=$(PI_CC) GOARCH=arm64 GOOS=linux go build -o bin/lineasflowd-pi ./cmd/lineasflowd
