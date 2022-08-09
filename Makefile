# raccoon
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: raccoon

raccoon:
	$(GO) build $(GOFLAGS)

clean:
	$(RM) raccoon

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f raccoon $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/raccoon

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/raccoon

.DEFAULT_GOAL := all

.PHONY: all raccoon clean install uninstall
