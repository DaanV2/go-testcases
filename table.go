package testcases

import "testing"

type TestCase[Exptected any, Actual any] struct {
	Title    string
	Expected Exptected
	Actual   Actual
}

type Table[Exptected any, Actual any] struct {
	Cases []TestCase[Exptected, Actual]
}

// DefineTable creates a Table of test cases from the provided items.
// Example:
//	tbl := DefineTable[int, string](
//	    TestCase[int, string]{Title: "Test 1", Expected: 1, Actual: "one"},
//	    TestCase[int, string]{Title: "Test 2", Expected: 2, Actual: "two"},
//	)
//	tbl.Run(t, func(t *testing.T, expected int, actual string) {
// 	   // test logic here
//	})
func DefineTable[Exptected any, Actual any](items ...TestCase[Exptected, Actual]) Table[Exptected, Actual] {
	return Table[Exptected, Actual]{Cases: items}
}

// Run executes the test cases defined in the Table.
func (tbl Table[Exptected, Actual]) Run(t *testing.T, fn func(t *testing.T, expected Exptected, actual Actual)) {
	t.Helper()

	for _, tc := range tbl.Cases {
		t.Run(tc.Title, func(t *testing.T) {
			fn(t, tc.Expected, tc.Actual)
		})
	}
}
