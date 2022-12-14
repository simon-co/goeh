package goeh

import (
	"errors"
	"log"
	"testing"
)

var testErr = AppErr{
	File:      "filename",
	Operation: "operation",
	Line:      10,
	Message:   "Error message",
	Cause:     errors.New("Error Cause"),
}

func TestAppErrError(t *testing.T) {
	errorStr := testErr.Error()
	log.Println(errorStr)
}

func TestAppErrIs(t *testing.T) {
	testCases := []struct {
		err      error
		expected bool
	}{
		{
			err:      errors.New("No match"),
			expected: false,
		},
		{
			err:      errors.New(testErr.Message),
			expected: false,
		},
		{
			err: &AppErr{
				Message: testErr.Message,
			},
			expected: true,
		},
	}
	for i, tc := range testCases {
		ok := errors.Is(tc.err, &testErr)
		if ok != tc.expected {
			t.Errorf("case failed:%d expected:%t received:%t", i, tc.expected, ok)
		}
	}
}
