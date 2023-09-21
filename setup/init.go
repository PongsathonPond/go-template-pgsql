package setup

import (
	"idev-cms-service/config"
	"idev-cms-service/service/util/logs"
)

type setup struct {
	appConfig *config.Config
	log       logs.Log
}

func New(appConfig *config.Config, log logs.Log) *setup {
	return &setup{appConfig, log}
}
