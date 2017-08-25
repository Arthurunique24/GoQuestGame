package game

import (
	"fmt"
	"math/rand"
	"time"
)

type gameInit struct {
	startPos int
	endPos int
	keyPos int
	hasKey bool
	curPos int
	matrix [MAP_SIZE][MAP_SIZE] int
}

const MAP_SIZE = 11


func (gi *gameInit) createMatrix(){
	gi.matrix = [MAP_SIZE][MAP_SIZE]int {
		{0, 1, 0, 0, 0 ,0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	}
}

var gi gameInit
var sessionStarted bool = false

func Start() (bool, []int){
	fmt.Println("Welcome to quest! \nYour task is to find the key, and get out of the maze. \nGood Luck!\n")
	gi.createMatrix()
	gi.startPos = generateRandPosition(MAP_SIZE, []int {})
	gi.endPos = generateRandPosition(MAP_SIZE, []int {gi.startPos})
	gi.keyPos = generateRandPosition(MAP_SIZE, []int {gi.startPos, gi.endPos})
	gi.hasKey = false
	gi.curPos = gi.startPos
	sessionStarted = true

	fmt.Println("Key:", gi.keyPos, "Start:", gi.startPos, "End:", gi.endPos)

	return true, Answer()
}


func Update(newState int) (bool, bool, string) { // correct, finished, message
	possibleStates := Answer()
	possible := false
	for j := 0; j < len(possibleStates) && !possible; j++ {
		possible = newState == possibleStates[j]
	}
	if !possible {
		return false, false, "Incorrect turn"
	}
	gi.curPos = newState
	if gi.curPos == gi.keyPos {
		gi.hasKey = true
		return true, false, "Found key"
	}
	if gi.curPos == gi.endPos && gi.hasKey {
		sessionStarted = false
		return true, true, "Finished"
	} else {
		return true, false, "Ok"
	}
}

func Turn(newState int) (bool, bool, []int, string) {
	if !sessionStarted {
		return false, false, []int{}, "Game not started"
	}
	correct, end, message := Update(newState)
	if !correct {
		return false, gi.hasKey, Answer(), message
	}
	if end {
		return true, gi.hasKey, []int{}, message
	} else {
		return false, gi.hasKey, Answer(), message
	}
}


func Answer() ([]int){
	var states []int
	for j := 0; j < MAP_SIZE; j++ {
		if gi.matrix[gi.curPos][j] == 1 {
			fmt.Println(j)
			states = append(states, j)
		}
	}

	return states
}

func generateRandPosition(max int, exclusions []int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	placed := false
	pos := rand.Intn(max)

	if len(exclusions) == 0 {
		return pos
	}

	for !placed {
		for j := 0; j < len(exclusions) && !placed; j++ {
			pos = rand.Intn(max)
			//fmt.Println("alo")
			placed = pos != exclusions[j]
		}
	}
	return pos
}
