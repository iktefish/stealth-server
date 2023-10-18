package utils

import (
	"strconv"
	"time"
)

var monthMap = map[time.Month]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func DateToday() string {
	var today string
	var year, month, day = time.Now().Date()
	today = strconv.Itoa(year) + " " + monthMap[month] + ", " + strconv.Itoa(day)
	return today
}
