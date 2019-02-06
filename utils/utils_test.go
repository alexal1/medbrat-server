package utils

import "testing"

func TestGetFirstNWords(t *testing.T) {
	sample := "Hey you, out there in the cold\nGetting lonely, getting old\nCan you feel me?"
	var firstN string
	var trueFirstN string

	firstN = GetFirstNWords(sample, 3)
	trueFirstN = "Hey you, out"
	if firstN != trueFirstN {
		t.Fatalf("Wrong output: %s instead of: %s", firstN, trueFirstN)
	}

	firstN = GetFirstNWords(sample, 100000)
	trueFirstN = "Hey you, out there in the cold Getting lonely, getting old Can you feel me?"
	if firstN != trueFirstN {
		t.Fatalf("Wrong output: %s instead of: %s", firstN, trueFirstN)
	}

	firstN = GetFirstNWords(sample, 0)
	trueFirstN = ""
	if firstN != trueFirstN {
		t.Fatalf("Wrong output: %s instead of: %s", firstN, trueFirstN)
	}

	sample = "Гемоглобин \t143 \tг/л \t132 173 \t\r"
	firstN = GetFirstNWords(sample, 1)
	trueFirstN = "Гемоглобин"
	if firstN != trueFirstN {
		t.Fatalf("Wrong output: %s instead of: %s", firstN, trueFirstN)
	}
}
