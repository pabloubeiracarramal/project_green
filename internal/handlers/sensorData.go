package handlers

import (
	"encoding/json"
	"net/http"
	"project_green/internal/repository"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// DeviceRoutes initializes the routes for the device-related endpoints
func SensorDataRoutes(r chi.Router) {
	r.Get("/sensorData/latest/{id}", sensorDataLatestHandler)
	r.Get("/sensorData/historic/{id}", sensorDataHistoricHandler)
	r.Post("/sensorData/period/{id}", sensorDataPeriodHandler)
	r.Post("/sensorData/sendData/{id}", sensorDataSendDataHandler)
}

// HANDLER FUNCTIONS
// =================

// sensorDataLatestHandler returns the latest sensor data for a specific device
func sensorDataLatestHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "id")
	if deviceID == "" {
		http.Error(w, "Device ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(deviceID)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	// Sample data for the latest sensor data
	sensorData, err := repository.GetLatestSensorData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensorData) // Encode and send the JSON response
}

// sensorDataHistoricHandler returns historic sensor data for a specific device
func sensorDataHistoricHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "id")
	if deviceID == "" {
		http.Error(w, "Device ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(deviceID)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	// Sample data for the historic sensor data
	sensorData, err := repository.GetHistoricSensorData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensorData) // Encode and send the JSON response
}

// sensorDataPeriodHandler returns sensor data for a specific device within a specified period
func sensorDataPeriodHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "id")
	if deviceID == "" {
		http.Error(w, "Device ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(deviceID)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	startTime, err := time.Parse(time.RFC3339, input.Start)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	endTime, err := time.Parse(time.RFC3339, input.End)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}

	// Sample data for the sensor data within the specified period
	sensorData, err := repository.GetSensorDataByPeriod(id, startTime, endTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensorData) // Encode and send the JSON response
}

// sensorDataSendDataHandler sending sensor data to the server
func sensorDataSendDataHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "id")
	if deviceID == "" {
		http.Error(w, "Device ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(deviceID)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Temp         float64 `json:"temp"`
		Humidity     float64 `json:"humidity"`
		LightLevel   float64 `json:"light_level"`
		SoilMoisture float64 `json:"soil_moisture"`
		WaterLevel   float64 `json:"water_level"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Sample data for the sensor data
	sensorData, err := repository.SendSensorData(repository.SensorData{
		DeviceID:     id,
		Temp:         input.Temp,
		Humidity:     input.Humidity,
		LightLevel:   input.LightLevel,
		SoilMoisture: input.SoilMoisture,
		WaterLevel:   input.WaterLevel,
		DateTime:     time.Now(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensorData) // Encode and send the JSON response
}
