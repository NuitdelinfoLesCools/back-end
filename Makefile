GOOS=linux
GO111MODULE=on
OS_TYPE=$(shell uname -a)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITCOMMIT=$(shell git rev-parse HEAD)
BUILDDATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
VERSION=1.0.0

compile:
	GOOS=linux go build \
		-o server.o \
		cmd/*.go

compile-windows:
    # Well after this your on your own ;-)
	GOOS=windows go build \
		-o server.exe
		cmd/*.go

run-dev:
	go run cmd/*.go -config dev.config.yaml

go-mod:
	GO111MODULE=on go mod init && \
		go mod tidy

build-image:
	docker build -t nilescools/back-end:latest .


docker-login:
	docker login \
		-u $(DOCKER_USER) \
		-p $(DOCKER_PASS)

push-image:
	docker push nilescools/back-end


