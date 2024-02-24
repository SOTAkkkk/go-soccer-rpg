package main

import (
	"fmt"
)

type Player struct {
	name         string
	age          int
	position     string
	abilityValue int
}

func (p *Player) setName(name string) {
	p.name = name
}

func (p *Player) setAge(age int) {
	p.age = age
}

func (p *Player) setPosition(position string) {
	p.position = position
}

func (p *Player) setAbilityValue(abilityValue int) {
	p.abilityValue = abilityValue
}

func main() {
	var player Player

	player.setName("SOTAkkkk")
	player.setAge(27)
	player.setPosition("FW")
	player.setAbilityValue(99)

	fmt.Println(player)

}
