package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

type player struct {
	health int
	damage int
	armor  int
}

func items(file string) []item {
	f, _ := os.Open(file)
	defer f.Close()
	s := bufio.NewScanner(f)
	items := []item{}
	for i := 0; s.Scan(); i++ {
		line := s.Text()
		re := regexp.MustCompile("^([a-zA-Z0-9]+)\\s+([0-9]+)\\s+([0-9]+)\\s+([0-9]+)$")
		parts := re.FindStringSubmatch(line)
		if len(parts) == 0 {
			continue
		}
		name := parts[1]
		cost, _ := strconv.Atoi(parts[2])
		damage, _ := strconv.Atoi(parts[3])
		armor, _ := strconv.Atoi(parts[4])
		items = append(items, item{name, cost, damage, armor})
	}
	return items
}

func boss(file string) player {
	f, _ := os.Open(file)
	defer f.Close()
	s := bufio.NewScanner(f)
	boss := player{}
	for s.Scan() {
		line := s.Text()
		re := regexp.MustCompile("^(Hit Points|Damage|Armor): ([0-9]+)$")
		parts := re.FindStringSubmatch(line)
		key := parts[1]
		value, _ := strconv.Atoi(parts[2])
		switch key {
		case "Hit Points":
			boss.health = value
		case "Damage":
			boss.damage = value
		case "Armor":
			boss.armor = value
		}
	}
	return boss
}

func config(w, a, r1, r2 int, items []item) []item {
	weapons := items[0:5]
	armor := items[5:10]
	rings := items[10:16]
	result := []item{weapons[w]}
	if a >= 0 {
		result = append(result, armor[a])
	}
	if r1 >= 0 {
		result = append(result, rings[r1])
	}
	if r2 >= 0 {
		result = append(result, rings[r2])
	}
	return result
}

func configs(items []item) [][]item {
	result := [][]item{}
	for w := 0; w < 5; w++ {
		for a := -1; a < 5; a++ {
			for r1 := -1; r1 < 6; r1++ {
				for r2 := -1; r2 < 6; r2++ {
					if r2 == r1 && r1 != -1 {
						continue
					}
					result = append(result, config(w, a, r1, r2, items))
				}
			}
		}
	}
	return result
}

func you(items []item) player {
	you := player{100, 0, 0}
	for _, item := range items {
		you.armor += item.armor
		you.damage += item.damage
	}
	return you
}

func costs(items []item) int {
	costs := 0
	for _, item := range items {
		costs += item.cost
	}
	return costs
}

func attack(attacker player, defender *player) {
	damage := attacker.damage - (*defender).armor
	if damage < 1 {
		damage = 1
	}
	(*defender).health -= damage
}

func main() {
	items := items("items")
	configs := configs(items)
	least := 100
	for _, items := range configs {
		costs := costs(items)
		you := you(items)
		boss := boss("input")
		for {
			attack(you, &boss)
			if boss.health <= 0 {
				if costs < least {
					least = costs
				}
				break
			}
			attack(boss, &you)
			if you.health <= 0 {
				break
			}
		}
	}
	fmt.Println(least)
}
