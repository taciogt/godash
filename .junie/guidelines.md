# Development Guidelines for godash

This document provides guidelines and instructions for developing and contributing to the godash project.

## Build/Configuration Instructions

### Prerequisites

- Go >= 1.21
- [direnv](https://direnv.net/docs/installation.html) (required for environment management)
- [gotest](https://github.com/rakyll/gotest) (optional, for colorful test output)

### Setup

1. Clone the repository:
   ```shell
   git clone https://github.com/taciogt/godash.git
   cd godash
   ```

2. Create a `.env` file in the project root if it doesn't exist:
   ```shell
   touch .env
   ```

3. Install development dependencies:
   ```shell
   make setup
   ```
   This will install:
   - gotest: A colorful test runner
   - godoc: For generating and viewing documentation

## Testing Information

### Running Tests

To run all tests:
```shell
make test
```

To run specific tests:
```shell
go test -v -run "TestName"
```

To run benchmarks:
```shell
make bench
```

To generate and view test coverage:
```shell
make coverage-report
```

### Adding New Tests

The project follows Go's standard testing conventions with some additional patterns:

1. **Regular Tests**: 
   - Create a file named `<feature>_test.go` in the same package as the code being tested.
   - Use table-driven tests for comprehensive test cases.
   - Example:
     ```go
     func TestFeature(t *testing.T) {
         tests := []struct {
             name     string
             input    InputType
             expected OutputType
         }{
             {
                 name:     "test case description",
                 input:    inputValue,
                 expected: expectedValue,
             },
             // More test cases...
         }

         for _, tt := range tests {
             t.Run(tt.name, func(t *testing.T) {
                 result := Feature(tt.input)
                 if result != tt.expected {
                     t.Errorf("Feature(%v) = %v, expected %v", tt.input, result, tt.expected)
                 }
             })
         }
     }
     ```

2. **Example Tests**:
   - Create a file named `<feature>_example_test.go` in the `godash_test` package.
   - These serve as both tests and documentation.
   - Example:
     ```go
     package godash_test

     import (
         "fmt"
         "github.com/taciogt/godash"
     )

     func ExampleFeature() {
         result := godash.Feature(input)
         fmt.Println(result)
         
         // Output:
         // expected output
     }
     ```

3. **Benchmarks**:
   - Add benchmarks to measure performance.
   - Example:
     ```go
     func BenchmarkFeature(b *testing.B) {
         b.ReportAllocs()
         for i := 0; i < b.N; i++ {
             Feature(input)
         }
     }
     ```

### Test Example

Here's a complete example of adding a new feature with tests:

1. Create a new function in a file (e.g., `utils.go`):
   ```go
   package godash

   // IsEven returns true if the given number is even, false otherwise.
   func IsEven(n int) bool {
       return n%2 == 0
   }
   ```

2. Create regular tests in `utils_test.go`:
   ```go
   package godash

   import "testing"

   func TestIsEven(t *testing.T) {
       tests := []struct {
           name     string
           input    int
           expected bool
       }{
           {
               name:     "zero is even",
               input:    0,
               expected: true,
           },
           // More test cases...
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               result := IsEven(tt.input)
               if result != tt.expected {
                   t.Errorf("IsEven(%d) = %v, expected %v", tt.input, result, tt.expected)
               }
           })
       }
   }
   ```

3. Create example tests in `utils_example_test.go`:
   ```go
   package godash_test

   import (
       "fmt"
       "github.com/taciogt/godash"
   )

   func ExampleIsEven() {
       fmt.Println(godash.IsEven(0))
       fmt.Println(godash.IsEven(2))
       fmt.Println(godash.IsEven(3))
       
       // Output:
       // true
       // true
       // false
   }
   ```

4. Run the tests:
   ```shell
   go test -v -run "IsEven"
   ```

## Additional Development Information

### Code Style

- Follow standard Go code style and conventions.
- Use `gofmt` or `goimports` to format your code.
- Write clear and concise comments, especially for exported functions and types.

### Documentation

To generate and view the project documentation locally:
```shell
make doc
```

This will start a local documentation server at http://localhost:6060/pkg/github.com/taciogt/godash

### Makefile Commands

The project includes a Makefile with several useful commands:

- `make help`: Show available commands and their descriptions
- `make setup`: Install dependencies for local development
- `make test`: Run all tests
- `make bench`: Run benchmarks
- `make coverage-report`: Generate and display test coverage report
- `make doc`: Generate and serve project documentation

### Project Structure

- The project is organized as a single Go package with multiple files for different features.
- Each feature typically has:
  - A main implementation file (e.g., `set.go`)
  - A test file (e.g., `set_test.go`)
  - An example test file (e.g., `set_example_test.go`)

### Contributing

Before submitting a pull request:
1. Ensure all tests pass
2. Add tests for new functionality
3. Update documentation if necessary
4. Follow the code style guidelines