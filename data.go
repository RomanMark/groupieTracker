package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var API Api

type Api struct {
	Artists   Artists
	Dates     Dates
	Locations Locations
	Relation  Relation
}

type Artists []struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations"`
	Dates        []string            `json:"dates"`
	Relations    map[string][]string `json:"relations"`
}

type Dates struct {
	Index []struct {
		Id    int
		Dates []string
	}
}
type Locations struct {
	Index []struct {
		Id        int
		Locations []string
	}
}
type Relation struct {
	Index []struct {
		Id             int
		DatesLocations map[string][]string
	}
}

func getData() {

	getJson("https://groupietrackers.herokuapp.com/api/artists", &API.Artists)
	getJson("https://groupietrackers.herokuapp.com/api/dates", &API.Dates)
	getJson("https://groupietrackers.herokuapp.com/api/locations", &API.Locations)
	getJson("https://groupietrackers.herokuapp.com/api/relation", &API.Relation)

	for i := range API.Artists {

		API.Artists[i].Dates = API.Dates.Index[i].Dates
		API.Artists[i].Locations = API.Locations.Index[i].Locations
		API.Artists[i].Relations = fixNames(API.Relation.Index[i].DatesLocations)
	}

}

func getJson(URL string, target interface{}) {
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &target)
	if err != nil {
		panic(err)
	}

}

func fixNames(rel map[string][]string) map[string][]string {

	newMap := make(map[string][]string)
	for index := range rel {
		name := strings.Replace(index, "_", " ", -1)
		name = strings.Replace(name, "-", ", ", -1)
		name = strings.Title(strings.ToLower(name))
		name = strings.Replace(name, "Usa", "USA", -1)
		name = strings.Replace(name, "Uk", "UK", -1)
		newMap[name] = rel[index]
	}
	return newMap
}
