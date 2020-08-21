# start running server
.PHONY: run
run:
	@go run .

# e2e test
.PHONY: e2e
e2e:
	@go clean -testcache
	@go test -v ./test/e2e/ 

# unit test
.PHONY: unit
unit:
	@go clean -testcache
	@go test -v ./msg/processor
	@go test ./msg/processor/ -cover
