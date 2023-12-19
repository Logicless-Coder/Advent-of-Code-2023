package main

import "fmt"

func main() {
  n := 1
  sum := 0

  for n > 0 {
    var line string
    n, _ = fmt.Scanln(&line)

    first := 0
    for i := 0; i < len(line); i++ {
      if '0' <= line[i] && line[i] <= '9' {
        first = int(line[i]) - int('0')
        break
      }
    }

    last := 0
    for i := len(line) - 1; i >= 0; i-- {
      if '0' <= line[i] && line[i] <= '9' {
        last = int(line[i]) - int('0')
        break
      }
    }

    number := first * 10 + last
    sum += number
  }

    fmt.Println("Sum:", sum)
}
