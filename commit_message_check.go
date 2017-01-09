package main

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

var (
	subjectTaskRegexp = regexp.MustCompile(`^([0-9A-Z\-\_]+)\s`)

	errSubjectRequired     = errors.New("Сообщение должно содержать заголовок, отделённый пустой строкой от содержимого")
	errSubjectTaskRequired = errors.New("Заголовок должен содержать номер задачи")
	errSubjectTooLong      = errors.New("Заголовок не должен превышать 50 символов")
	errSubjectWrongCase    = errors.New("Заголовок должен начинаться с большой буквы")
	errSubjectRedundantDot = errors.New("Заголовок не должен заканчиваться точкой")
	errMessageWrongCase    = errors.New("Содержимое должно начинаться с большой буквы")
	errMessageWithoudDot   = errors.New("Содержимое должно заканчиваться точкой")
)

// IsEndWithDot check
func IsEndWithDot(str string) bool {
	return str[len(str)-1:] == "."
}

// IsFirstCharUpper check
func IsFirstCharUpper(str string) bool {
	return strings.IndexFunc(str[0:1], unicode.IsUpper) == 0
}

// IsSubjectWithTask check
func IsSubjectWithTask(str string) bool {
	if s := subjectTaskRegexp.FindString(str); s != "" {
		return true
	}
	return false
}

// CommitMessageCheck text
func CommitMessageCheck(text string) error {
	data := strings.SplitN(text, "\n\n", 2)

	if len(data) != 2 {
		return errSubjectRequired
	}

	subject := strings.TrimSpace(data[0])
	if subject != "" {
		if !IsSubjectWithTask(subject) {
			return errSubjectTaskRequired
		}
		subject = subjectTaskRegexp.ReplaceAllString(subject, "")
		if len(subject) > 50 {
			return errSubjectTooLong
		} else if !IsFirstCharUpper(subject) {
			return errSubjectWrongCase
		} else if IsEndWithDot(subject) {
			return errSubjectRedundantDot
		}
	}

	message := strings.TrimSpace(data[1])
	if message != "" {
		if !IsFirstCharUpper(message) {
			return errMessageWrongCase
		} else if !IsEndWithDot(message) {
			return errMessageWithoudDot
		}
	}

	return nil
}
