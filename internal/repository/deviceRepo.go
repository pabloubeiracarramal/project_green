package repository

import (
	"context"
	"fmt"
	"project_green/db"
	"time"
)

type Device struct {
	DeviceID   int       `json:"device_id"`
	DeviceName string    `json:"device_name"`
	CreatedAt  time.Time `json:"created_at"`
}

func GetDevices() ([]Device, error) {
	// Print a debug message
	fmt.Println("GetDevices called")

	// Execute the query using pgx
	rows, err := db.DB.Query(context.Background(), "SELECT device_id, device_name, created_at FROM devices")
	if err != nil {
		return nil, fmt.Errorf("error querying devices: %w", err)
	}
	defer rows.Close()

	// Slice to hold devices
	var devices []Device

	// Iterate over the rows
	for rows.Next() {
		var device Device
		if err := rows.Scan(&device.DeviceID, &device.DeviceName, &device.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning devices: %w", err)
		}
		devices = append(devices, device)
	}

	// Check for errors during iteration
	if rows.Err() != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", rows.Err())
	}

	return devices, nil
}

func GetDeviceByID(deviceID int) (Device, error) {
	// Print a debug message
	fmt.Println("GetDeviceByID called")

	// Query the database for the device
	var device Device
	err := db.DB.QueryRow(context.Background(), "SELECT device_id, device_name, created_at FROM devices WHERE device_id = $1", deviceID).Scan(&device.DeviceID, &device.DeviceName, &device.CreatedAt)
	if err != nil {
		return Device{}, fmt.Errorf("error querying device: %w", err)
	}

	return device, nil
}

func CreateDevice(device Device) (Device, error) {
	// Print a debug message
	fmt.Println("CreateDevice called")

	device.CreatedAt = time.Now()

	// Execute the insert query
	err := db.DB.QueryRow(
		context.Background(),
		"INSERT INTO devices (device_name, created_at) VALUES ($1, $2) RETURNING device_id",
		device.DeviceName, device.CreatedAt,
	).Scan(&device.DeviceID)
	if err != nil {
		return device, fmt.Errorf("error inserting device: %w", err)
	}

	return device, nil

}

func DeleteDevice(deviceID int) error {
	// Print a debug message
	fmt.Println("DeleteDevice called")

	// Execute the delete query
	_, err := db.DB.Exec(context.Background(), "DELETE FROM devices WHERE device_id = $1", deviceID)
	if err != nil {
		return fmt.Errorf("error deleting device: %w", err)
	}

	return nil
}
