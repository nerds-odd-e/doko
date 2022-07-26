package main

import (
	"reflect"
	"testing"
)

func TestFactorOf1(t *testing.T) {
	input := 1
	expect := []int{}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf2(t *testing.T) {
	input := 2
	expect := []int{2}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf3(t *testing.T) {
	input := 3
	expect := []int{3}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf4(t *testing.T) {
	input := 4
	expect := []int{2, 2}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf6(t *testing.T) {
	input := 6
	expect := []int{2, 3}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf9(t *testing.T) {
	input := 9
	expect := []int{3, 3}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf8(t *testing.T) {
	input := 8
	expect := []int{2, 2, 2}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf12(t *testing.T) {
	input := 12
	expect := []int{2, 2, 3}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf27(t *testing.T) {
	input := 27
	expect := []int{3, 3, 3}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf10(t *testing.T) {
	input := 10
	expect := []int{2, 5}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf25(t *testing.T) {
	input := 25
	expect := []int{5, 5}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf50(t *testing.T) {
	input := 50
	expect := []int{2, 5, 5}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}

func TestFactorOf13(t *testing.T) {
	input := 13
	expect := []int{13}

	assert := PrimeFactors(input)

	if !reflect.DeepEqual(expect, assert) {
		t.Errorf("Expect %v but got %v .", expect, assert)
	}
}