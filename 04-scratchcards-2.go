package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`Card [0-9 ]+: ([0-9 ]+) \| ([0-9 ]+)`)

	sum, card_number := 0, 0
  copies := make(map[int]int)
  
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
    card_number += 1
    copies[card_number] += 1

		match := re.FindAllStringSubmatch(line, -1)
    if len(match) == 0 {
      fmt.Println("THIS SHOULD NOT HAPPEN")
      continue
    }

		var winners []int
		winners_str := strings.Split(strings.Trim(match[0][1], " "), " ")
		ours_str := strings.Split(strings.Trim(match[0][2], " "), " ")

		for _, item_str := range winners_str {
      item, _ := strconv.Atoi(strings.Trim(item_str, " "))
			winners = append(winners, item)
		}

    count := 0
		for _, item_str := range ours_str {
      item, _ := strconv.Atoi(strings.Trim(item_str, " "))
      if item == 0 {
        continue
      }
      for _, x := range winners {
        if x == item {
          count += 1
          break
        }
      }
		}

    if count > 0 {
      for i := card_number + 1; i <= card_number + count; i += 1 {
        copies[i] += copies[card_number]
      }
    }
	}

  for i := 1; i <= card_number; i += 1 {
    sum += copies[i]
  }

	fmt.Println("Sum:", sum)
}
