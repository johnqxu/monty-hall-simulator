package main

import (
	"fmt"
	"math/rand"
)

const doorsCount int = 3

type game struct {
	doors          *[]bool
	initSelection  int
	openedDoor     int
	finalSelection int
	bonusIndex     int
}

func (g *game) initGame() {
	g.bonusIndex = rand.Intn(doorsCount)
	ds := []bool{}
	for i := 0; i < doorsCount; i++ {
		if i != g.bonusIndex {
			ds = append(ds, false)
		} else {
			ds = append(ds, true)
		}
	}
	g.doors = &ds
}

func (g *game) initalSelect() {
	g.initSelection = rand.Intn(doorsCount)
	g.finalSelection = g.initSelection
}

func (g *game) openOneDoor() {
	for i := 0; i < doorsCount; i++ {
		if i != g.initSelection && i != g.bonusIndex {
			g.openedDoor = i
		}
	}
}

func (g *game) finialSelect() {
	for i := 0; i < doorsCount; i++ {
		if i != g.initSelection && i != g.openedDoor {
			g.finalSelection = i
		}
	}
}

func (g *game) checkFinalResult() bool {
	return (*g.doors)[g.finalSelection]
}

func (g *game) String() string {
	return fmt.Sprintf("doors:%v, initSelection:%v, openedDoor:%v, finalSelection: %v, bonusIndex: %v, score:%v", *g.doors, g.initSelection, g.openedDoor, g.finalSelection, g.bonusIndex, g.checkFinalResult())
}

func main() {
	var g game
	var score int64
	rounds := 1000
	for i := 0; i < rounds; i++ {
		g.initGame()
		g.initalSelect()
		g.openOneDoor()
		g.finialSelect()
		if g.checkFinalResult() {
			score++
		}
		fmt.Println(&g)
	}
	fmt.Printf("共计%v轮，答对%v次", rounds, score)
}
