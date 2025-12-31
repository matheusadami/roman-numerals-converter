APPNAME=roman-numerals-converter
CMDPATH=./cmd/roman-numerals-converter
INSTALLPATH=github.com/matheusadami/roman-numerals-converter/cmd/roman-numerals-converter@v1.0.0
BINDIR=bin

GO=go

# Build the app
build:
	$(GO) build -o $(BINDIR)/$(APPNAME) $(CMDPATH)

# Run the app
# E.g.: make run ARGS="--number=100"
run:
	$(GO) run $(CMDPATH) $(ARGS)

# Install the app globally
install:
	$(GO) install $(INSTALLPATH)

# Run tests
test:
	$(GO) test ./...

# Format the code
fmt:
	$(GO) fmt ./...

# Clean the build artifacts
clean:
	rm -rf $(BINDIR)
