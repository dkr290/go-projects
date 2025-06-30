# Contributing to Go Auto Swagger

Thank you for your interest in contributing to Go Auto Swagger! 🎉 We welcome contributions from the community and are excited to see what you'll bring to the project.

## 📋 Table of Contents

- [Getting Started](#getting-started)
- [How to Contribute](#how-to-contribute)
- [Development Setup](#development-setup)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)
- [Issue Guidelines](#issue-guidelines)
- [Community](#community)

## 🚀 Getting Started

### Prerequisites

- Go 1.19 or higher
- Git
- A GitHub account

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:

   ```bash
   git clone https://github.com/YOUR_USERNAME/go-auto-swagger.git
   cd go-auto-swagger
   ```

3. Add the original repository as upstream:
   ```bash
   git remote add upstream https://github.com/dani/go-auto-swagger.git
   ```

## 🛠️ How to Contribute

### Types of Contributions

We welcome several types of contributions:

- 🐛 **Bug Reports** - Help us identify and fix issues
- ✨ **Feature Requests** - Suggest new functionality
- 🔧 **Code Contributions** - Submit bug fixes or new features
- 📖 **Documentation** - Improve or add documentation
- 🧪 **Tests** - Add or improve test coverage
- 💡 **Examples** - Create examples for different use cases

### Areas We Need Help With

- **Router Adapters**: Support for Gin, Echo, Fiber, Gorilla Mux
- **Schema Generation**: Enhanced struct tag support
- **Validation**: Parameter and request body validation
- **Security**: Authentication/authorization documentation
- **Performance**: Optimization and benchmarking
- **Examples**: Real-world usage examples

## 💻 Development Setup

1. **Install Dependencies**:

   ```bash
   go mod download
   ```

2. **Build the Project**:

   ```bash
   go build ./...
   ```

3. **Run Examples**:

   ```bash
   # Basic example
   go run examples/main.go

   # Chi example
   go run examples/chi/main.go

   # Net/HTTP example
   go run examples/nethttp/main.go
   ```

4. **Verify Everything Works**:
   ```bash
   go test ./...
   ```

## 🔄 Pull Request Process

### Before You Start

1. **Check existing issues** - See if someone is already working on it
2. **Create an issue** - For new features or major changes
3. **Get feedback** - Discuss your approach before coding

### Step by Step

1. **Create a Feature Branch**:

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make Your Changes**:
   - Write clean, readable code
   - Follow the coding standards
   - Add tests for new functionality
   - Update documentation as needed

3. **Test Your Changes**:

   ```bash
   # Run tests
   go test ./...

   # Test examples
   go run examples/main.go

   # Check formatting
   go fmt ./...

   # Run linter (if available)
   golangci-lint run
   ```

4. **Commit Your Changes**:

   ```bash
   git add .
   git commit -m "feat: add support for XYZ router"
   ```

   Use conventional commit messages:
   - `feat:` - New features
   - `fix:` - Bug fixes
   - `docs:` - Documentation changes
   - `test:` - Test additions/changes
   - `refactor:` - Code refactoring
   - `style:` - Code formatting
   - `chore:` - Maintenance tasks

5. **Push and Create PR**:

   ```bash
   git push origin feature/your-feature-name
   ```

   Then create a Pull Request on GitHub.

### PR Checklist

- [ ] Code follows the project's coding standards
- [ ] Tests are added for new functionality
- [ ] Documentation is updated
- [ ] Examples are updated if needed
- [ ] All tests pass
- [ ] Code is properly formatted
- [ ] PR description clearly explains the changes
- [ ] Linked to relevant issue (if applicable)

## 📏 Coding Standards

### Go Standards

- Follow standard Go conventions and idioms
- Use `gofmt` for formatting
- Use meaningful variable and function names
- Write clear comments for exported functions
- Keep functions small and focused

### Project Specific

```go
// Good: Clear function names
func (rb *RouteBuilder) Summary(summary string) *RouteBuilder {
    rb.info.Summary = summary
    return rb
}

// Good: Proper error handling
func (s *Swagger) ToJSON() ([]byte, error) {
    return json.MarshalIndent(s, "", "  ")
}

// Good: Clear struct tags
type User struct {
    ID   int    `json:"id" description:"User ID" example:"1"`
    Name string `json:"name" description:"User name" example:"John Doe"`
}
```

### File Organization

```
go-auto-swagger/
├── go-auto-swagger/          # Main package
│   ├── auto.go              # Core AutoSwagger struct
│   ├── builder.go           # RouteBuilder implementation
│   ├── swagger.go           # OpenAPI spec structures
│   ├── adapters.go          # Net/HTTP adapter
│   ├── chi_adapter.go       # Chi adapter
│   └── middleware.go        # Middleware utilities
├── examples/                # Usage examples
│   ├── main.go             # Basic example
│   ├── chi/                # Chi router example
│   └── nethttp/            # Advanced net/http example
└── docs/                   # Additional documentation
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -run TestFunctionName ./...
```

### Writing Tests

```go
func TestRouteBuilder_Summary(t *testing.T) {
    swagger := New("Test API", "1.0.0")
    builder := swagger.GET("/test")

    result := builder.Summary("Test Summary")

    if result.info.Summary != "Test Summary" {
        t.Errorf("Expected 'Test Summary', got '%s'", result.info.Summary)
    }
}
```

### Test Guidelines

- Write tests for all new functionality
- Include edge cases and error conditions
- Use table-driven tests for multiple scenarios
- Mock external dependencies
- Keep tests simple and focused

## 📖 Documentation

### Code Documentation

- Document all exported functions, types, and constants
- Use clear, concise descriptions
- Include examples in documentation

```go
// RouteBuilder helps build routes with automatic documentation.
// It provides a fluent interface for configuring route metadata
// such as summary, description, parameters, and response types.
//
// Example:
//   builder := swagger.GET("/users/{id}")
//   builder.Summary("Get User").Path("id", "User ID", IntSchema)
type RouteBuilder struct {
    // ...
}
```

### README and Guides

- Keep examples up to date
- Add new features to the README
- Include real-world usage scenarios
- Update API reference documentation

## 🐛 Issue Guidelines

### Bug Reports

When reporting bugs, please include:

- **Go version**: `go version`
- **Operating System**: OS and version
- **Expected behavior**: What should happen
- **Actual behavior**: What actually happens
- **Minimal reproduction**: Code to reproduce the issue
- **Error messages**: Full error output

### Feature Requests

When requesting features, please include:

- **Use case**: Why you need this feature
- **Proposed solution**: How it might work
- **Alternatives**: Other ways to achieve the goal
- **Examples**: How the API might look

### Good Issue Template

```markdown
## Description

Brief description of the issue/feature

## Environment

- Go version: 1.21
- OS: macOS 14.0
- Router: Chi v5.0.0

## Steps to Reproduce

1. Create a new project
2. Add this code...
3. Run the server
4. See error

## Expected Behavior

What should happen

## Actual Behavior

What actually happens

## Additional Context

Any other relevant information
```

## 💬 Community

### Getting Help

- **GitHub Issues**: For bugs and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Code Review**: We welcome code review on PRs

### Recognition

Contributors will be:

- Listed in the README contributors section
- Mentioned in release notes for significant contributions
- Thanked in commit messages

## 🏷️ Release Process

### Versioning

We use [Semantic Versioning](https://semver.org/):

- **MAJOR**: Incompatible API changes
- **MINOR**: New functionality (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

### Release Checklist

- [ ] All tests pass
- [ ] Documentation is updated
- [ ] Examples work with new version
- [ ] CHANGELOG is updated
- [ ] Version is tagged
- [ ] Release notes are written

## 🙏 Thank You

Every contribution, no matter the size, is valuable and appreciated. Thank you for helping make Go Auto Swagger better for everyone!

---

**Questions?** Feel free to ask in the issues or discussions section. We're here to help! 🚀

