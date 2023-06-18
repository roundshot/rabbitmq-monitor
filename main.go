
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/barryz/rmqmonitor/cron"
	"github.com/barryz/rmqmonitor/falcon"
	"github.com/barryz/rmqmonitor/g"
	v "github.com/barryz/rmqmonitor/version"
	"github.com/barryz/rmqmonitor/witch"
)

func collect() {
	go metricCollector(g.Config().Interval)
}

func rotateQLog() {
	go cron.Start()
}

func witchLaunch() {
	go witch.Launch()
}

func metricCollector(sec int64) {
	t := time.NewTicker(time.Second * time.Duration(sec)).C