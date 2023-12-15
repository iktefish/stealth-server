package clockChecker

import (
	"log"
	"time"

	"github.com/iktefish/stealth-server/constants"
	"github.com/iktefish/stealth-server/db"
	"github.com/iktefish/stealth-server/schema"
)

/** @_ To be used in "cron-job" coroutine **/

func CheckerLoop(r *db.Database) {
	for {
		/** DONE: Fetch clocked in empoyees **/
		var attendanceDataList []schema.EmployeeAttendanceData

		err, statusCode := r.DEBUG_GetClockedInEmployeesAttendanceData(&attendanceDataList)
		if err != nil {
			log.Printf("[LOG] CheckerLoop: Failed to get attendanceDataList, error with code ~~> %v\n", statusCode)
			time.Sleep(time.Minute * 60)
			continue
		}

		/** DONE: Clock out the employees that are clocked in if past office hours **/
		weekday := time.Now().Weekday()
		todaysEndingHour := constants.CloseTime[weekday].End.Hour
		currentHour := time.Now().Hour()
		if currentHour >= todaysEndingHour {
			/** DONE: Mark them "Clocked Out" **/
			for _, attendanceData := range attendanceDataList {
				r.ClockOut(attendanceData.TentID, attendanceData.EmployeeID)
			}
		}

		time.Sleep(time.Minute * 20)
	}
}

/** // **/
