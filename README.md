# Go Testcases

Go Testcases provides utilities for writing table-driven and matrix-style tests in Go, making it easier to cover combinations of test cases with minimal boilerplate.

## Install

```bash
go get github.com/daanv2/go-testcases
```

## Features

- **Each**: Run subtests for each element in a single input slice.
- **MatrixN**: Run subtests for all combinations of multiple input slices (Matrix2, Matrix3, etc.).
- **Table**: Define and run table-driven tests with custom titles and expected/actual values.

## Examples

### Each Example

```go
testcases.Each(t, []int{1, 2, 3}, func(t *testing.T, num int) {
    // Your test logic here
    t.Logf("Testing number: %d", num)
})
```

### Matrix2 Example

```go
testcases.Matrix2(t, "Test with %v and %v", []int{1, 2}, []string{"a", "b"}, func(t *testing.T, num int, str string) {
	// Your test logic here
	t.Logf("Testing %d and %s", num, str)
})
```

### Matrix3 Example

```go
testcases.Matrix3(t, "Test with %v, %v and %v", []int{1, 2}, []string{"a", "b"}, []bool{true, false}, func(t *testing.T, num int, str string, flag bool) {
	// Your test logic here
	t.Logf("Testing %d, %s, %v", num, str, flag)
})
```

### Table Example

```go
tbl := testcases.DefineTable(
	testcases.TestCase[int, string]{Title: "One", Expected: 1, Actual: "1"},
	testcases.TestCase[int, string]{Title: "Two", Expected: 2, Actual: "2"},
)
tbl.Run(t, func(t *testing.T, expected int, actual string) {
	if strconv.Itoa(expected) != actual {
		t.Errorf("expected %d, got %s", expected, actual)
	}
})
```

## How to Contribute

Contributions are welcome! To contribute:

1. Fork the repository and create your branch from `main`.
2. Add or update tests for your changes.
3. Run `go test ./...` to ensure all tests pass.
4. Open a pull request with a clear description of your changes.

Please follow Go best practices and keep code well-documented.

---
Licensed under the MIT License.