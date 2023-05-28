package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/sgade/randomorg"
)

func makeClient() *randomorg.Random {
	apiKey := os.Getenv("RANDOMORG_API_KEY")
	client := randomorg.NewRandom(apiKey)
	return client
}

func rollD6() int {
	random := makeClient()
	value, err := random.GenerateIntegers(1, 1, 6)
	//fmt.Printf("Rolled a %v\n", value)
	if err != nil {
		panic(err)
	}
	return int(value[0])
}

func rollAbilityScores() [][]int {
	allRolls := make([][]int, 6)
	for i := 0; i < 6; i++ {
		allRolls[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			allRolls[i][j] = rollD6()
		}
		//fmt.Println()
	}
	return allRolls
}

func generateAbilityScores() []int {
	abilityRolls := rollAbilityScores()
	abilityScores := make([]int, 6)
	for i := 0; i < 6; i++ {
		sort.Ints(abilityRolls[i])
		for j := 1; j < 4; j++ {
			abilityScores[i] += abilityRolls[i][j]
		}
	}
	return abilityScores
}

func main() {
	abilityScores := generateAbilityScores()
	sort.Ints(abilityScores)
	fmt.Println(abilityScores)
}
