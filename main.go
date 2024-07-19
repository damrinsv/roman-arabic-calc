package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		a, b, operation, err := parseInput(input)
		if err != nil {
			panic(err)
		}

		roman := map[string]int{
			"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11,
			"XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20}

		val1, ok1 := roman[a]
		val2, ok2 := roman[b]

		if ok1 && ok2 {
			result, err := calculate(val1, val2, operation)
			if err != nil {
				panic(err)
			} else if result < 1 {
				panic(err)
			} else {
				fmt.Print("Результат: ")
				_, romanResult := intToRoman(result)
				fmt.Println(romanResult)
			}
		} else {
			val1, err1 := strconv.Atoi(a)
			val2, err2 := strconv.Atoi(b)
			if err1 == nil && err2 == nil {
				if (val1 > 0 && val1 <= 20) && (val2 > 0 && val2 <= 20) {
					result, err := calculate(val1, val2, operation)
					if err != nil {
						panic(err)
					} else {
						fmt.Printf("Результат: %v\n", result)
					}
				} else {
					panic("Некорректные значения")
				}
			} else {
				panic("Некорректные значения")
			}
		}
	}
}

func parseInput(input string) (string, string, string, error) {
	operators := []string{"+", "-", "*", "/"}

	for _, op := range operators {
		if strings.Contains(input, op) {
			parts := strings.Split(input, op)
			if len(parts) != 2 {
				return "", "", "", errors.New("Некорректный ввод  данных")
			}
			return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), op, nil
		}
	}
	return "", "", "", errors.New("некорректный символ математической операции")
}

// функция конвертации арабских чисел в римские
func intToRoman(num int) (err string, result string) {
	if num < 1 || num > 3999 {
		return "Ошибка! Значение должно быть в диапазоне от 1 до 3999", ""
	}

	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; num > 0; i++ {
		for num >= vals[i] {
			num -= vals[i]
			result += syms[i]
		}
	}
	return "", result
}

// функция с математическими операциями и проверкой на корректность введенного символа
func calculate(x, y int, operation string) (int, error) {
	switch operation {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "/":
		if y == 0 {
			return 0, errors.New("ошибка! деление на ноль")
		}
		return x / y, nil
	case "*":
		return x * y, nil
	default:
		return 0, errors.New("некорректный символ математической операции")
	}
}
