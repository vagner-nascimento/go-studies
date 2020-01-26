package table

import (
	"strings"
	"testing"
)

const errorMessage = "%s (piece: %s) - expected index (%d) is different to the found index (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	tests := []struct { // Using that you can test a lot of scenarios using only this struct and a for
		text     string
		piece    string
		expected int
	}{
		{"Cod3r is good", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Leonardo", "o", 2},
	}

	for _, test := range tests {
		t.Logf("Massa: %v", test)
		current := strings.Index(test.text, test.piece)

		if current != test.expected {
			t.Errorf(errorMessage, test.text, test.piece, test.expected, current)
		}
	}
}
