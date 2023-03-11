root = $(shell pwd)

export PROG			= pkflux
export BINROOT 		= $(root)/bin
export CMDROOT  	= $(root)/cmd
export TESTROOT 	= $(root)/tests

BINARY=$(BINROOT)/$(PROG)
BUILD_PKG=github.com/cciva/pkflux/build

LDFLAGS += -X $(BUILD_PKG).BuildVersion=v0.0.1
LDFLAGS += -X $(BUILD_PKG).BuildTime=$(shell date --iso=seconds)

.PHONY: all

all: dirs clean pkflux

dirs:
	@mkdir -p $(BINROOT)
	
pkflux:
	$(info [*] compiling '$(BINARY)'...)
	@go build -ldflags "$(LDFLAGS)" -o $(BINARY) $(CMDROOT)/$(PROG)/main.go
	
clean:
	$(info [*] purging'$(BINARY)'...)
	@rm -f $(BINARY)