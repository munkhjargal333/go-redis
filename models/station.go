package models

import (
	"fmt"

	"route-redis/shared"
)

type Station struct {
	ID           uint    `json:"id" gorm:"column:id"`
	BusStopID    int     `json:"busstop_id" gorm:"column:busstop_id"`
	BusRouteID   int     `json:"busroute_id" gorm:"column:busroute_id"`
	BusStopSeq   int     `json:"busstop_seq" gorm:"column:busstop_seq"`
	BusStopLen   int     `json:"busstop_len" gorm:"column:busstop_len"`
	AFCBusStopID string  `json:"afcbusstop_id" gorm:"column:afcbusstop_id"`
	StopFlag     int     `json:"stop_flag" gorm:"column:stop_flag"`
	BusStopName  string  `json:"busstop_name" gorm:"column:busstop_name"`
	GPXx         float32 `json:"gpx_x" gorm:"column:gpx_x"`
	GPXy         float32 `json:"gpx_y" gorm:"column:gpx_y"`
}

func (u *Station) TableName() string {
	return fmt.Sprintf("%s.ip_stations", shared.Config.DB.Schema)
}
