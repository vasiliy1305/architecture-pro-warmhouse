package main

import (
	"math/rand"
	"net/http"
	"time"
	"fmt" 
)

func main() {
	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")
		sensorID := r.URL.Query().Get("sensorId")
		
		if location == "" {
			switch sensorID {
			case "1": location = "Living Room"
			case "2": location = "Bedroom"
			case "3": location = "Kitchen"
			default: location = "Unknown"
			}
		}
		
		if sensorID == "" {
			switch location {
			case "Living Room": sensorID = "1"
			case "Bedroom": sensorID = "2"
			case "Kitchen": sensorID = "3"
			default: sensorID = "0"
			}
		}
		
		rand.Seed(time.Now().UnixNano())
		temp := 15 + rand.Float64()*15
		
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"value": ` + fmt.Sprintf("%.1f", temp) + `,
			"unit": "°C",
			"location": "` + location + `",
			"sensor_id": "` + sensorID + `"
		}`))
	})
	http.ListenAndServe(":8081", nil)
}