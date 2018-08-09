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
/*
File: csv.go
Description: create csv-file
*/

package smartpi

import (
	"encoding/xml"
	"time"
)

const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	xmlHeader = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

type tXmlValue struct {
	XMLName      xml.Name `xml:"valueset"`
	Date         string   `json:"date" xml:"date"`
	Current_1    float32  `json:"current_1" xml:"current_1"`
	Current_2    float32  `json:"current_2" xml:"current_2"`
	Current_3    float32  `json:"current_3" xml:"current_3"`
	Current_4    float32  `json:"current_4" xml:"current_4"`
	Voltage_1    float32  `json:"voltage_1" xml:"voltage_1"`
	Voltage_2    float32  `json:"voltage_2" xml:"voltage_2"`
	Voltage_3    float32  `json:"voltage_3" xml:"voltage_3"`
	Power_1      float32  `json:"power_1" xml:"power_1"`
	Power_2      float32  `json:"power_2" xml:"power_2"`
	Power_3      float32  `json:"power_3" xml:"power_3"`
	Cosphi_1     float32  `json:"cosphi_1" xml:"cosphi_1"`
	Cosphi_2     float32  `json:"cosphi_2" xml:"cosphi_2"`
	Cosphi_3     float32  `json:"cosphi_3" xml:"cosphi_3"`
	Frequency_1  float32  `json:"frequency_1" xml:"frequency_1"`
	Frequency_2  float32  `json:"frequency_2" xml:"frequency_2"`
	Frequency_3  float32  `json:"frequency_3" xml:"frequency_3"`
	Energy_pos_1 float32  `json:"energyPos_1" xml:"energyPos_1"`
	Energy_pos_2 float32  `json:"energyPos_2" xml:"energyPos_2"`
	Energy_pos_3 float32  `json:"energyPos_3" xml:"energyPos_3"`
	Energy_neg_1 float32  `json:"energyNeg_1" xml:"energyNeg_1"`
	Energy_neg_2 float32  `json:"energyNeg_2" xml:"energyNeg_2"`
	Energy_neg_3 float32  `json:"energyNeg_3" xml:"energyNeg_3"`
}

func CreateXML(start time.Time, end time.Time) string {

	var valueSeries []tXmlValue

	type valuelist []tXmlValue

	type dataset struct {
		valuelist
	}

	// type serie []tChartSerie

	if xmlstring, err := xml.MarshalIndent(dataset{valueSeries}, "", "    "); err == nil {
		xmlstring = []byte(xml.Header + string(xmlstring))
		return string(xmlstring)
	} else {
		return ""
	}

}
