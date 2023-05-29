package driver

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBMysqlOption struct {
	Host                 string
	Port                 int
	User                 string
	Password             string
	DBName               string
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
}

func NewMysqlDatabase(option DBMysqlOption) (*gorm.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		option.User, option.Password, option.Host, option.Port, option.DBName, option.AdditionalParameters))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(option.MaxOpenConns)
	db.SetMaxIdleConns(option.MaxIdleConns)
	db.SetConnMaxLifetime(option.ConnMaxLifetime)

	gorm, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return gorm, nil
}
