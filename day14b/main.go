package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type reindeer struct {
	speed    int
	flies    int
	flown    int
	sleep    int
	slept    int
	traveled int
	points   int
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	s := bufio.NewScanner(f)
	deers := map[string]reindeer{}
	for s.Scan() {
		line := s.Text()
		re := regexp.MustCompile("^([a-zA-Z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.$")
		parts := re.FindStringSubmatch(line)
		name := parts[1]
		speed, _ := strconv.Atoi(parts[2])
		flies, _ := strconv.Atoi(parts[3])
		sleep, _ := strconv.Atoi(parts[4])
		deers[name] = reindeer{speed, flies, 0, sleep, 0, 0, 0}
	}
	for i := 0; i < 2503; i++ {
		for k := range deers {
			deer := deers[k]
			if deer.flown < deer.flies {
				deer.flown++
				deer.traveled += deer.speed
			} else {
				if deer.slept < deer.sleep {
					deer.slept++
				}
				if deer.slept == deer.sleep {
					deer.slept = 0
					deer.flown = 0
				}
			}
			deers[k] = deer
		}
		farrest := 0
		for k := range deers {
			if deers[k].traveled > farrest {
				farrest = deers[k].traveled
			}
		}
		for k := range deers {
			if deers[k].traveled == farrest {
				deer := deers[k]
				deer.points++
				deers[k] = deer
			}
		}
	}
	highest := 0
	for _, deer := range deers {
		if deer.points > highest {
			highest = deer.points
		}
	}
	fmt.Println(highest)
}
