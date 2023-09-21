package migrations

import (
	"context"

	"idev-cms-service/config"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
)

var (
	ctx     = context.Background()
	uuid, _ = util.NewUUID()
	encrypt = util.NewEncrypt()
	log     = logs.New(&config.Config{
		AppLogLevel:       0,
		AppLogEnableColor: true,
	})
)
