current_time = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags := -s -X main.buildTime=$(current_time) -X main.version=$(git_description)

.PHONY: build/api

build/api:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="$(linker_flags)" -o app .