## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool celeritas and copies it to templates
build_cli:
	@go build -o ../myapp/celeritas ./cmd/cli

## build: builds the command line tool into the dist directory
build:
	@go build -o ./dist/celeritas ./cmd/cli

## install: builds and installs the command line tool
install: build
	@cp ./dist/celeritas ~/go/bin/celeritas
	@echo "Installed celeritas in ~/go/bin/celeritas"