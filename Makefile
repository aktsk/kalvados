NAME := kalvados

all: build

setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports

deps: setup
	glide install

test: deps lint
	go test $$(glide novendor | grep -v cmd)
	go test -race $$(glide novendor | grep -v cmd)

lint: setup
	go vet $$(glide novendor | grep -v cmd)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

fmt: setup
	goimports -w $$(glide nv -x)

build: deps
	go build -o bin/$(NAME) cmd/kalvados/kalvados.go
	go build -o bin/$(NAME)-server cmd/kalvados/kalvados_server.go

clean:
	rm bin/$(NAME)
