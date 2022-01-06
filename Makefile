.PHONY: build
build:
	cd dll && make build
	cd app && npm run build
