package component

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nikhildev/bugsbunny/clients"
	"github.com/nikhildev/bugsbunny/models"
)

func CreateComponentHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error reading request body"))
		return
	}

	var component models.Component
	err = json.Unmarshal(body, &component)

	fmt.Println("component", component)

	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error unmarshalling request body", err)
		return
	}

	db, err := clients.GetDbClient()

	result := db.Create(&component)
	if result.Error != nil {
		w.WriteHeader(500)
		fmt.Println("Error creating component", result.Error)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(component)
}
