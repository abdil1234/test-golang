package commons

import (
	"github.com/abdil1234/test-golang/config"
	"github.com/abdil1234/test-golang/internal/app/commons/appcontext"
	"gorm.io/gorm"
)

type Option struct {
	AppCtx   *appcontext.AppContext
	Config   *config.Configuration
	Database *gorm.DB
}
