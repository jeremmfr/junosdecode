default: gotest/html

.PHONY: gotest gotest/html pkgsite pkgsite/install 

gotest:
	go test -race -covermode=count -v ./...

gotest/html:
	go test -race -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

pkgsite:
	go doc -http
