pwd 			=	 $(shell pwd)
base_path		=	"https://github.com/golang/dep/releases/download"
utility_version	=	"v0.3.2"
utility_version	=	"v0.3.2"
utility_binary	=	"dep-linux-amd64"
dep				=	"$(pwd)/dep"

dep-install:
	- wget $(base_path)/$(utility_version)/$(utility_binary) -O $(dep)
	- chmod +x $(dep)
	- $(dep) ensure

go-lint:
	- go get -u github.com/golang/lint/golint
	- golint -set_exit_status .

lint: go-lint
	go vet ./...

build:
	- CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

.PHONY: lint go-lint dep-install
