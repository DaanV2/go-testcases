package testcases

import "testing"

// TestCase represents a single test case with expected and actual values.
type TestCase[Exptected any, Actual any] struct {
	Title  string
	Input  Exptected
	Result Actual
}

// NewCase creates a new TestCase instance.
func NewCase[Exptected any, Actual any](title string, input Exptected, result Actual) TestCase[Exptected, Actual] {
	return TestCase[Exptected, Actual]{
		Title:  title,
		Input:  input,
		Result: result,
	}
}

// Table holds a collection of test cases.
type Table[Exptected any, Actual any] struct {
	Cases []TestCase[Exptected, Actual]
}

// DefineTable creates a Table of test cases from the provided items.
// Example:
//	tbl := testcases.DefineTable[int, string](
//	    testcases.TestCase[int, string]{Title: "Test 1", Expected: 1, Actual: "one"},
//	    testcases.TestCase[int, string]{Title: "Test 2", Expected: 2, Actual: "two"},
//	)
//	tbl.Run(t, func(t *testing.T, expected int, actual string) {
// 	   // test logic here
//	})
func DefineTable[Exptected any, Actual any](items ...TestCase[Exptected, Actual]) Table[Exptected, Actual] {
	return Table[Exptected, Actual]{Cases: items}
}

// Add appends a new test case to the Table.
func (t *Table[Exptected, Actual]) Add(title string, input Exptected, result Actual) *Table[Exptected, Actual] {
	t.Cases = append(t.Cases, NewCase(title, input, result))
	return t
}

// Run executes the test cases defined in the Table.
func (tbl Table[Exptected, Actual]) Run(t *testing.T, fn func(t *testing.T, input Exptected, result Actual)) {
	t.Helper()

	for _, tc := range tbl.Cases {
		t.Run(tc.Title, func(t *testing.T) {
			fn(t, tc.Input, tc.Result)
		})
	}
}
