GOBUILD = go build
GOTEST = go test
GOGET = go get
GOCLEAN = go clean

# Binary name
BINARY_NAME = authx

# Main build target
all: deps test build

# Install dependencies
deps:
	$(GOGET) -v ./...

# Run tests
test:
	$(GOTEST) -v ./...

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)