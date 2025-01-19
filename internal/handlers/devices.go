package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"project_green/internal/repository"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// DeviceRoutes initializes the routes for the device-related endpoints
func DeviceRoutes(r chi.Router) {
	r.Get("/devices", devicesHandler)
	r.Get("/devices/{id}", deviceHandler)
	r.Post("/devices/register", deviceRegisterHandler)
	r.Delete("/devices/delete/{id}", deviceDeleteHandler)
}

// HANDLER FUNCTIONS
// =================

// devicesHandler returns a list of devices
func devicesHandler(w http.ResponseWriter, r *http.Request) {
	// Sample data to return as JSON
	devices, err := repository.GetDevices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Devices: %+v\n", devices)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices) // Encode and send the JSON response
}

// deviceHandler returns details of a specific device
func deviceHandler(w http.ResponseWriter, r *http.Request) {
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

	// Sample data for the device
	device, err := repository.GetDeviceByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)

}

// deviceNewHandler simulates creating a new device
func deviceRegisterHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		DeviceName string `json:"device_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	device, err := repository.CreateDevice(repository.Device{
		DeviceName: input.DeviceName,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}

// deviceDeleteHandler simulates deleting a device
func deviceDeleteHandler(w http.ResponseWriter, r *http.Request) {
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

	err = repository.DeleteDevice(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "success"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
