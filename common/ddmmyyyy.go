package common

import "fmt"

type DDMMYYYY struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

func (d DDMMYYYY) String() string {
	return fmt.Sprintf("%02d/%02d/%d", d.Day, d.Month, d.Year)
}
