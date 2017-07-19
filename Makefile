PROJS = scribe scribecmd evrtest scribevulnpolicy
GO = GO15VENDOREXPERIMENT=1 go
GOLINT = golint

all: $(PROJS) runtests

evrtest:
	$(GO) install github.com/mozilla/scribe/evrtest

scribe:
	$(GO) install github.com/mozilla/scribe

scribecmd:
	$(GO) install github.com/mozilla/scribe/scribecmd

scribevulnpolicy:
	$(GO) install github.com/mozilla/scribe/scribevulnpolicy

runtests: scribetests gotests

gotests:
	$(GO) test -v -covermode=count -coverprofile=coverage.out github.com/mozilla/scribe

showcoverage: gotests
	$(GO) tool cover -html=coverage.out

scribetests: $(PROJS)
	cd test && SCRIBECMD=$$(which scribecmd) EVRTESTCMD=$$(which evrtest) $(MAKE) runtests

lint:
	$(GOLINT) $(PROJECT)

vet:
	$(GO) vet $(PROJECT)

go_vendor_dependencies:
	govend -u
	rm -rf vendor/github.com/mozilla/scribe
	[ $$(ls -A vendor/github.com/mozilla) ] || rm -r vendor/github.com/mozilla
	[ $$(ls -A vendor/github.com) ] || rm -r vendor/github.com

clean:
	rm -rf pkg
	rm -f bin/*
	cd test && $(MAKE) clean

.PHONY: $(PROJS) runtests gotests showcoverage scribetests lint vet clean
