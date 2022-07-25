package primefactor

import (
	"reflect"
	"testing"
)

func TestPrimeFactor1(t *testing.T) {
	result := PrimeFactors(1)
	expectedResult := []int{}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor2(t *testing.T) {
	result := PrimeFactors(2)
	expectedResult := []int{2}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor3(t *testing.T) {
	result := PrimeFactors(3)
	expectedResult := []int{3}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor4(t *testing.T) {
	result := PrimeFactors(4)
	expectedResult := []int{2, 2}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor5(t *testing.T) {
	result := PrimeFactors(5)
	expectedResult := []int{5}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor6(t *testing.T) {
	result := PrimeFactors(6)
	expectedResult := []int{2, 3}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor8(t *testing.T) {
	result := PrimeFactors(8)
	expectedResult := []int{2, 2, 2}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor9(t *testing.T) {
	result := PrimeFactors(9)
	expectedResult := []int{3, 3}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor10(t *testing.T) {
	result := PrimeFactors(10)
	expectedResult := []int{2, 5}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor12(t *testing.T) {
	result := PrimeFactors(12)
	expectedResult := []int{2, 2, 3}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}

func TestPrimeFactor25(t *testing.T) {
	result := PrimeFactors(25)
	expectedResult := []int{5, 5}
	if !reflect.DeepEqual(expectedResult, result) {
		t.Errorf("expected %q, got %q", expectedResult, result)
	}
}
