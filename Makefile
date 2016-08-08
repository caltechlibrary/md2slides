#
# Makefile to compile md2slides for Mac OS X, Linux, Windows 7
# as well as Raspberry Pi Zero, 1,2, and 3.
#

build:
	go build -o bin/md2slides cmds/md2slides/md2slides.go
	./mk-website.bash

test:
	go test

install:
	env GOBIN=$(HOME)/bin go install cmds/md2slides/md2slides.go

status:
	git status

save:
	./mk-website.bash
	git commit -am "Quick save"
	git push origin master

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f md2slides-binary-release.zip ]; then rm md2slides-binary-release.zip; fi

release:
	./mk-website.bash
	./mk-release.bash	
	
publish:
	./publish.bash
