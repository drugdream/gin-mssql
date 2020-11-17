package models

import (
	"app/dao"
	"app/pkg/config"
	"log"
)

func TableSync() {
	conf := config.GetConfig()
	if !conf.TB_SYNC {
		log.Println("未开启表格同步")
	}
	dao.DB.Set("gorm:table_options", "charset=utf8mb4")
	tables := []interface{}{
		&User{},
		&Logger{},
	}
	for _, table := range tables {
		if dao.DB.HasTable(table) {
			dao.DB.AutoMigrate(table)
		} else {
			dao.DB.CreateTable(table)
		}
	}
}

func getParams(args ...interface{}) []interface{} {

	if len(args) == 1 && args[0] == nil {
		return []interface{}{}
	}

	vals := []interface{}{}
	for i := 0; i < len(args); i++ {
		vals = append(vals, args[i])
	}

	return vals

}
