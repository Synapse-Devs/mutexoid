# Mutexoid

<p align="center">
  <img src="/assets/logo.svg" width="200" height="200" alt="Mutexoid Logo">
</p>

A Go linter that checks for incorrect mutex usage patterns, specifically focusing on preventing mutex copying.

## Installation

```bash
go install github.com/Synapse-Devs/mutexoid/cmd/mutexoid@latest
```

## Usage

```bash
mutexoid ./...
```

## What it checks

1. Ensures that `sync.Mutex` and `sync.RWMutex` are always used as pointer types in structs
2. Prevents accidental mutex copying that could lead to race conditions

## Example

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

## Why it matters

Copying mutexes can lead to:
- Race conditions
- Deadlocks
- Hard-to-debug production issues

## License

MIT 