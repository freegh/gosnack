package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type score struct {
	score string
	wordScore string
}

var scores = map[int]score{
	0: score{
		score: "0", wordScore: "Love",
	},
	1: score{
		score: "15", wordScore: "Fifteen",
	},
	2: score{
		score: "30", wordScore: "Thirty",
	},
	3: score{
		score: "40", wordScore: "Forty",
	},
}

func PlayerAGetPoint(players map[string]int) {
	players["A"] +=1
}

func PlayerBGetPoint(players map[string]int) {
	players["B"] +=1
}

func CurrentScore(players map[string]int, gameSet *bool, deuce *bool)string {
	scoreChart := strconv.Itoa(players["A"]) + "-" + strconv.Itoa(players["B"])
	if !*deuce && players["A"] == players["B"] {
		if players["A"] == 3 {
			*deuce = true
		}
		return (scores[players["A"]].wordScore + "-All")
	} else if *deuce && players["B"] == players["A"] {
		return ("Deuce " + scoreChart)
	} else if !*deuce && players["A"] == 4 && players["A"] > players["B"] {
		*gameSet = true
		return ("A wins " + scoreChart)
	} else if !*deuce && players["B"] == 4 && players["B"] > players["A"] {
		*gameSet = true
		return ("B wins " + scoreChart)
	} else if *deuce && players["A"] == players["B"] + 2 {
		*gameSet = true
		return ("A wins " + scoreChart)
	} else if *deuce && players["B"] == players["A"] + 2 {
		*gameSet = true
		return ("B wins " + scoreChart)
	} else if !*deuce {
		return (scores[players["A"]].wordScore + "-" + scores[players["B"]].wordScore)
	} else {
		return (scoreChart)
	}
}

func main() {
	players := make(map[string]int)
	var gameSet bool
	deuce := false
	
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("A-B")
	for !gameSet {
		if rand.Intn(2) == 0 {
			PlayerAGetPoint(players)
		} else {
			PlayerBGetPoint(players)
		}

		fmt.Println(CurrentScore(players, &gameSet, &deuce))

	}
}