package main

import (
	"errors"
	"strings"
	"unicode"
)

var (
	errSubjectRequired     = errors.New("Сообщение должно содержать заголовок, отделённый пустой строкой от содержимого")
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

// CommitMessageCheck text
func CommitMessageCheck(text string) error {
	data := strings.SplitN(text, "\n\n", 2)

	if len(data) != 2 {
		return errSubjectRequired
	}

	subject := strings.TrimSpace(data[0])
	if subject != "" {
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
