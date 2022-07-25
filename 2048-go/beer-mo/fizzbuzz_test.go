package main

import (
	"strconv"
	"testing"
)

func TestRunFirst(t *testing.T) {
	strNum := PrintFizzBuzz(1)
	if strNum != "1" {
		t.Error("got", strNum, "1")
	}
}

func TestPrintNumberWhenCannotPrintFizzBuzz(t *testing.T) {
	strNum := PrintFizzBuzz(2)
	expected := strconv.Itoa(2)
	if strNum != expected {
		t.Errorf("Got %v\n want %v", strNum, expected)
	}
}

func TestPrintFizz(t *testing.T) {
	strFizz := PrintFizzBuzz(3)
	if strFizz != "fizz" {
		t.Errorf("Cannot print fizz")
	}
}

func TestPrintBuzz(t *testing.T) {
	strBuzz := PrintFizzBuzz(5)
	if strBuzz != "buzz" {
		t.Errorf("Cannot print buzz")
	}
}

func TestPrintFizzWithSix(t *testing.T) {
	strFizz := PrintFizzBuzz(6)
	if strFizz != "fizz" {
		t.Errorf("Cannot print fizz")
	}
}

func TestPrintBuzzWithTen(t *testing.T){
	strBuzz := PrintFizzBuzz(10)
	if strBuzz != "buzz"{
		t.Errorf("Cannot print buzz")
	}
}

func TestPrintFizzBuzz(t *testing.T) {
	strFizzBuzz := PrintFizzBuzz(15)
	if strFizzBuzz != "fizzbuzz" {
		t.Errorf("Cannot print fizzbuzz")
	}
}
