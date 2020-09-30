functions := $(shell find functions -name \*handler.go | awk -F'/' '{print $$2}')
build: 
	@for function in $(functions) ; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$function functions/$$function/handler.go ; \
	done

clean: ## Clean build
	rm -rf ./bin