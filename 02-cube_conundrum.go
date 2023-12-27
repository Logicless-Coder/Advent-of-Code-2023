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
		game_id, _ := strconv.Atoi(strings.Split(first_split[0], " ")[1])
		hands := strings.Split(strings.TrimSpace(first_split[1]), "; ")

		good := true
		for _, hand := range hands {
			sets := strings.Split(hand, ", ")

			for _, set := range sets {
				split := strings.Split(set, " ")

				count_str, color := split[0], split[1]
				count, _ := strconv.Atoi(count_str)

				if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
					good = false
					goto skip
				}
			}
		}

	skip:
		if good {
			sum += game_id
		}
	}

	fmt.Println(sum)
}
