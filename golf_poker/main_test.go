package main

import (
	"testing"
)

func TestCard_Input_KH_ShouldBe_13_H(t *testing.T) {
	input := "KH"
	expect1, expect2 := 13, "H"

	assert1, assert2 := Card(input)

	if expect1 != assert1 {
		t.Errorf("expect1 %v but got %v .", expect1, assert1)
	}

	if expect2 != assert2 {
		t.Errorf("expect2 %v but got %v .", expect2, assert2)
	}
}

func TestCard_Input_TS_ShouldBe_10_S(t *testing.T) {
	input := "TS"
	expect1, expect2 := 10, "S"

	assert1, assert2 := Card(input)

	if expect1 != assert1 {
		t.Errorf("expect1 %v but got %v .", expect1, assert1)
	}

	if expect2 != assert2 {
		t.Errorf("expect2 %v but got %v .", expect2, assert2)
	}
}

func TestCard_Input_9C_ShouldBe_9_C(t *testing.T) {
	input := "9C"
	expect1, expect2 := 9, "C"

	assert1, assert2 := Card(input)

	if expect1 != assert1 {
		t.Errorf("expect1 %v but got %v .", expect1, assert1)
	}

	if expect2 != assert2 {
		t.Errorf("expect2 %v but got %v .", expect2, assert2)
	}
}

func TestCard_Input_5D_ShouldBe_5_D(t *testing.T) {
	input := "5D"
	expect1, expect2 := 5, "D"

	assert1, assert2 := Card(input)

	if expect1 != assert1 {
		t.Errorf("expect1 %v but got %v .", expect1, assert1)
	}

	if expect2 != assert2 {
		t.Errorf("expect2 %v but got %v .", expect2, assert2)
	}
}

func TestHandRank_Input_8C_TS_KC_9H_4S_Return_high_card(t *testing.T) {
	input := []string{"8C", "TS", "KC", "9H", "4S"}
	expect := "high_card"

	assert := HandRank(input)

	if expect != assert {
		t.Errorf("expect %v but got %v .", expect, assert)
	}
}

func TestHandRank_Input_AH_AC_KD_JS_7H_Return_pair_card(t *testing.T) {
	input := []string{"8C", "TS", "KC", "9H", "4S"}
	expect := "high_card"

	assert := HandRank(input)

	if expect != assert {
		t.Errorf("expect %v but got %v .", expect, assert)
	}
}

func TestWinner_Input_8C_TS_KC_9H_4S_7D_2S_5D_3S_AC_Return_p1(t *testing.T) {
	input := []string{"8C", "TS", "KC", "9H", "4S", "7D", "2S", "5D", "3S", "AC"}
	expect := "p1"

	assert := Winner(input)

	if expect != assert {
		t.Errorf("expect %v but got %v .", expect, assert)
	}
}

func TestWinner_Input_3H_7H_6S_KC_JS_QH_TD_JC_2D_8S_Return_p2(t *testing.T) {
	input := []string{"3H", "7H", "6S", "KC", "JS", "QH", "TD", "JC", "2D", "8S"}
	expect := "p2"

	assert := Winner(input)

	if expect != assert {
		t.Errorf("expect %v but got %v .", expect, assert)
	}
}

func TestWinner_Input_TD_8C_4H_7C_TC_KC_4C_3H_7S_KS_Return_p2(t *testing.T) {
	input := []string{"TD", "8C", "4H", "7C", "TC", "KC", "4C", "3H", "7S", "KS"}
	expect := "p2"

	assert := Winner(input)

	if expect != assert {
		t.Errorf("expect %v but got %v .", expect, assert)
	}
}
