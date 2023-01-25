package test

import (
	"testing"
)

func TestBuddyStrings(t *testing.T) {
	buddyTests := []struct {
		s              string
		goal           string
		expectedResult bool
	}{
		{"ab", "ba", true},
		{"ab", "ab", false},
		{"aa", "aa", true},
	}

	for _, e := range buddyTests {
		k := buddyStrings(e.s, e.goal)

		if k == e.expectedResult {
			t.Log("PASS")
		} else {
			t.Error("Unexpected ", k)
		}
	}
}
