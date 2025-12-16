// nolint:thelper,testableexamples // because of t.Helper usage in testcases package
package testcases_test

import (
	"strconv"
	"testing"

	"github.com/daanv2/go-testcases"
)

func TestDefineTableAndRun(t *testing.T) {
	tbl := testcases.DefineTable(
		testcases.NewCase("One", 1, "1"),
		testcases.NewCase("Two", 2, "2"),
	)
	var called []string
	tbl.Run(t, func(t *testing.T, input int, result string) {
		if strconv.Itoa(input) != result {
			t.Errorf("expected %d, got %s", input, result)
		}
		called = append(called, result)
	})
	if len(called) != 2 {
		t.Errorf("expected 2 test cases to run, got %d", len(called))
	}
}

// ExampleDefineTable demonstrates usage of DefineTable and Table.Run.
func ExampleDefineTable() {
	tbl := testcases.DefineTable(
		testcases.TestCase[int, string]{Title: "Two", Input: 2, Result: "2"},
	)
	tbl.Run(&testing.T{}, func(t *testing.T, input int, result string) {
		println("Test:", input, result)
	})
}

// ExampleDefineTable demonstrates usage of DefineTable and Table.Run.
func ExampleDefineTable_second() {
	tbl := testcases.DefineTable(
		testcases.NewCase("One", 1, "1"),
	)
	tbl.Run(&testing.T{}, func(t *testing.T, input int, result string) {
		println("Test:", input, result)
	})
}
