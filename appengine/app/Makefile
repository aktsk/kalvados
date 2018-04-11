setup:
	go get golang.org/x/vgo
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports

deps: setup
	cd ../.. && vgo vendor

deploy: deps
	GOPATH=../_gopath:../_gopath/vendor gcloud app deploy
