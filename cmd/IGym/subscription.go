package main

import "time"

type Subscription struct {
	Sub_id       int32
	UserSub_id   int32
	Startdate    time.Time
	durationdays int32
	enddate      time.Time
}
