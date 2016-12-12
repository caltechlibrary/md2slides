#
# Makefile to compile mkslides for Mac OS X, Linux, Windows 7
# as well as Raspberry Pi Zero, 1,2, and 3.
#

build:
	go build -o bin/mkslides cmds/mkslides/mkslides.go
	./mk-website.bash

test:
	go test

install:
	env GOBIN=$(HOME)/bin go install cmds/mkslides/mkslides.go

status:
	git status

save:
	./mk-website.bash
	git commit -am "Quick save"
	git push origin master

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f mkslides-release.zip ]; then rm mkslides-release.zip; fi

release:
	./mk-website.bash
	./mk-release.bash	
	
publish:
	./publish.bash
