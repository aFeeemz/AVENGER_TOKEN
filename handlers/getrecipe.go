package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"NGC_AVENGER/config"
	"NGC_AVENGER/models"

	"github.com/julienschmidt/httprouter"
)

func GetAllRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rows, err := config.DB.Query("SELECT nama_resep, deskripsi_resep, waktu_masak, rating FROM recipe")

	if err != nil {
		log.Println("Database query failed:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe

		if err := rows.Scan(&recipe.Name, &recipe.Description, &recipe.Time, &recipe.Rating); err != nil {
			log.Println("Error when scanning rows", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		log.Println("error appending rows", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode the recipes slice to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(recipes); err != nil {
		log.Println("Error encoding response to JSON:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
