NAME := kalvados
VERSION = $(shell gobump show -r ./version)
REVISION := $(shell git rev-parse --short HEAD)

all: build

setup:
	go get golang.org/x/vgo
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/tcnksm/ghr
	go get github.com/Songmu/goxz/cmd/goxz
	go get github.com/motemen/gobump/cmd/gobump

test: lint
	vgo test ./...
	vgo test -race ./...

lint: setup
	vgo vet ./...
	golint ./...

fmt: setup
	goimports -w .

build:
	vgo build -o bin/$(NAME) cmd/$(NAME)/main.go
	vgo build -o bin/$(NAME)-server cmd/$(NAME)-server/main.go

clean:
	rm bin/$(NAME)

package: setup
	@sh -c "'$(CURDIR)/scripts/package.sh'"

crossbuild: setup
	goxz -pv=v${VERSION} -build-ldflags="-X main.GitCommit=${REVISION}" \
        -arch=386,amd64 -d=./pkg/dist/v${VERSION} \
        -n ${NAME} ./cmd/${NAME}
	goxz -pv=v${VERSION} -build-ldflags="-X main.GitCommit=${REVISION}" \
        -arch=386,amd64 -d=./pkg/dist/v${VERSION} \
        -n ${NAME}-server ./cmd/${NAME}-server

release: package
	ghr -u aktsk v${VERSION} ./pkg/dist/v${VERSION}

bump:
	@sh -c "'$(CURDIR)/scripts/bump.sh'"
