PROJS = scribe scribecmd evrtest ubuntu-cve-tracker
GO = GOPATH=$(shell pwd):$(shell go env GOROOT)/bin go
export SCRIBECMD = $(shell pwd)/bin/scribecmd
export EVRTESTCMD = $(shell pwd)/bin/evrtest

all: $(PROJS)

ubuntu-cve-tracker:
	$(GO) install ubuntu-cve-tracker

evrtest:
	$(GO) install evrtest

scribe:
	$(GO) build scribe
	$(GO) install scribe

scribecmd:
	$(GO) install scribecmd

runtests: scribetests gotests

gotests:
	$(GO) test -v scribe

scribetests: $(PROJS)
	cd test && $(MAKE) runtests

clean:
	rm -rf pkg
	rm -f bin/*
	cd test && $(MAKE) clean
