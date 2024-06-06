# Setup Variable
GO = go
COVERAGE_FILE = cover.out

# Setup Rules Default
all: test coverage

# Setup Rules to Run Testing
test:
	@echo "Running testing "
	@echo "This may take a moment..."
	@echo ""
	$(GO) test ./... -v -coverprofile=$(COVERAGE_FILE)
	@echo ""

# Setup Rules for Generate Coverage Reports in HTML Format
coverage: test
	@echo ""
	@echo "Generating coverage report in HTML format.."
	@echo ""
	$(GO) tool cover -html=$(COVERAGE_FILE)