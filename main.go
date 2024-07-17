package roman_arabic_calc

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var a, b, operation string

	for {
		fmt.Print("> ")
		fmt.Scanln(&a, &operation, &b)

		//проверка ввода трех значений
		if a == "" || b == "" || operation == "" {
			fmt.Println("Некорректный ввод")
			break
		}
		//маппинг арабских и римских цифр
		roman := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

		val_1, ok_1 := roman[a]
		val_2, ok_2 := roman[b]
		if ok_1 && ok_2 {
			result, err := calculating(val_1, val_2, operation)
			if err != nil {
				fmt.Println(err)
				break

			} else if result < 1 {
				fmt.Println("Ошибка! Значение не положительное")
				break
			} else {
				fmt.Print("Результат:")
				fmt.Println(intToRoman(result))
			}

		} else {
			val_1, err := (strconv.Atoi(a))
			val_2, err := (strconv.Atoi(b))
			if err == nil {
				if (val_1 > 0 && val_1 <= 10) && (val_2 > 0 && val_2 <= 10) {
					result, err := calculating(val_1, val_2, operation)
					if err != nil {
						fmt.Println(err)
						break
					} else {
						fmt.Printf("Результат: %v\n", result)
					}
				} else {
					fmt.Println("Некорректные значения, попробуйте еще раз")
					break
				}
			} else {
				fmt.Println("Некорректные значения, попробуйте еще раз")
				break
			}
		}
	}
}

// функция конвертации арабских чисел в римские
func intToRoman(num int) (err, result string) {
	r := [][]string{
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	n := []int{100, 10, 1}
	for k, v := range n {
		result += r[k][num/v]
		num = num % v
	}
	return
}

// функция с математическими операциями и проверкой на корректность введенного символа
func calculating(x, y int, operation string) (int, error) {

	switch operation {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "/":
		return x / y, nil
	case "*":
		return x * y, nil
	default:
		return 0, errors.New("Некорректный символ математической операции")
	}
}
