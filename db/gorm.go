package db

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"sso/config"
	"sso/log"
)

var(
	Gorm *gorm.DB
)

// 初始化sqlx
func InitGorm(c *config.MysqlConfig) {

	log.Info("Gorm init ok")
}
