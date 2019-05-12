package main

import "github.com/jhungerford/auto-tank/tank"

func main() {
	var t = tank.Init()
	t.Move("Forward")
}