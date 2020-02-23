package store

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

//Controller : Todo
type Controller struct {
	Repository Repository
}

//GetAllTownHalls : Method to return all townhalls
func (c *Controller) GetAllTownHalls(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if townhall := c.Repository.GetAllTownhalls(); townhall != nil {

		respondWithJSON(w, http.StatusCreated, townhall)
		return
	}
}

//InsertTHalls : Methods to insert townhalls
func (c *Controller) InsertTHalls(w http.ResponseWriter, r *http.Request) {
	var townhall TownHalls

	defer r.Body.Close()
	townhall.ID = bson.NewObjectId()

	if err := json.NewDecoder(r.Body).Decode(&townhall); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := c.Repository.InsertTHalls(townhall); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, townhall)

}

// respondWithError if the request fail i send the error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

//respondWithJSON generate the response with the items from db
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
