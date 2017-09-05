package main

import (
	"fmt"
	"sync/atomic"
	"time"

	scheduler "github.com/singchia/go-scheduler"
)

func main() {
	sch := scheduler.NewScheduler()
	sch.Interval = time.Millisecond * 100

	//sch.SetDefaultHandler(SchedulerHandler)
	sch.SetMonitor(SchedulerMonitor)
	sch.SetMaxRate(0.95)
	sch.SetMaxGoroutines(5000)
	sch.StartSchedule()

	var val int64
	for i := 0; i < 100*10000*10; i++ {
		sch.PublishRequest(&scheduler.Request{Data: val, Handler: SchedulerHandler})
		atomic.AddInt64(&val, 1)
		//time.Sleep(time.Microsecond * 10)
	}
	time.Sleep(time.Second * 10)
	fmt.Printf("maxValue: %d\n", maxValue)
	sch.Close()
}

var maxValue int64 = 0

func SchedulerHandler(data interface{}) {
	val, ok := data.(int64)
	if ok {
		if val > maxValue {
			maxValue = val
		}
	}
}

func SchedulerMonitor(incomingReqsDiff, processedReqsDiff, diff, currentGotoutines int64) {
	fmt.Printf("%d, %d, %d, %d\n", incomingReqsDiff, processedReqsDiff, diff, currentGotoutines)
}
