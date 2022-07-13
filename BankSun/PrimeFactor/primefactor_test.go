package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InputPrimefactor(t *testing.T) {
	a := inputPrimefactor(1)
	b := []int64{}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor2(t *testing.T) {
	a := inputPrimefactor(2)
	b := []int64{2}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor3(t *testing.T) {
	a := inputPrimefactor(3)
	b := []int64{3}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor4(t *testing.T) {
	a := inputPrimefactor(4)
	b := []int64{2, 2}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor6(t *testing.T) {
	a := inputPrimefactor(6)
	b := []int64{2, 3}
	assert.Equal(t, a, b)
}
func Test_InputPrimefactor8(t *testing.T) {
	a := inputPrimefactor(8)
	b := []int64{2, 2, 2}
	assert.Equal(t, a, b)
}
func Test_InputPrimefactor9(t *testing.T) {
	a := inputPrimefactor(9)
	b := []int64{3, 3}
	assert.Equal(t, a, b)
}
func Test_InputPrimefactor10(t *testing.T) {
	a := inputPrimefactor(10)
	b := []int64{2, 5}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor12(t *testing.T) {
	a := inputPrimefactor(12)
	b := []int64{2, 2, 3}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor14(t *testing.T) {
	a := inputPrimefactor(14)
	b := []int64{2, 7}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor27(t *testing.T) {
	a := inputPrimefactor(27)
	b := []int64{3, 3, 3}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor25(t *testing.T) {
	a := inputPrimefactor(25)
	b := []int64{5, 5}
	assert.Equal(t, a, b)
}

func Test_InputPrimefactor99(t *testing.T) {
	a := inputPrimefactor(99)
	b := []int64{3, 3, 11}
	assert.Equal(t, a, b)
}
