package testcases

import (
	"fmt"
	"testing"
)

// Matrix2 runs a matrix of test cases for two sets of items.
// For each combination of items from items1 and items2, it runs a subtest
// with a title formatted using titleformat.
// The titleformat should contain two format verbs to include item1 and item2.
// Example:
//	Matrix2(t, "Test with %v and %v", []int{1, 2}, []string{"a", "b"}, func(t *testing.T, num int, str string) {
func Matrix2[T1 any, T2 any](t *testing.T, titleformat string, items1 []T1, items2 []T2, fn func(t *testing.T, item1 T1, item2 T2)) {
	t.Helper()

	for _, item1 := range items1 {
		for _, item2 := range items2 {
			title := fmt.Sprintf(titleformat, item1, item2)
			t.Run(title, func(t *testing.T) {
				fn(t, item1, item2)
			})
		}
	}
}

// Matrix3 runs a matrix of test cases for three sets of items.
// For each combination of items from items1, items2, and items3, it runs a subtest
// with a title formatted using titleformat.
// The titleformat should contain three format verbs to include item1, item2, and item3.
// Example:
//	Matrix3(t, "Test with %v, %v and %v", []int{1, 2}, []string{"a", "b"}, []bool{true, false}, func(t *testing.T, num int, str string, flag bool) {
func Matrix3[T1 any, T2 any, T3 any](t *testing.T, titleformat string, items1 []T1, items2 []T2, items3 []T3, fn func(t *testing.T, item1 T1, item2 T2, item3 T3)) {
	t.Helper()

	for _, item1 := range items1 {
		for _, item2 := range items2 {
			for _, item3 := range items3 {
				title := fmt.Sprintf(titleformat, item1, item2, item3)
				t.Run(title, func(t *testing.T) {
					fn(t, item1, item2, item3)
				})
			}
		}
	}
}

// Matrix4 runs a matrix of test cases for four sets of items.
// For each combination of items from items1, items2, items3, and items4, it runs a subtest
// with a title formatted using titleformat.
// The titleformat should contain four format verbs to include item1, item2, item3, and item4.
// Example:
//	Matrix4(t, "Test with %v, %v, %v and %v", []int{1, 2}, []string{"a", "b"}, []bool{true, false}, []float64{1.1, 2.2}, func(t *testing.T, num int, str string, flag bool, f float64) {
func Matrix4[T1 any, T2 any, T3 any, T4 any](t *testing.T, titleformat string, items1 []T1, items2 []T2, items3 []T3, items4 []T4, fn func(t *testing.T, item1 T1, item2 T2, item3 T3, item4 T4)) {
	t.Helper()
	for _, item1 := range items1 {
		for _, item2 := range items2 {
			for _, item3 := range items3 {
				for _, item4 := range items4 {
					title := fmt.Sprintf(titleformat, item1, item2, item3, item4)
					t.Run(title, func(t *testing.T) {
						fn(t, item1, item2, item3, item4)
					})
				}
			}
		}
	}
}

// Matrix5 runs a matrix of test cases for five sets of items.
// For each combination of items from items1, items2, items3, items4, and items5, it runs a subtest
// with a title formatted using titleformat.
// The titleformat should contain five format verbs to include item1, item2, item3, item4, and item5.
// Example:
//	Matrix5(t, "Test with %v, %v, %v, %v and %v", []int{1, 2}, []string{"a", "b"}, []bool{true, false}, []float64{1.1, 2.2}, []byte{0, 1}, func(t *testing.T, num int, str string, flag bool, f float64, b byte) {
func Matrix5[T1 any, T2 any, T3 any, T4 any, T5 any](t *testing.T, titleformat string, items1 []T1, items2 []T2, items3 []T3, items4 []T4, items5 []T5, fn func(t *testing.T, item1 T1, item2 T2, item3 T3, item4 T4, item5 T5)) {
	t.Helper()
	for _, item1 := range items1 {
		for _, item2 := range items2 {
			for _, item3 := range items3 {
				for _, item4 := range items4 {
					for _, item5 := range items5 {
						title := fmt.Sprintf(titleformat, item1, item2, item3, item4, item5)
						t.Run(title, func(t *testing.T) {
							fn(t, item1, item2, item3, item4, item5)
						})
					}
				}
			}
		}
	}
}