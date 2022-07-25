package main

import "testing"

func TestHighCardLose(t *testing.T) {
	result := ComparePokerHand("8C TS KC 9H 4S 7D 2S 5D 3S AC")
	expectedResult := "lose"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestHighCardWinWithAce(t *testing.T) {
	result := ComparePokerHand("AC TS KC 9H 4S 7D 2S 5D 3S KC")
	expectedResult := "win"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestHighCardWinWithAce2(t *testing.T) {
	result := ComparePokerHand("JC TS KC 9H AS 7D 2S 5D 3S KC")
	expectedResult := "win"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestHighCardWinWithKing(t *testing.T) {
	result := ComparePokerHand("2C TS KC 9H 4S 7D 2S 5D 3S QC")
	expectedResult := "win"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestHighCardWinWithQueen(t *testing.T) {
	result := ComparePokerHand("2C TS QC 9H 4S 7D 2S 5D 3S JC")
	expectedResult := "win"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestHighCardWinWithJack(t *testing.T) {
	result := ComparePokerHand("2C 9S 6C JH 4S 7D 2S 4D 3S 5C")
	expectedResult := "win"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestHighCardWinWithNumber(t *testing.T) {
	result := ComparePokerHand("2C 3S 6C 9H 4S 7D 2S 4D 3S 5C")
	expectedResult := "win"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestSecondHighCardWinWith10(t *testing.T) {
	result := ComparePokerHand("AC 3S 6C TH 4S AD 2S 4D 3S 5C")
	expectedResult := "win"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}

func TestSecondHighCardLoseWith10(t *testing.T) {
	result := ComparePokerHand("AC 3S 6C 9H 4S AD 2S TD 3S 5C")
	expectedResult := "lose"
	if result != expectedResult {
		t.Errorf("Got %q, want %q", result, expectedResult)
	}
}
