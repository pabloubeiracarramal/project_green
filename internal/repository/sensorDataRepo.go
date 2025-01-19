package repository

import (
	"context"
	"fmt"
	"project_green/db"
	"time"
)

type SensorData struct {
	DeviceID     int       `json:"device_id"`
	Temp         float64   `json:"temp"`
	Humidity     float64   `json:"humidity"`
	LightLevel   float64   `json:"light_level"`
	SoilMoisture float64   `json:"soil_moisture"`
	WaterLevel   float64   `json:"water_level"`
	DateTime     time.Time `json:"date_time"`
}

func GetLatestSensorData(deviceID int) (SensorData, error) {
	// Print a debug message
	fmt.Println("GetLatestSensorData called")

	// Query the database for the latest sensor data
	var sensorData SensorData
	err := db.DB.QueryRow(context.Background(), "SELECT device_id, temp, humidity, light_level, soil_moisture, water_level, date_time FROM sensor_data WHERE device_id = $1 ORDER BY date_time DESC LIMIT 1", deviceID).Scan(&sensorData.DeviceID, &sensorData.Temp, &sensorData.Humidity, &sensorData.LightLevel, &sensorData.SoilMoisture, &sensorData.WaterLevel, &sensorData.DateTime)
	if err != nil {
		return SensorData{}, fmt.Errorf("error querying sensor data: %w", err)
	}

	return sensorData, nil
}

func GetHistoricSensorData(deviceID int) ([]SensorData, error) {
	// Print a debug message
	fmt.Println("GetHistoricSensorData called")

	// Execute the query using pgx
	rows, err := db.DB.Query(context.Background(), "SELECT device_id, temp, humidity, light_level, soil_moisture, water_level, date_time FROM sensor_data WHERE device_id = $1 ORDER BY date_time DESC", deviceID)
	if err != nil {
		return nil, fmt.Errorf("error querying sensor data: %w", err)
	}
	defer rows.Close()

	// Slice to hold sensor data
	var sensorData []SensorData

	// Iterate over the rows
	for rows.Next() {
		var data SensorData
		if err := rows.Scan(&data.DeviceID, &data.Temp, &data.Humidity, &data.LightLevel, &data.SoilMoisture, &data.WaterLevel, &data.DateTime); err != nil {
			return nil, fmt.Errorf("error scanning sensor data: %w", err)
		}
		sensorData = append(sensorData, data)
	}

	// Check for errors during iteration
	if rows.Err() != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", rows.Err())
	}

	return sensorData, nil
}

func GetSensorDataByPeriod(deviceID int, start time.Time, end time.Time) ([]SensorData, error) {
	// Print a debug message
	fmt.Println("GetSensorDataByPeriod called")

	// Execute the query using pgx
	rows, err := db.DB.Query(context.Background(), "SELECT device_id, temp, humidity, light_level, soil_moisture, water_level, date_time FROM sensor_data WHERE device_id = $1 AND date_time BETWEEN $2 AND $3 ORDER BY date_time DESC", deviceID, start, end)
	if err != nil {
		return nil, fmt.Errorf("error querying sensor data: %w", err)
	}
	defer rows.Close()

	// Slice to hold sensor data
	var sensorData []SensorData

	// Iterate over the rows
	for rows.Next() {
		var data SensorData
		if err := rows.Scan(&data.DeviceID, &data.Temp, &data.Humidity, &data.LightLevel, &data.SoilMoisture, &data.WaterLevel, &data.DateTime); err != nil {
			return nil, fmt.Errorf("error scanning sensor data: %w", err)
		}
		sensorData = append(sensorData, data)
	}

	// Check for errors during iteration
	if rows.Err() != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", rows.Err())
	}

	return sensorData, nil
}

func SendSensorData(data SensorData) (SensorData, error) {
	// Print a debug message
	fmt.Println("SendSensorData called")

	// Execute the query using pgx
	_, err := db.DB.Exec(context.Background(), "INSERT INTO sensor_data (device_id, temp, humidity, light_level, soil_moisture, water_level, date_time) VALUES ($1, $2, $3, $4, $5, $6, $7)", data.DeviceID, data.Temp, data.Humidity, data.LightLevel, data.SoilMoisture, data.WaterLevel, data.DateTime)
	if err != nil {
		return SensorData{}, fmt.Errorf("error inserting sensor data: %w", err)
	}

	return data, nil
}
