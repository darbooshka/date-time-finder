package datetimefinder

import (
	"strings"

	"testing"
	"github.com/stretchr/testify/assert"
)

type testCase[expected any] struct {
	text     string
	expected []expected
	ignore   *string
}

func reason(s string) *string { return &s }

func getTestName(text string, length ...int) string {
	maxLength := 85
	if len(length) > 0 {
		maxLength = length[0]
	}

	// Cut the string to the specified length
	if len(text) > maxLength {
		text = text[:maxLength]
	}

	// Replace slashes with spaces
	text = strings.ReplaceAll(text, "/", " ")

	return text
}

func TestDateTimeFinder_FindDateTime(t *testing.T) {
	for _, tc := range testCases {
		t.Run(getTestName(tc.text), func(t *testing.T) {
			if tc.ignore != nil {
				t.Logf("Ignoring test case: %s", *tc.ignore)
				t.Skip(*tc.ignore)
				return // continue
			}

			finder := NewDateTimeFinder()
			result := finder.FindDateTime(tc.text)

			assert.Equal(t, tc.expected, result)

			// if len(result) != len(tc.expected) {
			// 	t.Errorf("Expected %d date-time mention(s), but got %d", len(tc.expected), len(result))
			// }

			// for i, match := range result {
			// 	if i >= len(tc.expected) {
			// 		break
			// 	}
			// 	if match != tc.expected[i] {
			// 		t.Errorf("Expected '%s', but got '%s'", tc.expected[i], match)
			// 	}
			// }
		})
	}
}
