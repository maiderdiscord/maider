EXTENSION=so
ifeq ($(OS), Windows_NT)
	EXTENSION=dll
endif

.PHONY: build
build:
	go build -buildmode=c-shared -o ../app/golib.$(EXTENSION) ./...
