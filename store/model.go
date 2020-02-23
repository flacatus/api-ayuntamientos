package store

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	//CORDINATES struct for townhall object
	CORDINATES struct {
		LATITUDE  string `bson:"latitude" json:"latitude"`
		LONGITUDE string `bson:"longitude" json:"longitude"`
	}
	//Party: aa
	Party struct {
		PHOTO string `bson:"photo" json:"photo"`
		NAME  string `bson:"name" json:"name"`
	}
	//Maior: brief data
	Maior struct {
		AGE       string `bson:"age" json:"age"`
		NAME      string `bson:"name" json:"name"`
		PHOTO     string `bson:"photo" json:"photo"`
		COMPLETED string `bson:"completed" json:"completed"`
		PARTY     Party  `bson:"party" json:"party"`
	}
	//Cities with their key value
	Cities struct {
		//ID             bson.ObjectId `bson:"_id" json:"id"`
		NAME           string     `bson:"city_name" json:"city_name"`
		CORDINATES     CORDINATES `bson:"coordinates" json:"coordinates"`
		MAIOR          Maior      `bson:"maior" json:"maior"`
		POLITICALPARTY string     `bson:"politicalparty" json:"politicalparty"`
		POPULATION     string     `bson:"population" json:"population"`
		WEBSITE        string     `bson:"website" json:"website"`
	}
	//TownHalls object with all key:value
	TownHalls struct {
		ID       bson.ObjectId `bson:"_id" json:"id"`
		CITIES   []Cities      `bson:"cities" json:"cities"`
		COMUNITY string        `bson:"comunity" json:"comunity"`
		COUNTRY  string        `bson:"country" json:"country"`
	}
)

//TownHall: Obj to slice
type TownHall []TownHalls
