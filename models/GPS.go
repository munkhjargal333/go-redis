package models

type GPSData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	BusID     int     `json:"busId"`
	RouteID   int     `json:"routeId"`
	Direction int     `json:"direction"`
	DeviceID  int     `json:"deviceId"`
}
