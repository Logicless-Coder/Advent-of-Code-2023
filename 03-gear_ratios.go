package main

import (
	"fmt"
	"strconv"
)

type Symbol struct {
	row int
	col int
}

type Number struct {
	row  int
	col1 int
	col2 int
	num  int
}

func main() {
	n := 1

	symbols := []Symbol{}
	numbers := []Number{}

	row := 0
	for n > 0 {
		var line string
		n, _ = fmt.Scanln(&line)

		current_number := 0
		for col := 0; col < len(line); col++ {
			if '0' <= line[col] && line[col] <= '9' {
				digit := int(line[col])
				current_number = current_number*10 + digit - 48
			} else {
				if current_number > 0 {
					col_start := col - len(strconv.Itoa(current_number))
					number := Number{row: row, col1: col_start, col2: col - 1, num: current_number}
					numbers = append(numbers, number)
					current_number = 0
				}
				if line[col] != '.' {
					symbol := Symbol{row: row, col: col}
					symbols = append(symbols, symbol)
				}
			}
		}

		if current_number > 0 {
			col_start := len(line) - len(strconv.Itoa(current_number))
			number := Number{row: row, col1: col_start, col2: len(line) - 1, num: current_number}
			numbers = append(numbers, number)
			current_number = 0
		}

		row++
	}

	sum := 0
	for _, number := range numbers {
		for _, symbol := range symbols {
			if (number.row-1 <= symbol.row && symbol.row <= number.row+1) && (number.col1-1 <= symbol.col && symbol.col <= number.col2+1) {
				sum += number.num
				break
			}
		}
	}

	fmt.Println("Sum:", sum)
}
