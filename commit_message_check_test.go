package main

import "testing"

type TestMessage struct {
	in  string
	out error
}

type TestBool struct {
	in  string
	out bool
}

var testsDot = []TestBool{
	{"Abc.", true},
	{"Abc", false},
}

var testsUpper = []TestBool{
	{"First char is upper", true},
	{"first char is not upper", false},
}

var testsMessage = []TestMessage{
	{"One line message.", errSubjectRequired},
	{"two line message with a long-long-long-long-long-long-long-long-long-long-long subject\n\nanother line", errSubjectTooLong},
	{"two line message\n\nanother line", errSubjectWrongCase},
	{"Two line message\n\nanother line", errSubjectRedundantDot},
	{"Two line message.\n\nanother line", errMessageWrongCase},
	{"Two line message.\n\nAnother line", errMessageWithoudDot},
	{"Two line message.\n\nAnother line.", nil},
}

func TestIsEndWithDot(t *testing.T) {
	for _, test := range testsDot {
		result := IsEndWithDot(test.in)
		if result != test.out {
			t.Errorf("%v want %v", result, test.out)
		}
	}
}

func TestIsFirstCharUpper(t *testing.T) {
	for _, test := range testsUpper {
		result := IsFirstCharUpper(test.in)
		if result != test.out {
			t.Errorf("%v want %v", result, test.out)
		}
	}
}

func TestCommitMessageCheck(t *testing.T) {
	for _, test := range testsMessage {
		result := CommitMessageCheck(test.in)
		if result != test.out {
			t.Errorf("IsFirstCharUpper(%s) == %v, want %v", test.in, result, test.out)
		}
	}
}
