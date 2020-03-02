OUTPUT := harbor-exporter

build: build-mac

build-mac:
	@echo "Building for mac .."
	GOOS=darwin go build -o $(OUTPUT) -ldflags "-X main.appVersion=$(shell cat VERSION)"
	@echo "Done"
