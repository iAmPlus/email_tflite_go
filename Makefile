.PHONY: build
PROJECT_PATH := $(CURDIR)

build:
	CGO_CFLAGS=-I$(PROJECT_PATH) CGO_LDFLAGS=-L"$(PROJECT_PATH)/libs" CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./application . 