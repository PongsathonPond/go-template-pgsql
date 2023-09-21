package implement

import (
	"context"

	"idev-cms-service/service/cron"
	"idev-cms-service/service/cron/tasks"
	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
)

type implementation struct {
	ctx context.Context
	util.DateTime
	logs.Log
	tasks *tasks.Tasks
}

type CronServiceConfig struct {
	RepoTokens   util.Repository
	FilterString util.FilterString
	DateTime     util.DateTime
	Log          logs.Log
}

func New(config *CronServiceConfig) (service cron.Service) {
	return &implementation{
		ctx:      context.Background(),
		DateTime: config.DateTime,
		Log:      config.Log,
		tasks: tasks.New(&tasks.CronTasksConfig{
			RepoTokens:   config.RepoTokens,
			FilterString: config.FilterString,
			DateTime:     config.DateTime,
			Log:          config.Log,
		}),
	}
}
