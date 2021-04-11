package models

import (
	"sso/utils/datetime"
)

type TokenResult struct {
	Token string `json:"token"`
}

type User struct {
	Id int                           `db:"id"`
	UserCode string                  `db:"user_code"`
	UserName string                  `db:"user_name"`
	Password string                  `db:"password"`
	CreateTime datetime.Datetime     `db:"create_time"`
	UpdateTIme datetime.Date         `db:"update_time"`
}
