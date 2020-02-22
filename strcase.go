// SPDX-License-Identifier: BSD-3-Clause

package strcase

import (
	"strings"

	"github.com/martingallagher/runes"
)

const (
	hyphen     = 0x2d // -
	underscore = 0x5f // _
)

// convert converts a s to (lower)CamelCase.
func convert(s string, needCap bool) (str string) {
	var sb strings.Builder

	nextCap := needCap
	for _, r := range strings.TrimSpace(s) {
		switch {
		case runes.IsLowerUnsafe(r):
			if nextCap {
				sb.WriteByte(byte(runes.ToUpperUnsafe(r)))
				break
			}
			sb.WriteByte(byte(r))
		case runes.IsUpperUnsafe(r):
			sb.WriteByte(byte(r))
		case runes.IsDigitUnsafe(r):
			sb.WriteByte(byte(r))
			nextCap = true
			continue
		}

		if runes.IsSpaceUnsafe(r) || runes.IsPunctUnsafe(r) {
			nextCap = true
			continue
		}
		nextCap = false
	}

	str = sb.String()
	return
}

// ToCamelCase converts a s to CamelCase.
func ToCamelCase(s string) (str string) {
	if s == "" {
		return s
	}

	str = convert(s, true)
	return
}

// ToLowerCamelCase converts a s to lower CamelCase.
func ToLowerCamelCase(s string) (str string) {
	if s == "" {
		return s
	}

	if r := rune(s[0]); runes.IsUpperUnsafe(r) {
		str = convert(string(runes.ToLowerUnsafe(r))+s[1:], false)
		return
	}

	str = convert(s, false)
	return
}

// ToSnakeCase converts a s to snake_case.
func ToSnakeCase(s string) (str string) {
	str = ToDelimited(s, underscore)
	return
}

// ToScreamingSnakeCase converts a s to SCREAMING_SNAKE_CASE.
func ToScreamingSnakeCase(s string) (str string) {
	str = ToScreamingDelimited(strings.TrimSpace(s), underscore, true)
	return
}

// ToKebab converts a s to kebab-case.
func ToKebab(s string) (str string) {
	str = ToDelimited(s, hyphen)
	return
}

// ToScreamingKebab converts a s to SCREAMING-KEBAB-CASE.
func ToScreamingKebab(s string) (str string) {
	str = ToScreamingDelimited(strings.TrimSpace(s), hyphen, true)
	return
}

// ToDelimited converts a s to delimited.snake.case (in this case `delim = '.'`).
func ToDelimited(s string, delim uint8) (str string) {
	str = ToScreamingDelimited(strings.TrimSpace(s), delim, false)
	return
}

// ToScreamingDelimited converts a s to any delim and screaming.
//
//  SCREAMING.DELIMITED.SNAKE.CASE
// in this case delim = '.', screaming = true
//
//  delimited.snake.case
// in this case delim = '.', screaming = false.
func ToScreamingDelimited(s string, delim uint8, screaming bool) (str string) {
	var sb strings.Builder

	isDelim := true
	for i, r := range s {
		changeNext := false

		if i < len(s)-1 {
			next := rune(s[i+1])
			if (runes.IsUpperUnsafe(r) && runes.IsLowerUnsafe(next)) || (runes.IsLowerUnsafe(r) && runes.IsUpperUnsafe(next)) {
				changeNext = true
			}
		}

		switch {
		case i > 0 && isDelim && changeNext:
			switch {
			case runes.IsLowerUnsafe(r):
				sb.WriteByte(byte(r))
				sb.WriteByte(delim)
				isDelim = false
			case runes.IsUpperUnsafe(r):
				sb.WriteByte(delim)
				sb.WriteByte(byte(r))
				isDelim = true
			}

		case runes.IsSpaceUnsafe(r), runes.IsPunctUnsafe(r):
			sb.WriteByte(delim)
			isDelim = false

		default:
			sb.WriteByte(byte(r))
			isDelim = true
		}
	}

	ss := sb.String()
	sb.Reset()

	if screaming {
		for _, r := range ss {
			if runes.IsLowerUnsafe(r) {
				sb.WriteByte(byte(runes.ToUpperUnsafe(r)))
				continue
			}
			sb.WriteByte(byte(r))
		}
		str = sb.String()
		return
	}

	for _, r := range ss {
		if runes.IsUpperUnsafe(r) {
			sb.WriteByte(byte(runes.ToLowerUnsafe(r)))
			continue
		}
		sb.WriteByte(byte(r))
	}

	str = sb.String()
	return str
}
