.PHONY: run
run:
	@go run .

.PHONY: e2e
e2e:
	@go clean -testcache
	@go test ./test/e2e/ -v