package models

import (
	"fmt"

	"route-redis/database"

	"route-redis/shared"

	"gorm.io/gorm"
)

type Route struct {
	ID           int    `json:"id" gorm:"column:id"`
	Status       string `json:"status" gorm:"column:status"`
	BusRouteID   int    `json:"busroute_id" gorm:"column:busroute_id"`
	BusRouteNo   string `json:"busroute_no" gorm:"column:busroute_no"`
	BusRouteName string `json:"busroute_name" gorm:"column:busroute_name"`

	Model
}

func (u *Route) TableName() string {
	return fmt.Sprintf("%s.ip_routes", shared.Config.DB.Schema)
}

func (u *Route) Migrate() error {
	db := database.DB

	if err := db.AutoMigrate(&Route{}); err != nil {
		return err
	}

	return nil
}

func (u *Route) Drop() error {
	db := database.DB

	if err := db.Exec("drop table ?", gorm.Expr(u.TableName())).Error; err != nil {
		return err
	}

	return nil
}

type RouteAmount struct {
	//	Model
	RouteID    uint `json:"route_id"`
	CardTypeID uint `json:"card_type_id"`
	Amount     uint `json:"amount"`
}

func (r *RouteAmount) TableName() string {
	return fmt.Sprintf("%s.bs_route_amounts", shared.Config.DB.Schema)
}

func (r *RouteAmount) Migrate() error {
	db := database.DB

	if err := db.AutoMigrate(&RouteAmount{}); err != nil {
		return err
	}

	return nil
}

func (r *RouteAmount) Drop() error {
	db := database.DB

	if err := db.Exec("drop table ?", gorm.Expr(r.TableName())).Error; err != nil {
		return err
	}

	return nil
}
