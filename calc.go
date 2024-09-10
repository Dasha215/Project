package main

import ( // импорт пакетов
	"bufio"   // для чтения ввода
	"fmt"     // для форматированного вывода
	"os"      // для чтения ввода, работа с ос
	"strconv" // для преобразования строк в числа
	"strings" // для работы со строками
)

var roman = map[string]int{ // Переменная roman содержит отображение римских чисел в их арабские эквиваленты.
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convIntToRoman = [14]int{ // это массив, используемый для преобразования арабских чисел в римские
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var operators = map[string]func(int, int) int{ // содержит отображение операторов на соответствующие функции для выполнения арифметических операций.
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"/": func(a, b int) int { return a / b },
	"*": func(a, b int) int { return a * b },
}

const ( // Константы используются для сообщений об ошибках, которые могут возникнуть в программе.
	LOW   = "Ошибка: строка не является математической операцией."
	HIGH  = "Ошибка: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Ошибка: используются одновременно разные системы счисления."
	DIV   = "Ошибка: в римской системе нет отрицательных чисел."
	ZERO  = "Ошибка: в римской системе нет числа 0."
	RANGE = "Ошибка: калькулятор работает только с целыми числами от 1 до 10 включительно."
)

func main() { // Функция main запускает цикл

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите выражение (или 'выход' для завершения): ")
		console, _ := reader.ReadString('\n')
		s := strings.TrimSpace(console)
		if strings.ToLower(s) == "выход" {
			fmt.Println("Завершение работы калькулятора.")
			break
		}
		result := calculate(s) // для вычисления результата и выводит его
		fmt.Println("Результат:", result)
	}
}

func calculate(input string) string { // функция вычисления
	input = strings.ReplaceAll(input, " ", "")
	operator, data := findOperator(input)
	if operator == "" {
		panic(LOW)
	}

	if len(data) != 2 {
		panic(HIGH)
	}

	num1, num2, isRoman := parseOperands(data)

	if val, ok := operators[operator]; ok {
		result := val(num1, num2)
		if isRoman {
			resultStr, err := intToRoman(result)
			if err != nil {
				panic(err.Error())
			}
			return resultStr
		}
		return strconv.Itoa(result)
	}
	panic(HIGH)
}

func findOperator(s string) (string, []string) { // ищет оператор в строке и разбивает строку на операнды.
	for op := range operators {
		if strings.Contains(s, op) {
			return op, strings.Split(s, op)
		}
	}
	return "", nil
}

func parseOperands(data []string) (int, int, bool) { // пробует преобразовать строики в числа,проверяет корректность
	num1, err1 := strconv.Atoi(data[0])
	num2, err2 := strconv.Atoi(data[1])

	if err1 == nil && err2 == nil {
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			panic(RANGE)
		}
		return num1, num2, false
	}

	val1, ok1 := roman[data[0]]
	val2, ok2 := roman[data[1]]
	if ok1 && ok2 {
		if val1 < 1 || val1 > 10 || val2 < 1 || val2 > 10 {
			panic(RANGE)
		}
		return val1, val2, true
	}

	if (ok1 && err2 == nil) || (ok2 && err1 == nil) {
		panic(SCALE)
	}

	panic(HIGH)
}

func intToRoman(number int) (string, error) { // функция преобразует число в римское представление
	if number == 0 {
		return "", fmt.Errorf(ZERO)
	} else if number < 0 {
		return "", fmt.Errorf(DIV)
	}

	var result strings.Builder
	for _, value := range convIntToRoman {
		for number >= value {
			for key, val := range roman {
				if val == value {
					result.WriteString(key)
					number -= val
				}
			}
		}
	}
	return result.String(), nil
}
