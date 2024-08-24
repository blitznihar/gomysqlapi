package main

import (
	"encoding/json"
	"gomysqlapi/gomysql"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getAlbums(w http.ResponseWriter, r *http.Request) {

	albums, err := gomysql.GetAlbums()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(albums)

}

func sendAlbums(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var albums []album
	json.Unmarshal(reqBody, &albums)
	var albumsMySql []gomysql.Album

	for i := 0; i < len(albums); i++ {
		albumsMySql = append(albumsMySql, gomysql.Album{
			Title:  albums[i].Title,
			Artist: albums[i].Artist,
			Price:  float32(albums[i].Price),
		})
	}

	err := gomysql.InsertAlbums(albumsMySql)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(albums)

}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/albums", getAlbums).Methods("GET")
	myRouter.HandleFunc("/albums", sendAlbums).Methods("POST")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {

	handleRequests()

}
