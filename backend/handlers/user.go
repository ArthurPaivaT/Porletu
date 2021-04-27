package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ArthurPaivaT/Porletu/mongohandler"
	"go.mongodb.org/mongo-driver/mongo"
)

// validate tags are used to validate during mongohandler insert operation
type User struct {
	Userid   string
	Name     string `validate:"required"`
	LinkedIn string `validate:"required"`
	GitHub   string `validate:"required"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if len(r.Header["Userid"]) < 1 {
		err := fmt.Errorf("Missing Userid header")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userid := r.Header["Userid"][0]

	_, readErr := mongohandler.Read("Users", map[string]interface{}{"userid": userid})
	if readErr != nil {
		if readErr != mongo.ErrNoDocuments {
			err := fmt.Errorf("Error checking if user already exists: %w", readErr)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err := fmt.Errorf("user already exists")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var body []byte
	r.Body.Read(body)
	defer r.Body.Close()

	var newUser User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err := fmt.Errorf("Error reading request body: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		err := fmt.Errorf("Error creating user json: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newUser.Userid = userid

	insertResponse, err := mongohandler.Insert("Users", newUser)
	if err != nil {
		err := fmt.Errorf("Error inserting user into mongo: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(insertResponse)
	if err != nil {
		err := fmt.Errorf("Could not marshal insert response: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)
}
