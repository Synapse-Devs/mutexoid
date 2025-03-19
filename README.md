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
```

### Development Mode
If you're working on the codebase directly, you can run analyzers without installation:
```bash
go run cmd/all/main.go ./...            # Run all analyzers
go run cmd/mutexoid/main.go ./...       # Only check mutex usage
go run cmd/englishcomments/main.go ./... # Only check English comments
go run cmd/testpackage/main.go ./...    # Only check test package names
go run cmd/paralleltests/main.go ./...  # Only check parallel tests
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
  
linters-settings:
  custom:
    mutexoid:
      path: path/to/mutexoid.so
      description: Collection of code quality analyzers
      original-url: github.com/Synapse-Devs/mutexoid
```

## Examples

### Mutex Usage

```go
// Bad - will trigger warnings
type BadStruct struct {
    sync.Mutex      // Warning: mutex should be a pointer type
    mu sync.Mutex   // Warning: mutex should be a pointer type
}

// Good - no warnings
type GoodStruct struct {
    mu *sync.Mutex
}
```

### English Comments

```go
// Good - English comment
func Calculate() int {
    return 42
}

// Плохо - не английский комментарий (Warning: comment should be in English)
func BadExample() int {
    return 42
}
```

### Test Package Names

```go
// Bad - will trigger warning
package mypackage // in file mypackage_test.go

// Good - no warning
package mypackage_test // in file mypackage_test.go
```

### Parallel Tests

```go
// Bad - will trigger warnings
func TestBad(t *testing.T) {
    // Warning: test should call t.Parallel()
    // Warning: test should use table-driven tests
    result := Calculate()
    if result != 42 {
        t.Error("wrong result")
    }
}

// Good - no warnings
func TestGood(t *testing.T) {
    t.Parallel() // Enable parallel execution

    tests := []struct{
        name string
        want int
    }{
        {"case 1", 42},
        {"case 2", 43},
    }

    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Enable parallel execution for subtests
            result := Calculate()
            if result != tt.want {
                t.Errorf("got %d, want %d", result, tt.want)
            }
        })
    }
}
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

## License

MIT 