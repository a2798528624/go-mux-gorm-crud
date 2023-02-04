package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-mux-gorm-crud/internal/model"
	"log"
	"net/http"
	"strconv"
)

func respondWithJSON(w http.ResponseWriter, ok int, places []model.Place) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ok)
	err := json.NewEncoder(w).Encode(places)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	places, err := model.GetAllPlaces()
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, places)
		return
	}
	respondWithJSON(w, http.StatusOK, places)
}

func GetPlace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, nil)
		return
	}
	place, err := model.GetPlace(id)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, []model.Place{place})
		return
	}
	respondWithJSON(w, http.StatusOK, []model.Place{place})
}

func CreatePlace(w http.ResponseWriter, r *http.Request) {
	var place model.Place
	err := json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, []model.Place{place})
		return
	}
	place, err = model.CreatePlace(place)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, []model.Place{place})
		return
	}
	respondWithJSON(w, http.StatusCreated, []model.Place{place})
}

func DeletePlace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, nil)
		return
	}
	err = model.DeletePlace(id)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, nil)
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func UpdatePlace(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, nil)
		return
	}
	var place model.Place
	err = json.NewDecoder(r.Body).Decode(&place)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, []model.Place{place})
		return
	}
	place, err = model.UpdatePlace(id, place)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, []model.Place{place})
		return
	}
	respondWithJSON(w, http.StatusOK, []model.Place{place})
}
