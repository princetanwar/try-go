package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FactType struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	GetCatFact()
}

func GetCatFact() {

	var CatFact FactType

	err := ParseJson("https://catfact.ninja/fact", &CatFact)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Fact %s \n", CatFact.Fact)
}

func ParseJson(url string, target interface{}) error {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error", err)
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
