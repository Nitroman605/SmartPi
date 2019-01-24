/*
    Copyright (C) Jens Ramhorst
	  This file is part of SmartPi.
    SmartPi is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    SmartPi is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with SmartPi.  If not, see <http://www.gnu.org/licenses/>.
    Diese Datei ist Teil von SmartPi.
    SmartPi ist Freie Software: Sie können es unter den Bedingungen
    der GNU General Public License, wie von der Free Software Foundation,
    Version 3 der Lizenz oder (nach Ihrer Wahl) jeder späteren
    veröffentlichten Version, weiterverbreiten und/oder modifizieren.
    SmartPi wird in der Hoffnung, dass es nützlich sein wird, aber
    OHNE JEDE GEWÄHRLEISTUNG, bereitgestellt; sogar ohne die implizite
    Gewährleistung der MARKTFÄHIGKEIT oder EIGNUNG FÜR EINEN BESTIMMTEN ZWECK.
    Siehe die GNU General Public License für weitere Details.
    Sie sollten eine Kopie der GNU General Public License zusammen mit diesem
    Programm erhalten haben. Wenn nicht, siehe <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Nitroman605/SmartPi/src/smartpi"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/exp/io/i2c"

	"github.com/fsnotify/fsnotify"

	//import the Paho Go MQTT library

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/version"
)

func makeReadoutAccumulator() (r smartpi.ReadoutAccumulator) {
	r.Current = make(smartpi.Readings)
	r.Voltage = make(smartpi.Readings)
	r.ActiveWatts = make(smartpi.Readings)
	r.CosPhi = make(smartpi.Readings)
	r.Frequency = make(smartpi.Readings)
	r.WattHoursConsumed = make(smartpi.Readings)
	r.WattHoursProduced = make(smartpi.Readings)
	return r
}

func makeReadout() (r smartpi.ADE7878Readout) {
	r.Current = make(smartpi.Readings)
	r.Voltage = make(smartpi.Readings)
	r.ActiveWatts = make(smartpi.Readings)
	r.CosPhi = make(smartpi.Readings)
	r.Frequency = make(smartpi.Readings)
	r.ApparentPower = make(smartpi.Readings)
	r.ReactivePower = make(smartpi.Readings)
	r.PowerFactor = make(smartpi.Readings)
	r.ActiveEnergy = make(smartpi.Readings)
	return r
}

func pollSmartPi(config *smartpi.Config, device *i2c.Device) {
	//var mqttclient MQTT.Client
	var p smartpi.Phase

	/*if config.MQTTenabled {
		mqttclient = newMQTTClient(config)
	}*/

	//i := 0

	tick := time.Tick(1 * time.Millisecond)
	now := time.Now()
	currentHour := now.Hour()
	file, err := os.OpenFile(strings.Join([]string{now.Format("2006-01-02h15"), ".csv"}, ""), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{"timestamp", "Current #1", "Current #2", "Current #3"})
	i := 0
	for {
		readouts := makeReadout()
		// Restart the accumulator loop every 60 seconds.
		/*if i > 59 {
			i = 0
		}*/

		//startTime := time.Now()
		//smartpi.ReadPhase(device, config, smartpi.PhaseN, &readouts)
		// Update readouts and the accumlator.
		for _, p = range smartpi.MainPhases {
			smartpi.ReadCurrent(device, config, p)

		}
		readouts.Current[1] = smartpi.ReadCurrent(device, config, smartpi.MainPhases[0])
		readouts.Current[2] = smartpi.ReadCurrent(device, config, smartpi.MainPhases[1])
		readouts.Current[3] = smartpi.ReadCurrent(device, config, smartpi.MainPhases[2])
		// Update metrics endpoint.
		updatePrometheusMetrics(&readouts)
		//publishMQTTReadouts(config, mqttclient, &readouts)

		//go sendPost(config, &readouts)
		/*
			delay := time.Since(startTime) - (1000 * time.Millisecond)
			if int64(delay) > 0 {
				log.Errorf("Readout delayed: %s", delay)
			}
		*/
		now = time.Now()
		if currentHour != now.Hour() {
			csvWriter.Flush()
			currentHour = now.Minute()
			file, err = os.OpenFile(strings.Join([]string{now.Format("2006-01-02h15"), ".csv"}, ""), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			csvWriter = csv.NewWriter(file)
			csvWriter.Write([]string{"timestamp", "Current #1", "Current #2", "Current #3"})
		}
		csvWriter.Write([]string{now.Format("2006-01-02 15:04:05.000"), strconv.FormatFloat(readouts.Current[1], 'E', -1, 64), strconv.FormatFloat(readouts.Current[2], 'E', -1, 64), strconv.FormatFloat(readouts.Current[3], 'E', -1, 64)})
		if i == 1000 {
			csvWriter.Flush()
			i = 0
		}
		<-tick
		i++
	}
}

func configWatcher(config *smartpi.Config) {
	log.Debug("Start SmartPi watcher")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	log.Debug("init done 1")
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					config.ReadParameterFromFile()
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	log.Debug("init done 2")
	err = watcher.Add("./smartpi")
	if err != nil {
		log.Fatal(err)
	}
	<-done
	log.Debug("init done 3")
}

func sendPost(c *smartpi.Config, values *smartpi.ADE7878Readout) {
	/*t1 := fmt.Sprintf("C1 : %v , C2 : %v , C3 : %v \n", values.Current[1], values.Current[2], values.Current[3])
	fmt.Println(t1)
	t2 := fmt.Sprintf("V1 : %v , V2 : %v , V3 : %v \n", values.Voltage[1], values.Voltage[2], values.Voltage[3])
	fmt.Println(t2)
	t3 := fmt.Sprintf("AW1 : %v , AW2 : %v , AW3 : %v \n", values.ActiveWatts[1], values.ActiveWatts[2], values.ActiveWatts[3])
	fmt.Println(t3)
	fmt.Printf("~~~~~~~~~~~~~~~~~~~~~~~~~~~ \n")*/
	url := c.Address
	jsonReadout, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReadout))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Debug(err)
		return
	}
	defer resp.Body.Close()
}
func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	prometheus.MustRegister(currentMetric)
	prometheus.MustRegister(voltageMetric)
	prometheus.MustRegister(activePowerMetirc)
	prometheus.MustRegister(cosphiMetric)
	prometheus.MustRegister(frequencyMetric)
	prometheus.MustRegister(apparentPowerMetric)
	prometheus.MustRegister(reactivePowerMetric)
	prometheus.MustRegister(powerFactorMetric)
	prometheus.MustRegister(version.NewCollector("smartpi"))
}

var appVersion = "No Version Provided"

func main() {
	config := smartpi.NewConfig()

	go configWatcher(config)

	version := flag.Bool("v", false, "prints current version information")
	flag.Parse()
	if *version {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	log.SetLevel(config.LogLevel)

	listenAddress := config.MetricsListenAddress

	log.Debug("Start SmartPi readout")

	device, _ := smartpi.InitADE7878(config)

	go pollSmartPi(config, device)

	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head><title>SmartPi Readout Metrics Server</title></head>
            <body>
            <h1>SmartPi Readout Metrics Server</h1>
            <p><a href="/metrics">Metrics</a></p>
            </body>
            </html>`))
	})

	log.Debugf("Listening on %s", listenAddress)
	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		panic(fmt.Errorf("Error starting HTTP server: %s", err))
	}
}
