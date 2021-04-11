package services

import (
	"encoding/json"
	"errors"
	"sso/log"
	"sso/server/dao"
	"sso/server/models"
	"sso/utils/jwt"
)

// login
func Login(userCode string,password string) (*models.TokenResult,error){
	user,err := dao.GetUserByUserCode(userCode)
	if err != nil{
		return nil,err
	}
	userInfo,err := json.Marshal(user)
	log.Info(string(userInfo))
	if user.Password != password {
		return nil,errors.New("密码错误！")
	}
	j := jwt.NewJWT()
	token,err1 := j.CreateToken(user.Id,user.UserCode,user.UserName)
	if err1 != nil{
		return nil,err
	}
	log.Info(token)
	tokenResult := &models.TokenResult{Token:token}
	return tokenResult,nil
}