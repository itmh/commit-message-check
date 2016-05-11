package main

import "testing"

type Test struct {
	in  string
	out error
}

var tests = []Test{
	{"One line message.", errSubjectRequired},
	{"two line message with a long-long-long-long-long-long-long-long-long-long-long subject\n\nanother line", errSubjectTooLong},
	{"two line message\n\nanother line", errSubjectWrongCase},
	{"Two line message\n\nanother line", errSubjectRedundantDot},
	{"Two line message.\n\nanother line", errMessageWrongCase},
	{"Two line message.\n\nAnother line", errMessageWithoudDot},
	{"Two line message.\n\nAnother line.", nil},
}

func TestIsEndWithDot(t *testing.T) {
	str := "Abc."
	result := IsEndWithDot(str)
	if result != true {
		t.Errorf("IsEndWithDot(%s) == %v, want %v", str, result, false)
	}

	str = "Abc"
	result = IsEndWithDot("Abc")
	if result != false {
		t.Errorf("IsEndWithDot(%s) == %v, want %v", str, result, false)
	}
}

func TestIsFirstCharUpper(t *testing.T) {
	str := "First char is upper"
	result := IsFirstCharUpper(str)
	if result != true {
		t.Errorf("IsFirstCharUpper(%s) == %v, want %v", str, result, false)
	}

	str = "first char is not upper"
	result = IsFirstCharUpper(str)
	if result != false {
		t.Errorf("IsFirstCharUpper(%s) == %v, want %v", str, result, false)
	}
}

func TestCommitMessageCheck(t *testing.T) {
	for _, test := range tests {
		result := CommitMessageCheck(test.in)
		if result != test.out {
			t.Errorf("IsFirstCharUpper(%s) == %v, want %v", test.in, result, test.out)
		}
	}
}
