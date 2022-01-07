EXTENSION=so
ifeq ($(OS), Windows_NT)
	EXTENSION=dll
endif

.PHONY: build
build:
ifeq ($(OS), Windows_NT)
	set GOPRIVATE=github.com/maiderdiscord
else
	export GOPRIVATE=github.com/maiderdiscord
endif
	go build -buildmode=c-shared -ldflags="-s -w" -trimpath -o ../app/golib.$(EXTENSION) ./...
