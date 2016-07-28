#
# Makefile to compile md2slides for Mac OS X, Linux, Windows 7
# as well as R-pi.
#

build:
	go build -o bin/md2slides cmds/md2slides/md2slides.go
	./mk-website.bash

install:
	env GOBIN=$(HOME)/bin go install cmds/md2slides/md2slides.go

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f md2slides-binary-release.zip ]; then rm md2slides-binary-release.zip; fi

release:
	./mk-website.bash
	./mk-release.bash	
	
save:
	git commit -am "Quick save"
	git push origin master

publish:
	./publish.bash
