package ctx

import (
	"github.com/SinedVonotirah/gopo/core/config"
	"github.com/SinedVonotirah/gopo/orms/beego"
	"github.com/SinedVonotirah/gopo/orms/gorm"
)

var (
	GlobalContext = loadContext()
)

type ApplicationContext struct {
	AppConfig *config.Config
}

func loadContext() *ApplicationContext {
	config := config.GetConfig("gopo_config")

	beego.InitBenchSuite("host=localhost user=postgres password=1 dbname=gopo sslmode=disable", 10)
	gorm.InitBenchSuite("host=localhost user=postgres password=1 dbname=gopo sslmode=disable", 10)

	return &ApplicationContext{
		AppConfig: config,
	}
}
