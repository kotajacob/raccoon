# raccoon
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: raccoon

raccoon:
	$(GO) build $(GOFLAGS)
	scdoc < raccoon.1.scd > raccoon.1

clean:
	$(RM) raccoon
	$(RM) raccoon.1

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f raccoon $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/raccoon
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f raccoon.1 $(DESTDIR)$(MANPREFIX)/man1/raccoon.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/raccoon.1

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/raccoon
	$(RM) $(DESTDIR)$(MANPREFIX)/man1/raccoon.1

.DEFAULT_GOAL := all

.PHONY: all raccoon clean install uninstall
