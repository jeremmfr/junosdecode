default: gotest/html

.PHONY: gotest gotest/html pkgsite pkgsite/install 

gotest:
	go test -race -covermode=count -v ./...

gotest/html:
	go test -race -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

pkgsite: pkgsite/install
	@echo "open -> http://localhost:8080/$(shell go list -m)"
	pkgsite

pkgsite/install:
	go install golang.org/x/pkgsite/cmd/pkgsite@latest
