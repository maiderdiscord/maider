EXTENSION=so
ifeq ($(OS), WINDOWS_NT)
	EXTENSION=dll
endif

.PHONY: build
build:
	go build -buildmode=c-shared -ldflags="-s -w" -trimpath -o ../app/golib.$(EXTENSION) ./...
