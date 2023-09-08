package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		slackName := r.URL.Query().Get("slack_name")
		track := r.URL.Query().Get("track")

		// Get current UTC time with +/-2 minutes window
		currentTime := time.Now().UTC()
		utcTime := currentTime.Format("2006-01-02T15:04:05Z")

		// Prepare the JSON response
		response := map[string]interface{}{
			"slack_name":     slackName,
			"current_day":    currentTime.Weekday().String(),
			"utc_time":       utcTime,
			"track":          track,
			"github_file_url": "https://github.com/Shiewhun/hng1/blob/main/main.go",
			"github_repo_url": "https://github.com/Shiewhun/hng1",
			"status_code":    200,
		}

		// Convert the response to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the content type and write the response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
