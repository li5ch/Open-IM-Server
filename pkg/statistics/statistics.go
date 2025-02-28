package statistics

import (
	"Open_IM/pkg/common/log"
	"time"
)

type Statistics struct {
	AllCount   *uint64
	ModuleName string
	PrintArgs  string
	SleepTime  int
}

func (s *Statistics) output() {
	var intervalCount uint64
	t := time.NewTicker(time.Duration(s.SleepTime) * time.Second)
	defer t.Stop()
	var sum uint64
	for {
		sum = *s.AllCount
		select {
		case <-t.C:
		}
		if *s.AllCount-sum <= 0 {
			intervalCount = 0
		} else {
			intervalCount = *s.AllCount - sum
		}
		log.NewWarn("", " system stat ", s.ModuleName, s.PrintArgs, intervalCount, "total:", *s.AllCount)
	}
}

func NewStatistics(allCount *uint64, moduleName, printArgs string, sleepTime int) *Statistics {
	p := &Statistics{AllCount: allCount, ModuleName: moduleName, SleepTime: sleepTime, PrintArgs: printArgs}
	go p.output()
	return p
}
