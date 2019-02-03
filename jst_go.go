package jst_go

import (
	"sync"
	"time"
)

var jstLoc time.Location
var jstOnce sync.Once

func JstNow() time.Time {
	jstOnce.Do(func() {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		jstLoc = *loc
	})
	return time.Now().In(&jstLoc)
}
