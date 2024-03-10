// internal/handler/hero_handler.go

package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"postgres-go-echo-htmx-bulma/pkg/db" // Import the generated code from sqlc
	"github.com/gorilla/mux"
)

// Hero represents the Hero model.
type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CreateHeroHandler handles the creation of a new hero.
func CreateHeroHandler(w http.ResponseWriter, r *http.Request) {
	var hero Hero
	if err := json.NewDecoder(r.Body).Decode(&hero); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err := db.CreateHero(ctx, hero.ID, hero.Name)
	if err != nil {
		http.Error(w, "Failed to create hero", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetHeroHandler retrieves a hero by ID.
func GetHeroHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing ID in the request", http.StatusBadRequest)
		return
	}

	// Convert the ID to an integer (you might want to add error handling here)
	heroID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID in the request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	hero, err := db.GetHeroByID(ctx, heroID)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch hero", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(hero)
}

// UpdateHeroHandler updates a hero by ID.
func UpdateHeroHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing ID in the request", http.StatusBadRequest)
		return
	}

	// Convert the ID to an integer (you might want to add error handling here)
	heroID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID in the request", http.StatusBadRequest)
		return
	}

	var updatedHero Hero
	if err := json.NewDecoder(r.Body).Decode(&updatedHero); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = db.UpdateHero(ctx, heroID, updatedHero.Name)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Failed to update hero", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteHeroHandler deletes a hero by ID.
func DeleteHeroHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing ID in the request", http.StatusBadRequest)
		return
	}

	// Convert the ID to an integer (you might want to add error handling here)
	heroID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID in the request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = db.DeleteHero(ctx, heroID)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Failed to delete hero", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
