package dao

import (
	"app/pkg/config"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var DB *gorm.DB

// 链接数据库
func Connect() {

	// 读取配置文件，初始化mongo链接配置信息
	conf := config.GetConfig()
	url := "sqlserver://" + conf.DB_USER + ":" + conf.DB_PWD + "@" + conf.DB_HOST + "?database=" + conf.DB_NAME
	fmt.Println(url)

	// 链接数据库
	var err error
	DB, err = gorm.Open("mssql", url)
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	// 开始数据库日志
	DB.LogMode(true)
	// 空闲连接的数量，不需要设置的太高，Go会检查并在必要时自动减少
	DB.DB().SetMaxIdleConns(50)
	// 允许并发连接量
	DB.DB().SetMaxOpenConns(50)
	// 设置连接池内连接可重用的时间为30分钟
	DB.DB().SetConnMaxLifetime(time.Second * 30)
	fmt.Println("链接成功，表格更新完毕")

}
