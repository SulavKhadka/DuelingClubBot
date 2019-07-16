package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

//PlayerList is the return payload of the API that contains all the players and thier spell deck.
type PlayerList struct {
	PlayerName string
	SpellLists []string
}

//Spells is a struct for stroing the spells
type Spells struct {
	Name   string
	Type   string
	Effect string
}

//SpellList is a list of spells.
type SpellList struct {
	SpellArray []Spells
	Error      string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//GetRandomSpells returns x amount of spells
func GetRandomSpells(numberOfSpells int) SpellList {

	//Read in all the spells into a spells slice
	spellsFile, _ := ioutil.ReadFile("./spells.json")
	var spells []Spells
	err := json.Unmarshal(spellsFile, &spells)
	check(err)

	var chosenSpells SpellList //Initialize return payload of the function

	//Check to see if number of spells exceeds the total spells in the file
	if numberOfSpells >= len(spells) {
		chosenSpells.SpellArray = nil
		chosenSpells.Error = "Invalid amount of spells. Please choose a number between 1, 100."
		return chosenSpells
	}

	rand.Seed(time.Now().Unix()) //Initialize global pseudo random generator

	//Add i random spells to the spell array
	for i := 0; i < numberOfSpells; i++ {
		spellNumber := rand.Intn(len(spells))
		chosenSpells.SpellArray = append(chosenSpells.SpellArray, spells[spellNumber])
	}
	chosenSpells.Error = ""
	return chosenSpells
}

func main() {
	spells := GetRandomSpells(3)
	for i := 0; i < len(spells.SpellArray); i++ {
		fmt.Println(spells.SpellArray[i].Name, ":", spells.SpellArray[i].Effect)
	}
}
