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
gcloud_deb_repo		= 	"http://packages.cloud.google.com/apt"
gcloud_gpg			=	"https://packages.cloud.google.com/apt/doc/apt-key.gpg"
gcloud_src_list		=	"/etc/apt/sources.list.d/google-cloud-sdk.list"
gcloud_sdk			=	"https://sdk.cloud.google.com"

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

remove_current_gcloud_dir:
	- rm -rf ${HOME}/google-cloud-sdk

install_gcloud_sdk: remove_current_gcloud_dir
	- curl $(gcloud_sdk) | bash

.PHONY: lint go-lint dep-install build install_gcloud_sdk remove_current_gcloud_dir
