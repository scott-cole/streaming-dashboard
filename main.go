package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"streaming-dashboard/obs"
)

func main() {
	password := os.Getenv("OBS_PASSWORD")

	if password == "" {
		log.Fatal("OBS_PASSWORD env var not set")
	}

	client, err := obs.New(password)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect()

	log.Println("Server starting on :8080")

	http.HandleFunc("GET /api/version", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"version": client.Version()})
	})

	http.HandleFunc("GET /api/scenes", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string][]string{
			"scenes": client.ListScenes()})
	})

	http.HandleFunc("POST /api/switch", func(w http.ResponseWriter, r *http.Request) {

		var body struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		client.SwitchScene(body.Name)

		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok"})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
