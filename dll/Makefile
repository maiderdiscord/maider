EXTENSION=so
ifeq ($(OS), Windows_NT)
	EXTENSION=dll
endif

.PHONY: build
build:
	GOPRIVATE=github.com/maiderdiscord go build -buildmode=c-shared -ldflags="-s -w" -trimpath -o ../app/golib.$(EXTENSION) ./...
