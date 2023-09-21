package implement

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/jasonlvhit/gocron"
)

type runType uint8

const (
	EveryMinute runType = iota
	EveryHour
	EveryDay
)

type intervalConfig struct {
	day  int
	hour int
	min  int
}

type runAtConfig struct {
	hour int
	min  int
}

func (impl *implementation) runTask(rt runType, task interface{}, intervalConfig intervalConfig, runAtConfig *runAtConfig) error {
	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	min := time.Now().Minute()

	if rt == EveryMinute {
		if intervalConfig.min == 0 {
			return errors.New("invalid param")
		}

		if runAtConfig != nil && runAtConfig.hour != 0 {
			hour = runAtConfig.hour
		}

		var minRunAt float64
		if intervalConfig.min > 1 {
			if runAtConfig != nil && runAtConfig.min != 0 {
				minRunAt = math.Ceil(float64(min)/float64(runAtConfig.min)) * float64(runAtConfig.min)
			} else {
				minRunAt = math.Ceil(float64(min)/float64(intervalConfig.min)) * float64(intervalConfig.min)
			}
		} else {
			minRunAt = float64(min) + 1
		}

		tStart := time.Date(year, month, day, hour, int(minRunAt), 0, 0, time.Local)

		if intervalConfig.min > 1 {
			return gocron.Every(uint64(intervalConfig.min)).Minutes().From(&tStart).Loc(impl.DateTime.GetNow().Location()).Do(task)
		} else {
			return gocron.Every(1).Minute().From(&tStart).Loc(impl.DateTime.GetNow().Location()).Do(task)
		}
	} else if rt == EveryHour {
		if intervalConfig.hour == 0 {
			return errors.New("invalid param")
		}

		if runAtConfig != nil && runAtConfig.hour != 0 {
			if hour >= runAtConfig.hour {
				day += 1
			}
			hour = runAtConfig.hour
		}
		if runAtConfig != nil {
			if min >= runAtConfig.min {
				hour += 1
			}
			min = runAtConfig.min
		}

		tStart := time.Date(year, month, day, hour, min, 0, 0, time.Local)

		return gocron.Every(uint64(intervalConfig.hour)).Hours().From(&tStart).Loc(impl.DateTime.GetNow().Location()).Do(task)
	} else if rt == EveryDay {
		if intervalConfig.day == 0 || runAtConfig == nil {
			return errors.New("invalid param")
		}

		return gocron.Every(uint64(intervalConfig.day)).Days().At(fmt.Sprintf("%02d:%02d", runAtConfig.hour, runAtConfig.min)).Loc(impl.DateTime.GetNow().Location()).Do(task)
	} else {
		return errors.New("invalid run type")
	}
}
