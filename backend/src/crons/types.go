package crons

import (
	"context"
	"time"
)

type Cronjob struct {
	cancel context.CancelFunc
}

type cronjobT interface {
	Interval() time.Duration
	Run()
}
