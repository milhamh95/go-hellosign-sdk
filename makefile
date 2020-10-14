TEST_OPTS := -covermode=atomic $(TEST_OPTS)

# Testing
.PHONY: unittest
unittest:
	go test -short -race $(TEST_OPTS) ./...
