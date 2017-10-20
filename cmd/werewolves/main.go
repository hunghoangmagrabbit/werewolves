package main

import (
	"github.com/GSamuel/werewolvesmillershollow/game"
	"github.com/GSamuel/werewolvesmillershollow/roles"
)

func main() {

	were := roles.WEREWOLF
	vil := roles.VILLAGER
	hun := roles.HUNTER

	deck := game.NewDeck()
	deck.Add(vil)
	deck.Add(vil)
	deck.Add(vil)
	deck.Add(vil)
	deck.Add(hun)
	deck.Add(vil)
	deck.Add(were)
	deck.Add(were)
	deck.Shuffle()

	g := game.New(deck)
	g.Run()
}