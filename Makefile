GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

PACKAGE := github.com/outprog/gozx

test:
	sh ./test.sh

dep:
	dep ensure -v
	cp -r ./dependency/hidapi ./vendor/github.com/karalabe/hid/
	cp -r ./dependency/libusb ./vendor/github.com/karalabe/hid/
	cp -r ./dependency/crypto ./vendor/github.com/ethereum/go-ethereum/