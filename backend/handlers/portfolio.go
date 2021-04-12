package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type dev struct {
	Name     string `json:"name"`
	MainRole string `json:"mainRole"`
	LinkedIn string `json:"linkedIn"`
	GitHub   string `json:"gitHub"`
}

func GetDevInfo(w http.ResponseWriter, r *http.Request) {

	arthurDev := dev{
		Name:     "Arthur Paiva Tavares",
		MainRole: "Anything Developer",
		LinkedIn: "linkedin.com/in/arthur-paiva-982405199/",
		GitHub:   "github.com/arthurpaivat",
	}

	arthurDevJson, err := json.Marshal(arthurDev)
	if err != nil {
		err := fmt.Errorf("Error creating user json: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(arthurDevJson)
}
