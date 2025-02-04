package strings

import (
	"testing"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		word     string
		reversed string
	}{
		{"umbrella", "ambrellu"},
		{"coat", "caot"},
		{"eventually", "avunteelly"},
		{"sooner rather than later", "seanar rethar then lotor"},
	}

	for i, test := range tests {
		got, err := ReverseVowels(test.word)
		if err != nil {
			t.Fatalf("Failed test case #%d. Unexpected Error. Error: %s", i, err)
		}
		if got != test.reversed {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.reversed, got)
		}
	}
}
