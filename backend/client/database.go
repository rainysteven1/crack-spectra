package client

import (
	"backend/config"
	"backend/database/entity"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassWord, mysqlHost, mysqlPort, mysqlDBName)

	maxRetries := 3
	retryDelay := 10 * time.Second

	for i := 0; i <= maxRetries; i++ {
		DB, err = gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
		if err == nil {
			logrus.Info("Successfully connected to the database.")
			break
		}

		logrus.Errorf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries+1, err)
		if i < maxRetries {
			logrus.Infof("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		logrus.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries+1, err)
	}
	err = DB.AutoMigrate(
		&entity.Model{},
		&entity.ModelVersion{},
		&entity.Project{},
		&entity.Task{},
		&entity.User{},
		&entity.Role{},
		&entity.Permission{},
		&entity.UserProject{},
		&entity.UserRole{},
		&entity.RolePermission{},
	)
	if err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
	}

}
