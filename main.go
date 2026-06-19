package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

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

	http.HandleFunc("GET /api/status", func(w http.ResponseWriter, r *http.Request) {
		status, err := client.Status()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		scene, err := client.CurrentScene()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{
			"streaming": status.Streaming,
			"recording": status.Recording,
			"scene":     scene,
		})
	})

	http.HandleFunc("POST /api/stream/start", func(w http.ResponseWriter, r *http.Request) {
		if err := client.StartStream(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("POST /api/stream/stop", func(w http.ResponseWriter, r *http.Request) {
		if err := client.StopStream(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("POST /api/record/start", func(w http.ResponseWriter, r *http.Request) {
		if err := client.StartRecord(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("POST /api/record/stop", func(w http.ResponseWriter, r *http.Request) {
		if err := client.StopRecord(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	http.HandleFunc("GET /api/preview", func(w http.ResponseWriter, r *http.Request) {
		scene := r.URL.Query().Get("scene")
		if scene == "" {
			http.Error(w, "missing scene param", http.StatusBadRequest)
			return
		}
		width := 640
		height := 360
		if w := r.URL.Query().Get("width"); w != "" {
			if v, err := strconv.Atoi(w); err == nil {
				width = v
			}
		}
		if h := r.URL.Query().Get("height"); h != "" {
			if v, err := strconv.Atoi(h); err == nil {
				height = v
			}
		}
		image, err := client.ScenePreview(scene, width, height)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"image": image})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
