package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var line string
	var err error
	sum := 0
	for true {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		first_split := strings.Split(line, ": ")
		hands := strings.Split(strings.TrimSpace(first_split[1]), "; ")

		red, green, blue := 0, 0, 0
		for _, hand := range hands {
			sets := strings.Split(hand, ", ")

			for _, set := range sets {
				split := strings.Split(set, " ")

				count_str, color := split[0], split[1]
				count, _ := strconv.Atoi(count_str)

				if color == "red" && count > red {
					red = count
				}
				if color == "green" && count > green {
					green = count
				}
				if color == "blue" && count > blue {
					blue = count
				}
			}
		}

		sum += red * green * blue
	}

	fmt.Println(sum)
}
