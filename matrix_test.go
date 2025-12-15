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
		results = append(results, s+"-"+strconv.Itoa(i)+"-"+map[bool]string{true: "T", false: "F"}[b])
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

// ExampleMatrix2 demonstrates Matrix2 usage.
func ExampleMatrix2() {
	numbers := []int{1, 2}
	words := []string{"x", "y"}

	testcases.Matrix2(&testing.T{}, "A=%v, B=%v", numbers, words, func(t *testing.T, a int, b string) {
		println("Test:", a, b)
	})
}

// ExampleMatrix3 demonstrates Matrix3 usage.
func ExampleMatrix3() {
	numbers := []int{1}
	words := []string{"x"}
	flags := []bool{true}

	testcases.Matrix3(&testing.T{}, "A=%v, B=%v, C=%v", numbers, words, flags, func(t *testing.T, a int, b string, c bool) {
		println("Test:", a, b, c)
	})
}

// ExampleMatrix4 demonstrates Matrix4 usage.
func ExampleMatrix4() {
	numbers := []int{1}
	words := []string{"x"}
	flags := []bool{true}
	decimals := []float64{1.1}

	testcases.Matrix4(&testing.T{}, "A=%v, B=%v, C=%v, D=%v", numbers, words, flags, decimals, func(t *testing.T, a int, b string, c bool, d float64) {
		println("Test:", a, b, c, d)
	})
}

// ExampleMatrix5 demonstrates Matrix5 usage.
func ExampleMatrix5() {
	numbers := []int{1}
	words := []string{"x"}
	flags := []bool{true}
	decimals := []float64{1.1}
	bytes := []byte{0}

	testcases.Matrix5(&testing.T{}, "A=%v, B=%v, C=%v, D=%v, E=%v", numbers, words, flags, decimals, bytes, func(t *testing.T, a int, b string, c bool, d float64, e byte) {
		println("Test:", a, b, c, d, e)
	})
}
