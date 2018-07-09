// Database Exporter

package main

import (
	"fmt"
	"time"

	"github.com/Nitroman605/SmartPi/src/smartpi"

	log "github.com/Sirupsen/logrus"
)

func updateSQLiteDatabase(c *smartpi.Config, data smartpi.ReadoutAccumulator, consumedWattHourBalanced float64, producedWattHourBalanced float64) {
	t := time.Now()

	logLine := "## SQLITE Database Update ##"
	logLine += fmt.Sprintf(t.Format(" 2006-01-02 15:04:05 "))
	log.Info(logLine)

}
