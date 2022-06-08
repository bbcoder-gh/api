package models

import (
	"github.com/bbcoder-gh/api/common"
)

type User struct {
	ID int `json:"id"`
	common.Name
	DOB   common.DDMMYYYY `json:"dob"`
	Email string          `json:"email"`
	Phone string          `json:"phone"`
}
