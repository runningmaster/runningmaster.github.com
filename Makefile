all: build

build:
	rm -rf *.html; \
	cd $(CURDIR)/src; \
	go run main.go; \
	go fmt; \
	cd $(CURDIR); \
	chromium-browser index.html;

install:
	go get -u github.com/russross/blackfriday