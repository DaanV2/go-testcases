package testcases

import (
	"fmt"
	"testing"
)

// Each runs a set of test cases for a single set of items.
// For each item in items, it runs a subtest
// with a title formatted using titleformat.
func Each[T any](t *testing.T, titleformat string, items []T, fn func(t *testing.T, item T)) {
	t.Helper()

	for _, item1 := range items {
		title := fmt.Sprintf(titleformat, item1)
		t.Run(title, func(t *testing.T) {
			fn(t, item1)
		})
	}
}