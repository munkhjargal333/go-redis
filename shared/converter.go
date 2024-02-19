package shared

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func UintToString(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

// StringToInt
func StringToInt(num string) int {
	res, err := strconv.Atoi(num)
	if err != nil {
		res = 0
	}
	return res
}

func StringToUint(num string) uint {
	return uint(StringToInt(num))
}

// DateStringToTime
func DateStringToTime(t string) time.Time {
	c, err := time.Parse("2006-01-02", t)
	if err != nil {
		c, _ = time.Parse("2006-01-02", "0000-00-00")
	}

	return c
}

// TimeToDateString
func TimeToDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

// TimeDateTimeString
func TimeToDateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// TimeTimeString
func TimeToTimeString(t time.Time) string {
	return t.Format("15:04:05")
}

// MapToStruct : map to struct
func MapToStruct(m interface{}, v interface{}) (err error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return err
	}
	json.Unmarshal(jsonString, v)
	return err
}

func InterfaceToUint(val interface{}) (uint, error) {
	switch v := val.(type) {
	case uint:
		return v, nil
	case string:
		return StringToUint(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("cannot convert negative float to uint")
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("cannot convert negative int64 to uint")
		}
		return uint(v), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", val)
	}
}
