package models

import "time"

type Inverter struct {
	ID                          int64     `json:"id"`
	Date                        time.Time `json:"date"`
	Energy_per_Inverter_kWh     float32   `json:"energy_per_inverter_kwh"`
	Energy_per_inverter_per_kWp float32   `json:"energy_per_inverter_per_kwp"`
	Total_system_kWh            float32   `json:"total_system_kwh`
}
