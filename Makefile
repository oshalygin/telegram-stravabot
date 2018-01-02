pwd 				=	 $(shell pwd)
base_path			=	"https://github.com/golang/dep/releases/download"
utility_version		=	"v0.3.2"
utility_version		=	"v0.3.2"
utility_binary		=	"dep-linux-amd64"
dep					=	"$(pwd)/dep"
gcloud_zone			=	"us-central1-a"
gcloud_project		=	"olegie-io"
gcloud_sdk_version	=	"183.0.0-0"
cloud_sdk_repo		= 	"cloud-sdk-$(lsb_release -c -s)"

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

install_gcloud_sdk:
	- rm -rf ${HOME}/google-cloud-sdk
  	- echo "deb http://packages.cloud.google.com/apt $(cloud_sdk_repo) main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
	- curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
	- sudo apt-get update && sudo apt-get install google-cloud-sdk=$(gcloud_sdk_version)

.PHONY: lint go-lint dep-install build install_gcloud_sdk
