package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type spell struct {
	cost     int
	health   int
	mana     int
	damage   int
	armor    int
	duration int
}

type player struct {
	health int
	damage int
	armor  int
	mana   int
	spells map[string]spell
}

func read(file string) player {
	player := player{}
	f, _ := os.Open(file)
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		re := regexp.MustCompile("^(Hit Points|Damage): ([0-9]+)$")
		parts := re.FindStringSubmatch(line)
		key := parts[1]
		value, _ := strconv.Atoi(parts[2])
		switch key {
		case "Hit Points":
			player.health = value
		case "Damage":
			player.damage = value
		}
	}
	return player
}

func spells() map[string]spell {
	return map[string]spell{
		"Magic Missile": spell{53, 0, 0, 4, 0, -1},
		"Drain":         spell{73, 2, 0, 2, 0, -1},
		"Shield":        spell{113, 0, 0, 0, 7, 6},
		"Poison":        spell{173, 0, 0, 3, 0, 6},
		"Recharge":      spell{229, 0, 101, 0, 0, 5},
	}
}

func you(health, mana int) player {
	you := player{}
	you.health = health
	you.mana = mana
	you.spells = map[string]spell{}
	return you
}

func attack(player, defender player) (player, player, bool) {
	defender, player = timer(defender, player)
	if player.health <= 0 {
		return player, defender, true
	}
	damage := player.damage - defender.armor
	if damage < 1 {
		damage = 1
	}
	defender.health -= damage
	return player, defender, defender.health <= 0
}

func cast(name string, spell spell, player, defender player) (player, player, bool) {
	player, defender = timer(player, defender)
	if defender.health <= 0 {
		return player, defender, true
	}
	if _, exists := player.spells[name]; exists {
		return player, defender, true
	}
	if player.mana < spell.cost {
		return player, defender, true
	}
	player.mana -= spell.cost
	if spell.duration == -1 {
		player.health += spell.health
		defender.health -= spell.damage
	} else {
		player.spells[name] = spell
	}
	return player, defender, defender.health <= 0
}

func timer(player, defender player) (player, player) {
	spells := map[string]spell{}
	player.armor = 0
	for name, spell := range player.spells {
		player.health += spell.health
		player.mana += spell.mana
		defender.health -= spell.damage
		player.armor += spell.armor
		spell.duration--
		if spell.duration > 0 {
			spells[name] = spell
		}
	}
	player.spells = spells
	return player, defender
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func play(mana int, you, boss player, spells map[string]spell, max, depth int) int {
	result := 10000
	if depth == 0 || mana > max {
		return result
	}
	for name, spell := range spells {
		nyou, nboss, ended := cast(name, spell, you, boss)
		if ended {
			if nboss.health <= 0 {
				result = min(result, mana+spell.cost)
			}
			continue
		}
		nboss, nyou, ended = attack(nboss, nyou)
		if ended {
			if nboss.health <= 0 {
				result = min(result, mana+spell.cost)
			}
			continue
		}
		result = min(result, play(mana+spell.cost, nyou, nboss, spells, max, depth-1))
	}
	return result
}

func main() {
	min := 10000
	for i := 1; i < 20; i++ {
		spells := spells()
		you := you(50, 500)
		boss := read("input")
		min = play(0, you, boss, spells, min, i)
	}
	fmt.Println(min)
}
