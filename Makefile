NAME=ci-runner-helper
BINDIR=bin
GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info

all: linux macos
deploy: clean linux docker scp


linux:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-linux

macos:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-macos

docker:
	cp $(BINDIR)/$(NAME)-linux dockerfiles
	cd dockerfiles && docker build . -t ci-runner-helper:latest && docker save ci-runner-helper:latest -o ci-runner-helper.tar

scp:
	scp dockerfiles/ci-runner-helper.tar root@172.20.10.4:/tmp

releases: linux macos
	chmod +x $(BINDIR)/$(NAME)-*
	gzip $(BINDIR)/$(NAME)-linux
	gzip $(BINDIR)/$(NAME)-macos

clean:
	rm $(BINDIR)/*
	rm dockerfiles/$(NAME)-linux
	rm dockerfiles/$(NAME).tar