package cron

//go:generate mockery --name=Service
type Service interface {
	Run()
}
