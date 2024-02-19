package models

import (
	"fmt"
	"route-redis/shared"
)

type RouteLine struct {
	ID         uint    `json:"id" gorm:"column:id"`
	BusRouteID int     `json:"busroute_id" gorm:"column:busroute_id"`
	Seq        int     `json:"seq" gorm:"column:seq"`
	GPXx       float32 `json:"gpx_x" gorm:"column:gpx_x"`
	GPXy       float32 `json:"gpx_y" gorm:"column:gpx_y"`
}

func (u *RouteLine) TableName() string {
	return fmt.Sprintf("%s.ip_route_lines", shared.Config.DB.Schema)
}
