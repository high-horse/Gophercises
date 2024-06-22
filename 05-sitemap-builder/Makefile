# Makefile for Sitemap Builder

.DEFAULT_GOAL := help
.PHONY: help run

help:
	@echo "Usage:"
	@echo "  make run [URL=] [DEPTH=]    Run the sitemap builder"
	@echo ""
	@echo "Options:"
	@echo "  URL=yoururl.com         Specify the root URL for sitemap generation"
	@echo "  DEPTH=2                 Specify the maximum depth of links to follow (default: 2)"
	@echo ""

run:
	@if [ -n "$(URL)" ] && [ -n "$(DEPTH)" ]; then \
		go run cmd/main.go --url="$(URL)" --depth="$(DEPTH)"; \
	elif [ -n "$(URL)" ]; then \
		go run cmd/main.go --url="$(URL)"; \
	elif [ -n "$(DEPTH)" ]; then \
		go run cmd/main.go --depth="$(DEPTH)"; \
	else \
		go run cmd/main.go; \
	fi
