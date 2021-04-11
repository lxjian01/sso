package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sso/config"
	"sso/log"
	"time"
)

var (
	Sqlx *sqlx.DB
)

// 初始化sqlx
func InitSqlx(c *config.MysqlConfig){
	constr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=%v",c.User,c.Password,c.Host,c.Port,c.Db,c.Charset,"Asia%2fShanghai")
	Sqlx = sqlx.MustConnect("mysql", constr)
	// 设置连接可以被使用的最长有效时间
	Sqlx.SetConnMaxLifetime(10 * time.Minute)
	// 设置连接池中的保持连接的最大连接数
	Sqlx.SetMaxIdleConns(10)
	// 设置打开数据库的最大连接数
	Sqlx.SetMaxOpenConns(30)
	log.Info("Sqlx init ok")
}

// 关闭sqlx
func CloseSqlx() {
	if err := Sqlx.Close();err != nil{
		log.Error("Close jump db error by",err)
	}
}