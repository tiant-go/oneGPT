VERSION=v3.0.0

build:
	go mod tidy &&  CGO_ENABLED=0  \
    go build --tags=kqueue,operator -ldflags "-w -s" -o bin/oneGPT
