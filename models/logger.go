package models

import (
	"app/dao"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

type Logger struct {
	gorm.Model
	User       uint      `gorm:"column:user"`
	StartTime  time.Time `gorm:"column:startTime"`
	EndTime    time.Time `gorm:"column:endTime"`
	UseTime    string    `gorm:"column:useTime"`
	IP         string    `gorm:"column:ip"`
	Method     string    `gorm:"column:method"`
	Url        string    `gorm:"column:url"`
	StatusCode int       `gorm:"column:statusCode"`
	Type       string    `gorm:"column:type"`
	ErrMsg     string    `gorm:"column:errMsg"`
	StackInfo  string    `gorm:"column:stackInfo;type:varchar(3200)"`
}

func (logger *Logger) Start() {
	logger.StartTime = time.Now()
}

func (logger *Logger) End(c *gin.Context) {

	// 结束时间
	logger.EndTime = time.Now()

	// 执行时间
	logger.UseTime = logger.EndTime.Sub(logger.StartTime).String()

	// 请求方式
	logger.Method = c.Request.Method

	// 请求路由
	logger.Url = c.Request.RequestURI

	// 状态码
	logger.StatusCode = c.Writer.Status()

	// 请求IP
	logger.IP = c.ClientIP()

	if logger.ErrMsg != "" {
		logger.Type = "error"
	} else {
		logger.Type = "normal"
	}

	if err := dao.Insert(logger); err != nil {
		fmt.Println("日志记录失败" + err.Error())
	}
}
