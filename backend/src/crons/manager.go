package crons

import (
	"context"
	"time"
)

func Start(task cronjobT) *Cronjob {
	ctx, cancel := context.WithCancel(context.Background())
	c := Cronjob{
		cancel: cancel,
	}

	go cronHandler(task, ctx)

	return &c
}

func (c *Cronjob) Stop() {
	c.cancel()
}

func cronHandler(task cronjobT, ctx context.Context) {
	for {
		task.Run()
		time.Sleep(task.Interval())
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
