all:
	rm -rf *.html; \
	cd $(CURDIR)/src; \
	go run main.go; \
	go fmt; \
	cd $(CURDIR); \
	chromium-browser index.html;
