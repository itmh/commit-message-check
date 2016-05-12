GOTOOLS = \
	github.com/aktau/github-release

.PHONY: build-all upload-release tools test

build-all:
	@./scripts/build-binaries.sh

upload-release:
	@./scripts/upload-github-release.sh

tools:
	go get -u -v $(GOTOOLS)

test:
	@go test --cover
