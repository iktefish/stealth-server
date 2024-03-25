package constants

import (
	"os"
	"time"
)

/* Absolute path of the 'service-key.json' file used to initiate Firebase SDK. */
var ServiceKeyPath = os.Getenv("PROJECT_ROOT") + "/service-key.json"

/* Database collection names */
const (
	EMPLOYEES                = "Employee"
	EX_EMPLOYEES             = "Former Employee"
	LOCATIONS                = "locations"
	UNCONFIRMED_APPOINTMENTS = "unconfirmed-appointments"
	CONFIRMED_APPOINTMENTS   = "confirmed-appointments"
	WORKDAYS                 = "work-days"
	ATTENDANCE_DATA          = "Attendance Record"
)

type WorkHours struct {
	Start struct {
		Hour   int
		Minute int
	}
	End struct {
		Hour   int
		Minute int
	}
}

var CloseTime = map[time.Weekday]WorkHours{
	0: {
		Start: struct {
			Hour   int
			Minute int
		}{
			10,
			45,
		},
		End: struct {
			Hour   int
			Minute int
		}{
			17,
			45,
		},
	},
	1: {
		Start: struct {
			Hour   int
			Minute int
		}{
			10,
			45,
		},
		End: struct {
			Hour   int
			Minute int
		}{
			18,
			15,
		},
	},
	2: {
		Start: struct {
			Hour   int
			Minute int
		}{
			10,
			45,
		},
		End: struct {
			Hour   int
			Minute int
		}{
			18,
			15,
		},
	},
	3: {
		Start: struct {
			Hour   int
			Minute int
		}{
			10,
			45,
		},
		End: struct {
			Hour   int
			Minute int
		}{
			18,
			15,
		},
	},
	4: {
		Start: struct {
			Hour   int
			Minute int
		}{
			10,
			45,
		},
		End: struct {
			Hour   int
			Minute int
		}{
			18,
			15,
		},
	},
	5: {
		Start: struct {
			Hour   int
			Minute int
		}{
			10,
			45,
		},
		End: struct {
			Hour   int
			Minute int
		}{
			17,
			45,
		},
	},
	6: {
		Start: struct {
			Hour   int
			Minute int
		}{
			10,
			45,
		},
		End: struct {
			Hour   int
			Minute int
		}{
			17,
			45,
		},
	},
}
