# CodeChecker

A custom code quality checker for the Cryptellation CI pipeline that enforces coding standards and best practices.

## Overview

CodeChecker is a Go-based tool that performs automated code quality checks to ensure consistency and maintainability across the Cryptellation codebase. It integrates with CI/CD pipelines to catch issues early in the development process.

## Features

### 1. TODO Validation
- **Valid Format**: Ensures TODOs follow the required format `TODO(#123)` where `#123` is a GitHub issue number
- **File Coverage**: Scans `.go` and `.yaml` files for invalid TODO comments
- **Exclusions**: Automatically excludes generated files (`.gen.go`) and internal Dagger files
- **Detailed Reporting**: Provides file paths and line numbers for invalid TODOs

### 2. Test Tag Validation
- **Build Tags**: Ensures test files (`*_test.go`) have proper build tags (`// +build`)
- **Consistency**: Enforces consistent test organization across the codebase
- **Compliance**: Helps maintain proper test structure and organization

## Installation

### Prerequisites
- Go 1.23.8 or higher
- Docker (for containerized builds)

### Local Development
```bash
# Clone the repository
git clone <repository-url>
cd codechecker

# Install dependencies
go mod download

# Build the tool
go build -o codechecker .

# Make it available globally (optional)
go install .
```

### Docker
```bash
# Build the Docker image
docker build -t codechecker .

# Run the container
docker run codechecker [options]
```

## Usage

### Basic Usage
```bash
# Run all checks (default)
codechecker

# Run with custom path
codechecker --path ./src

# Run specific checks
codechecker --check-invalid-todos=false --check-test-tags=true
```

### Command Line Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--path` | `-p` | `.` | Path to check for issues |
| `--check-invalid-todos` | | `true` | Enable/disable TODO validation |
| `--check-test-tags` | | `true` | Enable/disable test tag validation |

### Subcommands

#### Check TODOs
```bash
# Check for invalid TODOs
codechecker todos

# Check with custom path
codechecker todos --path ./internal
```

#### Check Test Tags
```bash
# Check for missing test tags
codechecker test-tags

# Check with custom path
codechecker test-tags --path ./pkg
```

## Examples

### Valid TODO Format
```go
// ✅ Valid - includes issue number
// TODO(#123): Implement user authentication

// ❌ Invalid - missing issue number
// TODO: Add error handling
```

### Valid Test File
```go
// ✅ Valid - includes build tag
// +build integration

package mypackage

import "testing"

func TestMyFunction(t *testing.T) {
    // test implementation
}
```

## CI/CD Integration

CodeChecker is designed to integrate seamlessly with CI/CD pipelines. It returns appropriate exit codes:

- **Exit 0**: All checks passed
- **Exit 1**: Issues found or errors occurred

### GitHub Actions Example
```yaml
- name: Run CodeChecker
  run: |
    docker run ghcr.io/cryptellation/codechecker:latest
```

### Local Pre-commit Hook
```bash
#!/bin/bash
# .git/hooks/pre-commit

if ! codechecker; then
    echo "CodeChecker found issues. Please fix them before committing."
    exit 1
fi
```

## Configuration

The tool uses sensible defaults but can be customized through command-line flags. No configuration files are required.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run `codechecker` to ensure no issues
6. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Version

Current version: 1.0.0
