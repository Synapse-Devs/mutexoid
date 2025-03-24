# Mutexoid

<p align="center">
  <img src="/assets/logo.svg" width="200" height="200" alt="Mutexoid Logo">
</p>

A collection of Go linters that help maintain code quality and prevent common mistakes.

## Available Analyzers

### 1. Mutex Usage (mutexoid)
Checks for incorrect mutex usage patterns, specifically focusing on preventing mutex copying.

### 2. English Comments (englishcomments)
Ensures all code comments are written in English, helping maintain consistency in international projects.

### 3. Test Package (testpackage)
Ensures that all test files (*_test.go) use package names with '_test' suffix to properly separate test packages from the main ones.

### 4. Parallel Tests (paralleltests)
Ensures that all tests are executed in parallel and follow table-driven test patterns for better performance and maintainability.
- Checks for `t.Parallel()` calls in test functions
- Enforces table-driven test patterns
- Improves test execution speed
- Helps catch race conditions early

### 5. Test Cases (testcases)
Ensures that table-driven tests include both positive and negative test cases for better test coverage.
- Checks for at least one positive test case
- Checks for at least one negative test case
- Helps maintain comprehensive test suites
- Improves code quality through better testing

### 6. Config Validator (configvalidator)
Ensures that struct fields with configuration tags (koanf/json) also have validation tags for better configuration validation.
- Checks for validation tags on config fields
- Helps prevent invalid configurations
- Improves configuration reliability
- Enforces consistent validation practices

### 7. Useless Comments (uselesscomments)
Detects and reports redundant comments that don't add value to code understanding.
- Identifies comments that just repeat the code structure
- Catches obvious descriptions that don't provide additional context
- Helps maintain meaningful documentation

Examples of issues detected:
```go
// Bad - will trigger warnings
// UserService is a service for users
type UserService struct{}

// GetUser gets the user
func GetUser(id string) string {}

// ServiceInterface extends the base interface
type ServiceInterface interface {
    // GetData gets the data
    GetData() string
}

// Good - no warnings
// UserService handles user authentication and authorization with OAuth2
type UserService struct{}

// GetUser retrieves user data from cache or database according to configured priority
func GetUser(id string) string {}

// ServiceInterface defines the contract for handling complex data transformations
type ServiceInterface interface {
    // GetData retrieves and validates data according to business rules
    GetData() string
}
```

Common patterns that trigger warnings:
- Comments that just repeat the type/function/method name
- "is a type that...", "is a struct that...", "is an interface that..."
- "implements the...", "extends the..."
- "method that...", "function that..."
- Simple getter/setter descriptions
- Comments shorter than 3 words that just mention the element name

Exceptions (not flagged as useless):
- Comments containing implementation details
- Descriptions of business logic
- Comments with specific terms like "performs", "validates", "calculates"
- Comments explaining "according to" some rules or conditions

## Installation

### Global Installation (Recommended)
```bash
# Install all analyzers
go install github.com/Synapse-Devs/mutexoid/cmd/all@latest

# Or install individual analyzers
go install github.com/Synapse-Devs/mutexoid/cmd/mutexoid@latest        # Only mutex checks
go install github.com/Synapse-Devs/mutexoid/cmd/englishcomments@latest # Only English comments checks
go install github.com/Synapse-Devs/mutexoid/cmd/testpackage@latest    # Only test package checks
go install github.com/Synapse-Devs/mutexoid/cmd/paralleltests@latest  # Only parallel tests checks
go install github.com/Synapse-Devs/mutexoid/cmd/testcases@latest     # Only test cases checks
go install github.com/Synapse-Devs/mutexoid/cmd/configvalidator@latest # Only config validation checks
go install github.com/Synapse-Devs/mutexoid/cmd/uselesscomments@latest # Only useless comments checks
```

### Local Installation
```bash
# Clone the repository
git clone https://github.com/Synapse-Devs/mutexoid.git
cd mutexoid

# Build and install locally
go install ./cmd/all         # Install all analyzers
go install ./cmd/mutexoid    # Only mutex checks
go install ./cmd/englishcomments # Only English comments checks
go install ./cmd/testpackage    # Only test package checks
go install ./cmd/paralleltests  # Only parallel tests checks
go install ./cmd/testcases      # Only test cases checks
go install ./cmd/configvalidator # Only config validation checks
go install ./cmd/uselesscomments # Only useless comments checks
```

## Usage

### Standalone

Run all analyzers:
```bash
all ./...
```

Run individual analyzers:
```bash
mutexoid ./...         # Only check mutex usage
englishcomments ./...  # Only check English comments
testpackage ./...      # Only check test package names
paralleltests ./...    # Only check parallel tests
testcases ./...        # Only check test cases
configvalidator ./...  # Only check config validation
uselesscomments ./...  # Only check useless comments
```

### Development Mode
If you're working on the codebase directly, you can run analyzers without installation:
```bash
go run cmd/all/main.go ./...            # Run all analyzers
go run cmd/mutexoid/main.go ./...       # Only check mutex usage
go run cmd/englishcomments/main.go ./... # Only check English comments
go run cmd/testpackage/main.go ./...    # Only check test package names
go run cmd/paralleltests/main.go ./...  # Only check parallel tests
go run cmd/testcases/main.go ./...      # Only check test cases
go run cmd/configvalidator/main.go ./... # Only check config validation
go run cmd/uselesscomments/main.go ./... # Only check useless comments
```

### With golangci-lint

Note: Plugin mode is currently not recommended on macOS due to compatibility issues.

1. Build the plugin:
```bash
go build -buildmode=plugin -o path/to/mutexoid.so github.com/Synapse-Devs/mutexoid/cmd/mutexoid/plugin
```

2. Add to your `.golangci.yml`:
```yaml
linters:
  enable:
    - mutexoid        # For mutex checks
    - englishcomments # For English comments checks
    - testpackage     # For test package checks
    - paralleltests   # For parallel tests checks
    - testcases       # For test cases checks
    - configvalidator # For config validation checks
    - uselesscomments # For useless comments checks
  
linters-settings:
  custom:
    mutexoid:
      path: path/to/mutexoid.so
      description: Collection of code quality analyzers
      original-url: github.com/Synapse-Devs/mutexoid
```

## Why it matters

### Mutex Safety
- Prevents race conditions
- Avoids deadlocks
- Reduces hard-to-debug production issues

### English Comments
- Improves code maintainability
- Ensures better collaboration in international teams
- Makes code more accessible to the global Go community

### Test Package Separation
- Keeps test code separate from production code
- Prevents test dependencies from being included in production builds
- Improves package organization and maintainability

### Parallel Tests
- Significantly reduces test execution time
- Helps identify race conditions early
- Enforces consistent test patterns
- Makes tests more maintainable and readable

### Test Cases
- Improves code quality through better testing
- Helps maintain comprehensive test suites

### Useless Comments
- Reduces code noise and improves readability
- Encourages meaningful documentation
- Promotes self-documenting code practices
- Makes important comments stand out
- Helps maintain high-quality documentation

## License

MIT 