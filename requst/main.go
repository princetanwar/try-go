package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	// GetRequest()
	// EncodeJson()
	// PostRequest()
	DecodeJson()
}

func GetRequest() {
	res, err := http.Get("https://lco.dev")

	if err != nil {
		panic(err)

	}

	dataBytes, err := ioutil.ReadAll(res.Body)
	// fmt.Println("response is ", string(dataBytes))

	var responseString strings.Builder
	responseString.Write(dataBytes)
	fmt.Println("response langth is ", responseString.Len())
	fmt.Println("rsponse is ", responseString.String())

	defer res.Body.Close()
}

func PostRequest() {
	url := "https://jsonplaceholder.typicode.com/posts"

	type dataType struct {
		Title  string `json:"title"` // this will chagne to key name when convert to JSON
		Body   string
		UserId int
	}

	var todo = dataType{"foo", "bar", 1}

	jsonByte, er := json.Marshal(todo)

	if er != nil {
		fmt.Println("error", er)
		panic(er)
	}
	fmt.Println(" json data", string(jsonByte))

	var res, err = http.Post(url, "application/json", bytes.NewBuffer(jsonByte))

	defer res.Body.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("status code ", res.StatusCode)
	dataBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var stringBuilder strings.Builder

	stringCount, err := stringBuilder.Write(dataBytes)

	if err != nil {
		panic(err)
	}
	fmt.Println("string count", stringCount)

	fmt.Println("response is ", stringBuilder.String())
	// fmt.Println("reponse is ", string(dataBytes))
}

func EncodeJson() {
	type dataType struct {
		Title  string `json:"title"` // this will chagne to key name when convert to JSON
		Body   string
		UserId int
	}

	var todo = dataType{"foo", "bar", 1}

	jsonByte, er := json.Marshal(todo)

	if er != nil {
		fmt.Println("error", er)
		panic(er)
	}
	fmt.Println(" json data", string(jsonByte))
}

func DecodeJson() {

	type Country struct {
		Name      string
		Capital   string
		Continent string
	}

	var myCountry Country

	fmt.Println("my country is ", myCountry)
	var data = []byte(`{
		"Name": "India",  
		"Capital": "New Delhi",
		"Continent": "Asia"
		}`)

	json.Unmarshal(data, &myCountry)
	fmt.Println("my country is ", myCountry)

}
