package main

import (
  "fmt"
  "strings"
)

func main() {
  n := 1
  sum := 0

  for n > 0 {
    var line string
    n, _ = fmt.Scanln(&line)

    if n == 0 {
      break
    }

    first := len(line)
    for i := 0; i < len(line); i++ {
      if '0' <= line[i] && line[i] <= '9' {
        first = i
        break
      }
    }

    last := -1
    for i := len(line) - 1; i >= 0; i-- {
      if '0' <= line[i] && line[i] <= '9' {
        last = i
        break
      }
    }

    words := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    first_number, last_number := 0, 0

    first_word, first_word_i := len(line), -1
    for i := 0; i < 9; i++ {
      current_word := strings.Index(line, words[i])

      if -1 < current_word && current_word < first_word {
        first_word = current_word
        first_word_i = i
      }
    }

    if -2 < first_word && first_word < first {
      first_number = (first_word_i + 1)
    } else {
      first_number = int(line[first]) - int('0')
    }

    last_word, last_word_i := -2, -1
    for i := 0; i < 9; i++ {
      current_word := strings.LastIndex(line, words[i])
      
      if current_word > last_word {
        last_word = current_word
        last_word_i = i
      }
    }

    if last_word > last {
      last_number = (last_word_i + 1)
    } else {
      last_number = int(line[last]) - int('0')
    }

    number := first_number * 10 + last_number
    sum += number
  }

    fmt.Println("Sum:", sum)
}
