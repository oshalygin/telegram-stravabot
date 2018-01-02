pwd 			=	 $(shell pwd)
base_path		=	"https://github.com/golang/dep/releases/download"
utility_version	=	"v0.3.2"
utility_version	=	"v0.3.2"
utility_binary	=	"dep-linux-amd64"
dep				=	"$(pwd)/.travis.d/dep"

dep-install:
	- wget $(base_path)/$(utility_version)/$(utility_binary) -O $(dep)
	- chmod +x $(dep)
	- $(dep) ensure

lint:
	go vet ./...

.PHONY: lint dep-install
