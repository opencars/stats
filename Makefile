.PHONY: default all clean
APPS     := statsd
BLDDIR   := bin
VERSION  := $(shell cat VERSION)

.EXPORT_ALL_VARIABLES:
GO111MODULE  = on

default: clean all

all: $(APPS)

$(BLDDIR)/%:
	go build -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done
