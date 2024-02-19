package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	//	ID            uint           `json:"id" gorm:"primarykey" gorm:"column:id"`
	ObjectVersion uint           `json:"-" gorm:"default:1"`
	CreatedByID   uint           `json:"-"`
	CreatedAt     *time.Time     `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedByID   uint           `json:"-"`
	UpdatedAt     *time.Time     `json:"-" gorm:"autoUpdateTime"`
	DeletedByID   uint           `json:"-" gorm:"column:deleted"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type Address struct {
	AimagCode *string `json:"aimag_code,omitempty" gorm:"type:varchar(2)"`
	//City      City    `json:"city" gorm:"foreignKey:ID;references:AimagCode"`

	SumCode *string `json:"sum_code,omitempty" gorm:"type:varchar(2)"`
	//District District `json:"district" gorm:"foreignKey:ID;references:SumCode"`

	BagCode *string `json:"bag_code,omitempty" gorm:"type:varchar(2)"`
	//Khoroo  Khoroo  `json:"khoroo" gorm:"foreignKey:ID;references:BagCode"`
}

type Migrater interface {
	TableName() string
	Migrate() error
	Drop() error
}

type LocalTimeZone time.Time

const localTimeZoneFormat = "2006-01-02 15:04:05"

func (t *LocalTimeZone) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+localTimeZoneFormat+`"`, string(data), time.Local)
	*t = LocalTimeZone(now)
	return
}

func (t LocalTimeZone) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localTimeZoneFormat)+2)
	b = append(b, '"')
	b = append(b, []byte(t.String())...)
	b = append(b, '"')

	return b, nil
}

func (t LocalTimeZone) String() string {
	if time.Time(t).IsZero() {
		return "0000-00-00 00:00:00"
	}

	return time.Time(t).Format(localTimeZoneFormat)
}

func (t LocalTimeZone) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return time.Now().Local(), nil
	}
	return time.Time(t).Local(), nil
}

func (t *LocalTimeZone) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = LocalTimeZone(vt)
	case string:
		tTime, _ := time.Parse("2006-01-02 15:04:05", vt)
		*t = LocalTimeZone(tTime)
	default:
		return nil
	}
	return nil
}

type LocalDate time.Time

const localDateFormat = "2006-01-02 15:04:05"

func (t *LocalDate) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+localDateFormat+`"`, string(data), time.Local)
	*t = LocalDate(now)
	return
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localDateFormat)+2)
	b = append(b, '"')
	b = append(b, []byte(t.String())...)
	b = append(b, '"')

	return b, nil
}

func (t LocalDate) String() string {
	if time.Time(t).IsZero() {
		return "0000-00-00"
	}

	return time.Time(t).Format(localDateFormat)
}

func (t LocalDate) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return time.Now(), nil
	}
	return time.Time(t), nil
}

func (t *LocalDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = LocalDate(vt)
	case string:
		tTime, _ := time.Parse("2006/01/02 15:04:05", vt)
		*t = LocalDate(tTime)
	default:
		return nil
	}
	return nil
}

func (model *Model) AfterCreate(tx *gorm.DB) error {
	fmt.Println("Fucked up after create")

	return nil
}

func (model *Model) AfterUpdate(tx *gorm.DB) error {
	fmt.Println("Fucked up after update")

	return nil
}

func (model *Model) AfterDelete(tx *gorm.DB) error {
	fmt.Println("Fucked up after delete")

	return nil
}
