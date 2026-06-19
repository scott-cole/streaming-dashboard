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
		version, err := client.Version()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"version": version})
	})

	http.HandleFunc("GET /api/scenes", func(w http.ResponseWriter, r *http.Request) {
		scenes, err := client.ListScenes()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string][]string{"scenes": scenes})
	})

	http.HandleFunc("GET /api/current-scene", func(w http.ResponseWriter, r *http.Request) {
		scene, err := client.CurrentScene()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"scene": scene})
	})

	http.HandleFunc("POST /api/switch", func(w http.ResponseWriter, r *http.Request) {

		var body struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if err := client.SwitchScene(body.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
