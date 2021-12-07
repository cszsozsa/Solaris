package models

import "time"

type ElectricMeter struct {
	ID         int64     `json:"id"`
	Import_kWh int64     `json:"import_kwh"`
	Export_kWh int64     `json:"export_kwh"`
	Comment    string    `json:"comment"`
	Timestamp  time.Time `json:"timestamp"`
}
