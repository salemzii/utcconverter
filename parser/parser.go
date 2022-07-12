package parser

import (
	"strconv"
	"strings"
	"time"
)

//Receives string representation of time.time value or datetime string and returns time.Time equivalent
func Parse(val string) *time.Time {
	var mnth time.Month
	spltTimeVal := strings.Split(val, " ")
	date := spltTimeVal[0]
	tme := spltTimeVal[1]
	spltDateVal := strings.Split(date, "-") //2022-07-12
	spltTmeVal := strings.Split(tme, ":")   //02:10:36.101106135

	year, _ := strconv.Atoi(spltDateVal[0])
	month, _ := strconv.Atoi(spltDateVal[1])
	day, _ := strconv.Atoi(spltDateVal[2])

	hour, _ := strconv.Atoi(spltTmeVal[0])
	min, _ := strconv.Atoi(spltTmeVal[1])
	sec, _ := strconv.Atoi(spltTmeVal[2])
	nsec, _ := strconv.Atoi(strings.Split(spltTmeVal[2], ".")[1])

	switch month {
	case 1:
		mnth = time.January
	case 2:
		mnth = time.February
	case 3:
		mnth = time.March
	case 4:
		mnth = time.April
	case 5:
		mnth = time.May
	case 6:
		mnth = time.June
	case 7:
		mnth = time.July
	case 8:
		mnth = time.August
	case 9:
		mnth = time.September
	case 10:
		mnth = time.October
	case 11:
		mnth = time.November
	case 12:
		mnth = time.December
	}

	ParsedTimeVal := time.Date(year, mnth, day, hour, min, sec, nsec, time.UTC)

	return &ParsedTimeVal
}
