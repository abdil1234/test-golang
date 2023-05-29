package appcontext

import (
	"errors"

	"github.com/abdil1234/test-golang/config"

	"github.com/abdil1234/test-golang/internal/pkg/driver"
	"gorm.io/gorm"
)

const (
	DBDialectMysql = "mysql"
)

type AppOption struct {
	Host string
	Port int
	Env  string
}

type AppContext struct {
	config *config.Configuration
}

func NewAppContext(config *config.Configuration) *AppContext {
	return &AppContext{
		config: config,
	}
}

func (a *AppContext) GetAppOption() AppOption {
	return AppOption{
		Host: a.config.App.Host,
		Port: a.config.App.Port,
		Env:  a.config.App.Env,
	}
}

func (a *AppContext) GetDBInstance(dbType string, dbConfig config.DatabaseConfiguration) (*gorm.DB, error) {
	var gorm *gorm.DB
	var err error
	switch dbType {
	case DBDialectMysql:
		dbOption := a.GetMysqlOption(dbConfig)
		gorm, err = driver.NewMysqlDatabase(dbOption)
	default:
		err = errors.New("error get db instance, unknown db type")
	}

	return gorm, err
}

func (a *AppContext) GetMysqlOption(dbConfig config.DatabaseConfiguration) driver.DBMysqlOption {
	return driver.DBMysqlOption{
		Host:                 dbConfig.Host,
		Port:                 dbConfig.Port,
		User:                 dbConfig.User,
		Password:             dbConfig.Password,
		DBName:               dbConfig.Name,
		AdditionalParameters: dbConfig.AdditionalParameters,
		MaxOpenConns:         dbConfig.MaxOpenConns,
		MaxIdleConns:         dbConfig.MaxIdleConns,
		ConnMaxLifetime:      dbConfig.ConnMaxLifetime,
	}
}
