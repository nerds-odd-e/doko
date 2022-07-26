package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz1(t *testing.T) {
	got := FizzBuzz(1)
	want := "1"
	assert.Equal(t, got, want)
}

func TestFizzBuzz2(t *testing.T) {
	got := FizzBuzz(2)
	want := "2"
	assert.Equal(t, got, want)
}

func TestFizzBuzz3(t *testing.T) {
	got := FizzBuzz(3)
	want := "fizz"
	assert.Equal(t, got, want)
}

func TestFizzBuzz5(t *testing.T) {
	got := FizzBuzz(5)
	want := "buzz"
	assert.Equal(t, got, want)
}

func TestFizzBuzz6(t *testing.T) {
	got := FizzBuzz(6)
	want := "fizz"
	assert.Equal(t, got, want)
}
