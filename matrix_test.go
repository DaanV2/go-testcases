// nolint:thelper // because of t.Helper usage in testcases package
package testcases_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/daanv2/go-testcases"
)

func TestMatrix2(t *testing.T) {
	var results []string
	testcases.Matrix2(t, "%v-%v", []int{1, 2}, []string{"a", "b"}, func(t *testing.T, i int, s string) {
	 	results = append(results, strings.ToUpper(s)+strconv.Itoa(i))
	})
	if len(results) != 4 {
		t.Errorf("expected 4 results, got %d", len(results))
	}
}

func TestMatrix3(t *testing.T) {
	var results []string
	testcases.Matrix3(t, "%v-%v-%v", []int{1}, []string{"x"}, []bool{true, false}, func(t *testing.T, i int, s string, b bool) {
	 	results = append(results, s+"-"+strconv.Itoa(i)+"-"+map[bool]string{true:"T",false:"F"}[b])
	})
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

func TestMatrix4(t *testing.T) {
	var count int
	testcases.Matrix4(t, "%v-%v-%v-%v", []int{1}, []string{"a"}, []bool{true}, []float64{1.1, 2.2}, func(t *testing.T, i int, s string, b bool, f float64) {
		count++
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func TestMatrix5(t *testing.T) {
	var count int
	testcases.Matrix5(t, "%v-%v-%v-%v-%v", []int{1}, []string{"a"}, []bool{true}, []float64{1.1}, []byte{0, 1}, func(t *testing.T, i int, s string, b bool, f float64, by byte) {
		count++
	})
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}
