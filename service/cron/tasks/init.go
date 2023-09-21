package tasks

import (
	"context"

	"idev-cms-service/service/util"
	"idev-cms-service/service/util/logs"
)

type Tasks struct {
	ctx context.Context
	*CronTasksConfig
}

type CronTasksConfig struct {
	RepoTokens   util.Repository
	FilterString util.FilterString
	DateTime     util.DateTime
	Log          logs.Log
}

func New(config *CronTasksConfig) (tasks *Tasks) {
	return &Tasks{context.Background(), config}
}
