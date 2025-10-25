package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	PokeInfo struct {
		BaseExp int        `json:"base_experience"`
		Height  int        `json:"height"`
		Weight  int        `json:"weight"`
		Types   []TypeInfo `json:"types"`
		Stats   []StatInfo `json:"stats"`
	}
	TypeInfo struct {
		Type TName `json:"type"`
	}
	TName struct {
		Name string `json:"name"`
	}
	StatInfo struct {
		BaseStat int   `json:"base_stat"`
		Stat     SName `json:"stat"`
	}
	SName struct {
		Name string `json:"name"`
	}
)

func GetPokeInfo(url string) (PokeInfo, error) {
	var pokeInfo PokeInfo
	res, err := http.Get(url)
	if err != nil {
		log.Println("Error inside Get request attempt")
		return pokeInfo, err
	}

	if res.StatusCode > 299 {
		fmt.Println("That Pokemon doesn't exist!")
		return pokeInfo, nil
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&pokeInfo); err != nil {
		return pokeInfo, err
	}

	return pokeInfo, nil
}
