package constants

import (
	"os"
	"time"
)

/* Absolute path of the 'service-key.json' file used to initiate Firebase SDK. */
var ServiceKeyPath = os.Getenv("PROJECT_ROOT") + "/service-key.json"

/* Database collection names */
const (
	EMPLOYEES                = "employee-docs"
	EX_EMPLOYEES             = "former-employees"
	LOCATIONS                = "tent-docs"
	UNCONFIRMED_APPOINTMENTS = "postponed-appointments"
	CONFIRMED_APPOINTMENTS   = "appointment-docs"
	WORKDAYS                 = "work-days"
	ATTENDANCE_DATA          = "attendance-records"
	TENT_STATE_RECORDS       = "tent-state-records"
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
