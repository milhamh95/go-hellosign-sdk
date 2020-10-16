TEST_OPTS := -covermode=atomic $(TEST_OPTS)

# Install dependency
.PHONY: vendor
vendor: go.mod go.sum
	@GO111MODULE=on go get ./...

# Testing
.PHONY: unittest
unittest:
	go test -short -race $(TEST_OPTS) ./...
