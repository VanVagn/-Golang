package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var pattern string = "^[IVXLC]+$"

var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
	"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	cleandeText := strings.ReplaceAll(input, "\n", "")
	cleandeText = strings.ReplaceAll(input, "\r", "")
	cleandeText = strings.ReplaceAll(input, "\t", "")

	parts := strings.Fields(cleandeText)

	if len(parts) == 0 {
		panic("Выдача паники, так как введена пустая строка")
	} else if len(parts) == 1 {
		chars := "+-/*"
		if strings.IndexAny(parts[0], chars) != -1 {
			index := strings.IndexAny(parts[0], chars)
			str := parts[0]
			num1 := str[:index]
			oper := str[index]
			num2 := str[index+1:]
			parts[0] = num1
			parts = append(parts, string(oper))
			parts = append(parts, num2)
		} else {
			panic("Выдача паники, так как строка не является математической операцией.")
		}
	} else if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	isNumArab1, numArab1 := isArabic(parts[0])
	isNumArab2, numArab2 := isArabic(parts[2])
	isNumRom1, numRom1 := isRoman(parts[0])
	isNumRom2, numRom2 := isRoman(parts[2])

	if isNumArab1 && isNumArab2 {
		if !(numArab1 > 10 || numArab2 > 10) {
			result, err := calculation(numArab1, numArab2, parts[1])
			if err != nil {
				print(err)
			} else {
				fmt.Println(result)
			}
		} else {
			panic("Выдача паники, так как калькулятор должен принимать на вход числа от 1 до 10 включительно.")
		}
	} else if isNumRom1 && isNumRom2 && !(numRom1 == -1 || numRom2 == -1) {
		result, _ := calculation(numRom1, numRom2, parts[1])
		if result >= 0 {
			fmt.Println(arabicToRoman[result])
		} else {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
	} else if numRom1 == -1 || numRom2 == -1 {
		panic("Выдача паники, так как калькулятор должен принимать на вход числа от 1 до 10 включительно.")
	} else if isNumArab1 && isNumRom2 {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else if isNumRom1 && isNumArab2 {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда числа")
	}

}

func isArabic(num string) (bool, int) {
	number, err := strconv.Atoi(num)
	if err != nil {
		return false, -1
	}
	return true, number
}

func isRoman(num string) (bool, int) {
	matched, err := regexp.MatchString(pattern, num)
	if err != nil {
		return false, -1
	}
	if matched {
		value, ok := romanToArabic[num]
		if ok {
			return true, value
		} else {
			return true, -1
		}
	} else {
		return false, -1
	}
}

func calculation(a int, b int, oper string) (int, error) {
	switch oper {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return -1, errors.New("Деление на ноль")
		}
		return a / b, nil
	default:
		return -2, errors.New("Неверная операция")
	}
}
