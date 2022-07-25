package j_bas

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPrimeFactor1_ReturnEmpty(t *testing.T) {

	testData := []struct { 
		input int
		expectd []int
	}{
		{ 1, []int{} },
		{ 2, []int{2} },
		{ 3, []int{3} },
		{ 4, []int{2, 2} },
		{ 5, []int{5} },
		{ 6, []int{2, 3} },
		{ 7, []int{7} },
		{ 8, []int{2, 2, 2} },
		{ 9, []int{3, 3} },
		{ 10, []int{2, 5} },
		{ 18, []int{2, 3, 3} },
		{ 25, []int{5, 5} },
		{ 27, []int{3, 3, 3} },
		{ 1001, []int{7, 11, 13} },
	}

	for _, test := range testData {
		assert.Equal(t, test.expectd, PrimeFactor(test.input))
	}
}

func TestPrimesOf(t *testing.T) {
	testData := []struct { 
		input int
		expectd []int
	}{
		{ 1, []int{} },
		{ 2, []int{2} },
		{ 3, []int{2, 3} },
	}

	for _, test := range testData {
		assert.Equal(t, test.expectd, PrimesOf(test.input))
	}
}

func TestIsPrime(t *testing.T) {
	testData := []struct { 
		input int
		expectd bool
	}{
		{ 1, false },
		{ 2, true },
		{ 3, true },
		{ 4, false },
		{ 5, true },
		{ 7, true },
		{ 11, true },
		{ 13, true },
	}

	for _, test := range testData {
		assert.Equal(t, test.expectd, IsPrime(test.input), fmt.Sprintf("given %d got %v", test.input, test.expectd))
	}
}

// func TestPrimeFactor2_ReturnListContains2(t *testing.T) {
// 	assert.Equal(t, []int{2}, PrimeFactor(2))
// }

// func TestPrimeFactor3_ReturnListContains3(t *testing.T) {
// 	assert.Equal(t, []int{3}, PrimeFactor(3))
// }

// func TestPrimeFactor4_ReturnListContains2_2(t *testing.T) {
// 	assert.Equal(t, []int{2, 2}, PrimeFactor(4), )
// }

// func TestPrimeFactor6_ReturnListContains2_3(t *testing.T) {
// 	assert.Equal(t, []int{2, 3}, PrimeFactor(6))
// }

// func TestPrimeFactor8_ReturnListContains2_2_2(t *testing.T) {
// 	assert.Equal(t, []int{2, 2, 2}, PrimeFactor(8))
// }

// func TestPrimeFactor9_ReturnListContains3_3(t *testing.T) {
// 	assert.Equal(t, []int{3, 3}, PrimeFactor(9))
// }

// func TestPrimeFactor27_ReturnListContains3_3_3(t *testing.T) {
// 	assert.Equal(t, []int{3, 3, 3}, PrimeFactor(27))
// }

// func TestPrimeFactor5_ReturnListContains5(t *testing.T) {
// 	assert.Equal(t, []int{5}, PrimeFactor(5))
// }

// func TestPrimeFactor25_ReturnListContains5_5(t *testing.T) {
// 	assert.Equal(t, []int{5, 5}, PrimeFactor(25))
// }

// func TestPrimeFactor10_ReturnListContains2_5(t *testing.T) {
// 	assert.Equal(t, []int{2, 5}, PrimeFactor(10))
// }

// func TestPrimeFactor18_ReturnListContains2_3_3(t *testing.T) {
// 	assert.Equal(t, []int{2, 3, 3}, PrimeFactor(18))
// }

// func TestPrimeFactor7_ReturnListContains7(t *testing.T) {
// 	assert.Equal(t, []int{7}, PrimeFactor(7))
// }