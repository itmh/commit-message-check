.PHONY: build-all upload-release

build-all:
	@./scripts/build-binaries.sh

upload-release:
	@./scripts/upload-github-release.sh
