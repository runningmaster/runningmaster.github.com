all:
	rm -rf *.html; \
	cd $(CURDIR)/src; \
	go run main.go; \
	go fmt;
