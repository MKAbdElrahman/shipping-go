GO_VERSION := 1.21.0

# TODO: Add support for  MacOS
install-go:
	wget https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

setup: install-go init-go

test:
	go test  ./... -coverprofile=coverage.out
coverage: test
	go tool cover -func=coverage.out | grep "total:" |  awk '{print ((int($$3) > 80) != 1) }'
report: test
	go tool cover -html=coverage.out -o coverage.html
build:
	go build -o  api cmd/main.go
clean:
	rm -rf api

