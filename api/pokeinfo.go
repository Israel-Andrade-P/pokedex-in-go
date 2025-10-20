package api

import (
	"encoding/json"
	"net/http"
)

type (
	PokeExp struct {
		BaseExp int `json:"base_experience"`
	}
)

func GetPokeInfo(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	var exp PokeExp
	if err = json.NewDecoder(res.Body).Decode(&exp); err != nil {
		return 0, err
	}

	return exp.BaseExp, nil
}
