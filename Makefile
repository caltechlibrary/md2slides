#
# Makefile to compile mkslides for Mac OS X, Linux, Windows 7
# as well as Raspberry Pi Zero, 1,2, and 3.
#

build:
	env CGO_ENABLED=0 go build -o bin/mkslides cmds/mkslides/mkslides.go

lint:
	golint mkslides.go
	golint cmds/mkslides/mkslides.go

test:
	go test

install:
	env GOBIN=$(HOME)/bin go install cmds/mkslides/mkslides.go

status:
	git status

save:
	git commit -am "Quick save"
	git push origin master

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f mkslides-release.zip ]; then rm mkslides-release.zip; fi

release:
	./mk-website.bash
	./mk-release.bash	
	
website:
	./mk-website.bash

publish:
	./publish.bash
