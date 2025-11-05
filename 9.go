package main

import (
	"errors"
	"fmt"
	"unicode"
)

func disassemble(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	runes := []rune(s)
	result := make([]rune, 0)
	i := 0

	for i < len(runes) {
		if runes[i] == '\\' {
			if i+1 >= len(runes) {
				return "", errors.New("Неверное экранирование в конце")
			}
			result = append(result, runes[i+1])
			i += 2
		} else if unicode.IsLetter(runes[i]) || runes[i] == '_' {
			result = append(result, runes[i])
			i++
		} else if unicode.IsDigit(runes[i]) {
			if len(result) == 0 {
				return "", errors.New("Неэкранированная цифра в начале")
			}
			count := int(runes[i] - '0')
			for j := 1; j < count; j++ {
				result = append(result, result[len(result)-1])
			}
			i++
		} else {
			return "", fmt.Errorf("Неизвестный символ: %c", runes[i])
		}
	}

	return string(result), nil
}

func main() {

}
