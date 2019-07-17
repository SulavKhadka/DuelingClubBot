package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRandomSpells(t *testing.T) {
	spells := GetRandomSpells(3)

	if len(spells.SpellArray) != 3 && spells.Error == "" {
		t.Errorf("GetRandomSpells was incorrect, got: %d, want: %d.", len(spells.SpellArray), 3)
	}
}

func TestCreateMatch(t *testing.T) {

	numberOfPlayers := 3

	//API endpoint setup
	req, err := http.NewRequest("GET", "/newmatch/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMatch)

	//Making the API call
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Unmarshalling and storing API response in the proper struct
	matchDetails := make([]PlayerList, numberOfPlayers)
	err = json.Unmarshal(rr.Body.Bytes(), &matchDetails)
	if err != nil {
		t.Fatal(err)
	}

	if len(matchDetails) != 3 {
		t.Errorf("CreateMatch was incorrect, got: %d, want: %d.", len(matchDetails), 3)
	}

}
