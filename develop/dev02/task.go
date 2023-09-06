package dev02

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func Unpacking(str string) (string, error) {
	var builder strings.Builder
	reader := strings.NewReader(str)
	prevChar, _, _ := reader.ReadRune()

	if unicode.IsDigit(prevChar) {
		return "", fmt.Errorf("некорректная строка")
	}
	for {
		currChar, _, readErr := reader.ReadRune()
		if readErr != nil {
			builder.WriteRune(prevChar)
			break
		}

		digit, err := strconv.Atoi(string(currChar))
		if err == nil {
			builder.WriteString(strings.Repeat(string(prevChar), digit))
		} else {
			builder.WriteRune(prevChar)
		}

		if currChar == '\\' || err == nil {
			prevChar, _, readErr = reader.ReadRune()
			if readErr != nil {
				break
			}
		} else {
			prevChar = currChar
		}
	}
	return builder.String(), nil
}
