package string_util

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	. "github.com/gobeam/stringy"
)

func ToSnakeCase(str string, cleanText bool) string {
	if cleanText {
		str = RemoveSpecialCharacter(str, " ")
	}

	stringify := New(str)
	return stringify.SnakeCase("?", "").ToLower()
}

func ToKebabCase(str string, cleanText bool) string {
	if cleanText {
		str = RemoveSpecialCharacter(str, " ")
	}

	stringify := New(str)
	return stringify.KebabCase("?", "").ToLower()
}

func RemoveSpecialCharacter(str, replaceWith string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, replaceWith)
}

func PadLeftString(maxLength int, text, padChar string) string {
	if len(text) >= maxLength {
		return text
	}

	for insertedLength := maxLength - len(text); insertedLength > 0; insertedLength-- {
		text = padChar + text
	}
	return text
}

func PadRightString(maxLength int, text, padChar string) string {
	if len(text) >= maxLength {
		return text
	}

	for insertedLength := maxLength - len(text); insertedLength > 0; insertedLength-- {
		text += padChar
	}
	return text
}

func IntArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func ContainsArray(text string, array []string) bool {
	for _, item := range array {
		if strings.Contains(text, item) {
			return true
		}
	}
	return false
}
