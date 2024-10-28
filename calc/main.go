package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	rpn, err := toRPN(expression)
	if err != nil {
		return 0, err
	}

	return evalRPN(rpn)
}

func toRPN(expression string) ([]string, error) {
	var output []string
	var operators []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	isOperator := func(c string) bool {
		return c == "+" || c == "-" || c == "*" || c == "/"
	}

	isDigit := func(c string) bool {
		return c >= "0" && c <= "9"
	}

	for i := 0; i < len(expression); i++ {
		char := string(expression[i])

		if isDigit(char) || char == "." {
			number := char
			for i+1 < len(expression) && (isDigit(string(expression[i+1])) || string(expression[i+1]) == ".") {
				number += string(expression[i+1])
				i++
			}
			output = append(output, number)
		} else if char == "(" {
			operators = append(operators, char)
		} else if char == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, errors.New("неправильное выражение: несоответствие скобок")
			}
			operators = operators[:len(operators)-1] 
		} else if isOperator(char) {
			for len(operators) > 0 && operators[len(operators)-1] != "(" &&
				precedence[char] <= precedence[operators[len(operators)-1]] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, char)
		} else {
			return nil, fmt.Errorf("неизвестный символ в выражении: %s", char)
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, errors.New("неправильное выражение: несоответствие скобок")
		}
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output, nil
}

func evalRPN(rpn []string) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("неправильное выражение")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("деление на ноль")
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("неизвестный оператор: %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("неправильное выражение")
	}

	return stack[0], nil
}

func main() {
	expr := "18+(12-10)*2"
	result, err := Calc(expr)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
