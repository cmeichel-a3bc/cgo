.PHONY: sharedTest staticTest inlineTest stringTest

all: sharedTest staticTest inlineTest stringTest

inlineTest:
	go build -o $@ ./cmd/inline/...

sharedTest: shared_lib
	go build -o $@ ./cmd/shared/...

staticTest: static_lib
	go build -o $@ ./cmd/static/... \

stringTest: static_lib
	go build -o $@ ./cmd/string/... \

shared_lib: shared
	$(MAKE) -C src shared

static_lib: static
	$(MAKE) -C src static

shared:
	mkdir -p $@

static:
	mkdir -p $@

clean:
	rm -rf shared static inlineTest sharedTest staticTest stringTest
	$(MAKE) -C src clean