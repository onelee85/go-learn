package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	now := time.Now()
	p(now)

	then := time.Date(2017, 10, 22, 15, 30, 59, 65138, time.UTC)
	p(then)

	p(then.Year())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))

	diff := now.Sub(then)
	p(diff)

	now = time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	p(now)
	p(secs)
	p(nanos)
	p(millis)
	p(now.Format(time.RFC3339))
	p(now.Format("3:04PM"))
	p(now.Format("2006-01-02 15:04:05.999999-07:00"))
}
