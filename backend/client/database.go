package client

import (
	"backend/config"
	"backend/database/entity"
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局 DB 对象
var DB *gorm.DB

func InitDB() {
	var err error
	mysqlUser := config.GetString("database.mysql_user")
	mysqlPassWord := config.GetString("database.mysql_password")
	mysqlHost := config.GetString("database.mysql_host")
	mysqlPort := config.GetString("database.mysql_port")
	mysqlDBName := config.GetString("database.mysql_db")

	// 在InitDB函数中添加网络测试
	conn, err := net.DialTimeout("tcp", mysqlHost+":"+"6379", 5*time.Second)
	if err != nil {
		logrus.Errorf("Network connection failed: %v", err)
		return
	}
	conn.Close()
	logrus.Info("Network connection successful")

	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassWord, mysqlHost, mysqlPort, mysqlDBName)
	logrus.Info(mysqlDSN)
	if DB, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		logrus.Errorf("Failed to connect to database: %v", err)
	}

	if err = DB.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.Permission{},
		&entity.UserRole{}, &entity.RolePermission{},
		&entity.Project{}, &entity.UserProject{}, &entity.Task{},
		&entity.Model{}, &entity.ModelVersion{},
	); err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
	}
}
