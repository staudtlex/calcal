// Copyright (C) 2022  Alexander Staudt
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	lc "staudtlex.de/libcalendar"
)

// Target calendars names (internal)
var calendars = []string{
	"gregorian",
	"iso",
	"julian",
	"islamic",
	"hebrew",
	"mayanLongCount",
	"mayanHaab",
	"mayanTzolkin",
	"french",
	"oldHinduSolar",
	"oldHinduLunar",
}

// Map of internal calendar to external (fmt.Print-related) calendar names
var calendarNames = map[string]string{
	"gregorian":      "Gregorian",
	"iso":            "ISO",
	"julian":         "Julian",
	"islamic":        "Islamic",
	"hebrew":         "Hebrew",
	"mayanLongCount": "Mayan Long Count",
	"mayanHaab":      "Mayan Haab",
	"mayanTzolkin":   "Mayan Tzolkin",
	"french":         "French Revolutionary",
	"oldHinduSolar":  "Old Hindu Solar",
	"oldHinduLunar":  "Old Hindu Lunar",
}

// documentCalendars pretty prints the information of the map `calendarNames`
// to a string
func documentCalendars(names map[string]string, calendars []string) string {
	s := ""
	for _, v := range calendars {
		s = s + fmt.Sprintf("%-12s\t%s %s\n", v, names[v], "calendar")
	}
	return s[0:(len(s) - 1)]
}

// member checks returns true if a string is contained in a given slice, and
// false otherwise.
func member(e string, s []string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// unique returns the slice of unique strings contained in a given slice.
func unique(s []string) []string {
	helperMap := make(map[string]bool, len(s))
	uniqueStrings := make([]string, 0, len(s))
	for _, v := range s {
		if _, ok := helperMap[v]; !ok {
			helperMap[v] = true
			uniqueStrings = append(uniqueStrings, v)
		}
	}
	return uniqueStrings
}

// parseDate parses a given date string as a Gregorian date and returns a
// GregorianDate and an error value. If the date string is invalid, parseDate
// returns a non-nil error value and an GregorianDate set to {0, 0, 0}. Dates
// before the common era may be indicated with a leading "-".
func parseDate(s string) (lc.GregorianDate, error) {
	// deal with years before the common era (BCE)
	commonEra := 1.0
	if len(s) > 0 && s[0] == '-' {
		s = s[1:]
		commonEra = -1.0
	}
	// parse the rest of the date string
	parsed, err := time.Parse("2006-01-02", s)
	var date lc.GregorianDate
	if err != nil {
		return date, errors.New("invalid date string")
	}
	date = lc.GregorianDate{
		Year:  float64(parsed.Year()) * commonEra,
		Month: float64(parsed.Month()),
		Day:   float64(parsed.Day())}
	return date, nil
}

func main() {
	// command line arguments
	date_s := flag.String("d", "",
		fmt.Sprintf("%s\n%s (%s)",
			"date (format: yyyy-mm-dd). When omitted, '-d' defaults to the",
			"current date", time.Now().Format("2006-01-02")))
	list := flag.String("c", "",
		fmt.Sprintf("%s\n%s\n%s",
			"comma-separated list of calendars. Currently, calcal supports",
			fmt.Sprintf("%-12s\t%s", "all", "all calendars listed below (default)"),
			documentCalendars(calendarNames, calendars)))

	// parse arguments
	flag.Parse()

	// check if -c, -d have arguments
	var inputList []string
	if len(*list) > 0 { // if -c has an argument
		inputList = unique(strings.Split(*list, ","))
	} else { // otherwise, provide default argument
		inputList = append(inputList, "all")
	}

	var date string
	if *date_s != "" { // if -d has an argument
		date = *date_s
	} else { // otherwise, set default argument
		date = time.Now().Format("2006-01-02")
	}

	// ensure that Gregorian date is valid
	parsedDate, err := parseDate(date)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid date. Specify date as yyyy-mm-dd.\n")
		os.Exit(1)
	}

	// get absolute (fixed) date from Gregorian date
	d := lc.AbsoluteFromGregorian(parsedDate)

	// keep valid (non-duplicated) calendar names
	requiredCalendars := make([]string, 0, len(inputList))
	for _, v := range inputList {
		if member(v, calendars) {
			requiredCalendars = append(requiredCalendars, v)
		}
	}

	// display program limitations
	fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n",
		"Please note:",
		"- for dates before 1 January 1 (Gregorian), calcal may return incorrect",
		"  ISO and Julian dates",
		"- for dates before the beginning of the Islamic calendar (16 July 622",
		"  Julian), calcal returns an invalid Islamic date",
		"- for dates before the beginning of the French Revolutionary calendar",
		"  calendar (22 September 1792), calcal returns an invalid French date",
		"- Old Hindu (solar and lunar) calendar dates may be off by one day",
	)

	// compute and return dates
	if member("all", inputList) {
		for _, v := range calendars {
			s := lc.FromAbsolute(d, v)
			fmt.Printf("%-16s\t%-32s\n", calendarNames[v], s)
		}
	} else {
		for _, v := range requiredCalendars {
			s := lc.FromAbsolute(d, v)
			fmt.Printf("%-16s\t%-32s\n", calendarNames[v], s)
		}
	}
}
