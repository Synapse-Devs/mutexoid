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

## Installation

### Global Installation (Recommended)
```bash
# Install all analyzers
go install github.com/Synapse-Devs/mutexoid/cmd/all@latest

# Or install individual analyzers
go install github.com/Synapse-Devs/mutexoid/cmd/mutexoid@latest        # Only mutex checks
go install github.com/Synapse-Devs/mutexoid/cmd/englishcomments@latest # Only English comments checks
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
```

### Development Mode
If you're working on the codebase directly, you can run analyzers without installation:
```bash
go run cmd/all/main.go ./...            # Run all analyzers
go run cmd/mutexoid/main.go ./...       # Only check mutex usage
go run cmd/englishcomments/main.go ./... # Only check English comments
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

## Why it matters

### Mutex Safety
- Prevents race conditions
- Avoids deadlocks
- Reduces hard-to-debug production issues

### English Comments
- Improves code maintainability
- Ensures better collaboration in international teams
- Makes code more accessible to the global Go community

## License

MIT 