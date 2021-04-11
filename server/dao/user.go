package dao

import (
	"sso/db/sqlx"
	"sso/server/models"
)

func GetUserByUserCode(userCode string) (models.User,error){
	sql := "select * from user where user_code = ?"
	user := models.User{}
	err := sqlx.Sqlx.Get(&user, sql, userCode)
	return user,err
}