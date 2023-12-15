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

// func DateToday() string {
// 	var today string
// 	var year, month, day = time.Now().Date()
// 	today = strconv.Itoa(year) + " " + monthMap[month] + ", " + strconv.Itoa(day)
// 	return today
// }

func TodaysDateString() string {
	year, month, day := time.Now().Date()
	today := month.String() + " " + strconv.Itoa(day) + ", " + strconv.Itoa(year)
	return today
}

func DateObjFromString(dateString string) (time.Time, error) {
	dateTime, err := time.Parse("January 2, 2006", dateString)
	if err != nil {
		return time.Now(), err
	}

	return dateTime, nil
}

func TodaysDateObj() time.Time {
	dateTime, _ := time.Parse("January 2, 2006", TodaysDateString())
	return dateTime
}
