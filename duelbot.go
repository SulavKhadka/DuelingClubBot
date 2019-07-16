package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//PlayerList is the return payload of the API that contains all the players and thier spell deck.
type PlayerList struct {
	PlayerName string
	SpellLists SpellList
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

	//Add i random spells to the spell array
	for i := 0; i < numberOfSpells; i++ {
		spellNumber := rand.Intn(len(spells))
		chosenSpells.SpellArray = append(chosenSpells.SpellArray, spells[spellNumber])
	}
	chosenSpells.Error = ""
	return chosenSpells
}

//CreateMatch returns an array of PlayerList structs which contain every player and their spell list
func CreateMatch(w http.ResponseWriter, r *http.Request) {

	numberOfPlayers, err := strconv.Atoi(strings.Split(r.URL.String(), "/")[2]) //Parses request URL to get the number of players to generate spells
	check(err)
	numberOfSpells := 3

	//Main loop to create and populate spells for all players
	matchPayload := make([]PlayerList, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		currentPlayer := &matchPayload[i]
		currentPlayer.PlayerName = fmt.Sprintf("Player %d", i)
		currentPlayer.SpellLists = GetRandomSpells(numberOfSpells)
	}

	json.NewEncoder(w).Encode(matchPayload)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/newmatch/{numberOfPlayers}", CreateMatch).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
