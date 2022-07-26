package primefactor_test

import (
	"reflect"
	"sort"
	"testing"

	"tdd.com/v1/hok-pong/primefactor"
)

func Test_Primefactor_input_one_expected_empty_array(t *testing.T) {
	input := 1
	result := primefactor.PrimeFactor(input)
	expect := []int{}
	if len(result) != len(expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_two_expected_array_of_two(t *testing.T) {
	input := 2
	result := primefactor.PrimeFactor(input)
	expect := []int{2}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_three_expected_array_of_three(t *testing.T) {
	input := 3
	result := primefactor.PrimeFactor(input)
	expect := []int{3}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_four_expected_array_of_double_two(t *testing.T) {
	input := 4
	result := primefactor.PrimeFactor(input)
	expect := []int{2, 2}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_five_expected_array_of_five(t *testing.T) {
	input := 5
	result := primefactor.PrimeFactor(input)
	expect := []int{5}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_six_expected_array_of_two_and_three(t *testing.T) {
	input := 6
	result := primefactor.PrimeFactor(input)
	expect := []int{2, 3}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_seven_expected_array_of_seven(t *testing.T) {
	input := 7
	result := primefactor.PrimeFactor(input)
	expect := []int{7}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_eight_expected_array_of_tripple_two(t *testing.T) {
	input := 8
	result := primefactor.PrimeFactor(input)
	expect := []int{2, 2, 2}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_nine_expected_array_of_double_three(t *testing.T) {
	input := 9
	result := primefactor.PrimeFactor(input)
	expect := []int{3, 3}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_eleven_expected_array_of_eleven(t *testing.T) {
	input := 11
	result := primefactor.PrimeFactor(input)
	expect := []int{11}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_twelve_expected_array_of_2_2_3(t *testing.T) {
	input := 12
	result := primefactor.PrimeFactor(input)
	expect := []int{2, 2, 3}
	sort.Ints(result)
	sort.Ints(expect)
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_fourteen_expected_array_of_2_7(t *testing.T) {
	input := 14
	result := primefactor.PrimeFactor(input)
	expect := []int{2, 7}
	sort.Ints(result)
	sort.Ints(expect)
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}

func Test_Primefactor_input_twentyfive_expected_array_of_5_5(t *testing.T) {
	input := 25
	result := primefactor.PrimeFactor(input)
	expect := []int{5, 5}
	sort.Ints(result)
	sort.Ints(expect)
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("PrimeFactor(%v) = %v should be %v", input, result, expect)
	}
}
