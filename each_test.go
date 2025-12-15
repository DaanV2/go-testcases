// nolint:thelper,testableexamples // because of t.Helper usage in testcases package
package testcases_test

import (
	"testing"

	"github.com/daanv2/go-testcases"
)

func TestEach_Ints(t *testing.T) {
	items := []int{1, 2, 3}
	var seen []int
	testcases.Each(t, "item=%d", items, func(t *testing.T, item int) {
		seen = append(seen, item)
	})
	if len(seen) != len(items) {
		t.Errorf("Expected %d items, got %d", len(items), len(seen))
	}
}

func TestEach_Strings(t *testing.T) {
	items := []string{"a", "b", "c"}
	var seen []string
	testcases.Each(t, "item=%s", items, func(t *testing.T, item string) {
		seen = append(seen, item)
	})
	if len(seen) != len(items) {
		t.Errorf("Expected %d items, got %d", len(items), len(seen))
	}
}

func TestEach_Empty(t *testing.T) {
	items := []int{}
	testcases.Each(t, "item=%d", items, func(t *testing.T, item int) {
		t.Errorf("Should not be called for empty slice")
	})
}
