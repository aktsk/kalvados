setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports

deps: setup
	cd ../.. && glide install

deploy: deps
	GOPATH=../_gopath:../_gopath/vendor gcloud app deploy
