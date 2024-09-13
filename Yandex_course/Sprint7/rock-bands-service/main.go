package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"

	"github.com/go-chi/chi/v5"
)

type Artist struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Born  string   `json:"born"`
	Genre string   `json:"genre"`
	Songs []string `json:"songs"`
}

var artists = map[string]Artist{
	"1": {
		ID:    "1",
		Name:  "30 seconds to Mars",
		Born:  "1998",
		Genre: "alternative",
		Songs: []string{
			`The Kill`,
			`A Beautiful Lie`,
			`Attack`,
			`Live Like A Dream`,
		},
	},
	"2": {
		ID:    "2",
		Name:  "Garbage",
		Born:  "1994",
		Genre: "alternative",
		Songs: []string{
			`Queer`,
			`Shut Your Mouth`,
			`Cup of Coffee`,
			`Til the Day I Die`,
		},
	},
}

func getArtists(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func postArtists(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &artist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	artists[artist.ID] = artist

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func getArtist (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	artist, ok := artists[id]
	if !ok {
		http.Error(w, "Артист не был найден", http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	r := chi.NewRouter()
	
	r.Get("/artists", getArtists)
	r.Post("/artists", postArtists)
	r.Get("/artist/{id}", getArtist)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка запуска сервера: %s\n", err.Error())
		return
	}
}