// nolint:thelper // because of t.Helper usage in testcases package
package testcases_test

import (
	"strconv"
	"testing"

	"github.com/daanv2/go-testcases"
)

func TestDefineTableAndRun(t *testing.T) {
	tbl := testcases.DefineTable(
		testcases.TestCase[int, string]{Title: "One", Expected: 1, Actual: "1"},
		testcases.TestCase[int, string]{Title: "Two", Expected: 2, Actual: "2"},
	)
	var called []string
	tbl.Run(t, func(t *testing.T, expected int, actual string) {
		if strconv.Itoa(expected) != actual {
			t.Errorf("expected %d, got %s", expected, actual)
		}
		called = append(called, actual)
	})
	if len(called) != 2 {
		t.Errorf("expected 2 test cases to run, got %d", len(called))
	}
}

// ExampleDefineTable demonstrates usage of DefineTable and Table.Run.
func ExampleDefineTable() {
	tbl := testcases.DefineTable(
		testcases.TestCase[int, string]{Title: "One", Expected: 1, Actual: "1"},
		testcases.TestCase[int, string]{Title: "Two", Expected: 2, Actual: "2"},
	)
	tbl.Run(&testing.T{}, func(t *testing.T, expected int, actual string) {
		println("Test:", expected, actual)
	})
}
