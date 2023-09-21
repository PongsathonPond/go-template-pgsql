package implement

import (
	"time"

	"github.com/jasonlvhit/gocron"
)

func (impl *implementation) Run() {
	time.Sleep(2 * time.Second)
	gocron.Clear()
	impl.Log.Info("Cron job scheduling running ...")

	if err := impl.runTask(EveryHour, impl.tasks.ClearTokensExpired, intervalConfig{hour: 8}, &runAtConfig{hour: 0}); err != nil {
		impl.Log.Error(err)
	}

	<-gocron.Start()
}
