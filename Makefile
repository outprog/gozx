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
	cp -r ./dependence/hidapi ./vendor/github.com/karalabe/hid/
	cp -r ./dependence/libusb ./vendor/github.com/karalabe/hid/
	cp -r ./dependence/crypto ./vendor/github.com/ethereum/go-ethereum/