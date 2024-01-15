package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func update_seed_to_stuff(seed_to_stuff *map[int]int, temp_map map[int][]int) {
	for seed := range *seed_to_stuff {
    stuff := (*seed_to_stuff)[seed]
		just_lower := -1
		for key := range temp_map {
			if key <= stuff && key > just_lower && key+temp_map[key][1] >= stuff {
				just_lower = key
			}
		}

		if just_lower != -1 {
			diff := stuff - just_lower
			(*seed_to_stuff)[seed] = temp_map[just_lower][0] + diff
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	first_line := scanner.Text()
	seeds_str := strings.Split(first_line, ": ")[1]

	var seeds []int
	for _, item_str := range strings.Split(seeds_str, " ") {
		seed, _ := strconv.Atoi(strings.Trim(item_str, " "))
		seeds = append(seeds, seed)
	}

	seed_to_stuff := make(map[int]int)
	temp_map := make(map[int][]int)
  for _, seed := range seeds {
    seed_to_stuff[seed] = seed
  }

	map_start := false
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasSuffix(line, " map:") {
			map_start = true
		} else if line == "" {
			map_start = false
		}

		if map_start && !strings.HasSuffix(line, " map:") {
			nums_str := strings.Split(line, " ")
			dest_start, _ := strconv.Atoi(nums_str[0])
			src_start, _ := strconv.Atoi(nums_str[1])
			length, _ := strconv.Atoi(nums_str[2])

			temp_map[src_start] = []int{dest_start, length}
		} else if !map_start {
			update_seed_to_stuff(&seed_to_stuff, temp_map)
			for k := range temp_map {
				delete(temp_map, k)
			}
		}
	}
	update_seed_to_stuff(&seed_to_stuff, temp_map)

	lowest_location := math.MaxInt
	for k := range seed_to_stuff {
		lowest_location = min(lowest_location, seed_to_stuff[k])
	}

	fmt.Println(lowest_location)
}
