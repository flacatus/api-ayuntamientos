package store

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

//Repository struct
type Repository struct{}

// DBNAME the name of the DB instance
const DBNAME = "townhalls"

//collections used in this package
const (
	APIKCOLL = "townhalls"
)

//noSslConnect: TODO The idea of this is to migrate to an ssl connection to mongoDB
func noSslConnect() (*mgo.Session, error) {
	session, err := mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		panic(err)
	}

	return session, err
}

// GetAllTownhalls returns the list of Products
func (r Repository) GetAllTownhalls() TownHall {
	var results TownHall

	session, err := noSslConnect()
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(APIKCOLL)
	log.Println("Getting records of cities...")
	if err := c.Find(nil).All(&results); err != nil {
		log.Println("Failed to write results:", err)
	}

	return results
}

//InsertTHalls :Methods to allow inserts.Will be useful with the spider...
func (r Repository) InsertTHalls(townhalls TownHalls) error {
	session, err := noSslConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	log.Println("Inserting new record of Cities..")
	for _, key := range []string{"comunity"} {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := session.DB(DBNAME).C(APIKCOLL).EnsureIndex(index); err != nil {
			return err
		}
	}
	err = session.DB(DBNAME).C(APIKCOLL).Insert(&townhalls)

	return err
}
