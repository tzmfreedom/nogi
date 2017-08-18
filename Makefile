NAME := nogi
SRCS := $(shell find . -type d -name vendor -prune -o -type f -name "*.go" -print)
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\"" 
DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := bin/$(NAME) 

.PHONY: test
test: glide
	@go test -cover -v `glide novendor`

.PHONY: install
install: 
	@go install .

.PHONY: uninstall
uninstall:

.PHONY: clean
clean:
	@rm -rf bin/*
	@rm -rf vendor/*
	@rm -rf dist/*

.PHONY: dist-clean
dist-clean: clean
	@rm -f $(NAME).tar.gz

.PHONY: cross-build
cross-build: deps
	@for os in darwin linux windows; do \
	    for arch in amd64 386; do \
	        GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo \
	        -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME)-$$os-$$arch; \
	    done; \
	done

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: hash
hash: 
	openssl dgst -sha256 dist/$(NAME)-$(VERSION)-darwin-386.tar.gz
	openssl dgst -sha256 dist/$(NAME)-$(VERSION)-darwin-amd64.tar.gz

.PHONY: deps
deps: glide
	glide install

.PHONY: build
build: dev/deps
	go-bindata data/
	make bin/$(NAME)

.PHONY: dev/deps
dev/deps:
ifeq ($(shell command -v go-bindata 2> /dev/null),)
	go get -u github.com/jteeuwen/go-bindata/...
endif

.PHONY: bin/$(NAME) 
bin/$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME) .

.PHONY: dist
dist:
	@cd dist && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) cp ../completions/zsh/_nogi {} \; && \
		$(DIST_DIRS) tar zcf $(NAME)-$(VERSION)-{}.tar.gz {} \;

