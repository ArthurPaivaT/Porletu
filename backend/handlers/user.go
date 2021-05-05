package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ArthurPaivaT/Porletu/mongohandler"
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if len(r.Header["Userid"]) < 1 {
		err := fmt.Errorf("Missing Userid header")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userid := r.Header["Userid"][0]

	deleteResult, err := mongohandler.Delete("Users", map[string]interface{}{"userid": userid})
	if err != nil || deleteResult.DeletedCount < 1 {
		err := fmt.Errorf("Error deleting user: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
