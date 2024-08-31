package db

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/iktefish/stealth-server/constants"
	"github.com/iktefish/stealth-server/schema"
	"github.com/iktefish/stealth-server/utils"
	"google.golang.org/api/iterator"
)

type Database struct {
	app    *firebase.App
	client *firestore.Client
	auth   *auth.Client
}

func NewDatabase(app *firebase.App, client *firestore.Client, auth *auth.Client) Database {
	return Database{
		app,
		client,
		auth,
	}
}

/** @_ Debugging utilities **/

func (r *Database) DEBUG_GetEmployeeData(uid string, e *schema.Employee) (error, int) {
	var ctx = context.Background()
	var employeeCollection = r.client.Collection(constants.EMPLOYEES)
	var docRef, err = employeeCollection.Doc(uid).Get(ctx)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = docRef.DataTo(e)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	log.Printf("Log: docRef.Exists()~~> %v\n", docRef.Exists())
	return nil, http.StatusOK
}

func (r *Database) DEBUG_GetAllAttendanceData(attendenceData *schema.EmployeeAttendanceData) (error, int) {
	var ctx = context.Background()
	var attendanceDataCollection = r.client.Collection(constants.ATTENDANCE_DATA)
	var iterable = attendanceDataCollection.Documents(ctx)
	for {
		docRef, err := iterable.Next()
		if err != iterator.Done {
			break
		}
		if err != nil {
			return err, http.StatusInternalServerError
		}

		fmt.Println(docRef.Data())
		log.Printf("Log: docRef.Exists() ~~> %v\n", docRef.Exists())
	}

	return nil, http.StatusOK
}

func (r *Database) DEBUG_GetClockedInEmployeesAttendanceData(attendenceDataList *[]schema.EmployeeAttendanceData) (error, int) {
	ctx := context.Background()
	attendanceDataCol := r.client.Collection(constants.ATTENDANCE_DATA)
	dateObj := utils.TodaysDateObj()
	docRefIter := attendanceDataCol.Where("clockedIn", "==", true).Where("date", "==", dateObj).Documents(ctx)
	var counter int
	for {
		doc, err := docRefIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			return err, http.StatusInternalServerError
		}

		var atd schema.EmployeeAttendanceData
		err = doc.DataTo(&atd)
		if err != nil {
			return err, http.StatusInternalServerError
		}

		(*attendenceDataList) = append((*attendenceDataList), atd)
		counter++
	}

	log.Printf("DEBUG_GetClockedInEmployeesAttendanceData: Retrieved~~> %v docs\n", counter)
	return nil, http.StatusOK
}

func (r *Database) DEBUG_GetClockedOutEmployeesAttendanceData(attendenceDataList *[]schema.EmployeeAttendanceData) (error, int) {
	ctx := context.Background()
	attendanceDataCol := r.client.Collection(constants.ATTENDANCE_DATA)
	dateObj := utils.TodaysDateObj()
	docRefIter := attendanceDataCol.Where("clockedOut", "==", true).Where("date", "==", dateObj).Documents(ctx)
	var counter int
	for {
		doc, err := docRefIter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			return err, http.StatusInternalServerError
		}

		var atd schema.EmployeeAttendanceData
		err = doc.DataTo(&atd)
		if err != nil {
			return err, http.StatusInternalServerError
		}

		(*attendenceDataList) = append((*attendenceDataList), atd)
		counter++
	}

	log.Printf("DEBUG_GetClockedOutEmployeesAttendanceData: Retrieved~~> %v docs\n", counter)
	return nil, http.StatusOK
}

/** // **/

/** @_ Auth server **/

func (r *Database) RegisterEmployee(employee schema.EmployeeRegisterForm) (error, int, string) {
	var employeeToCreate = (&auth.UserToCreate{}).
		Email(employee.Email).
		EmailVerified(false).
		PhoneNumber(employee.Phone).
		Password(employee.Password).
		DisplayName(employee.FirstName + " " + employee.LastName).
		PhotoURL(employee.ProfilePhotoUrl.Url).
		Disabled(false)

	/* var userRecord, err = r.auth.GetUserByEmail(context.Background(), employee.Email) */
	var userRecord, err = r.auth.CreateUser(context.Background(), employeeToCreate)
	if err != nil {
		log.Printf("err~~> %v\n", err)
		return err, http.StatusInternalServerError, ""
	}

	log.Printf("userRecord~~> %v\n", userRecord.UID)

	// var id = userRecord.UID
	// err, statusCode := r.CloneEmployeeDataToFirestore(id, employee)
	// if err != nil {
	// 	log.Printf("err~~> %v\n", err)
	// 	return err, statusCode
	// }

	return nil, http.StatusOK, userRecord.UID
}

/*
Cloning employee account data into firestore will be done via the client SDK, therefore
the following code is redundant.
*/
// func (r *Database) CloneEmployeeDataToFirestore(id string, employeeRegisterForm schema.EmployeeRegisterForm) (error, int) {
// 	var ctx = context.Background()
// 	var employeeCollection = r.client.Collection(constants.EMPLOYEES)
// 	var employee = schema.Employee{
// 		FirstName:       employeeRegisterForm.FirstName,
// 		LastName:        employeeRegisterForm.LastName,
// 		Phone:           employeeRegisterForm.Phone,
// 		Position:        employeeRegisterForm.Position,
// 		Gender:          employeeRegisterForm.Gender,
// 		Email:           employeeRegisterForm.Email,
// 		Birthday:        employeeRegisterForm.Birthday,
// 		Address:         employeeRegisterForm.Address,
// 		Eligibility:     employeeRegisterForm.Eligibility,
// 		Car:             employeeRegisterForm.Car,
// 		CriminalOffense: employeeRegisterForm.CriminalOffense,
// 		ShiftsMayAug:    employeeRegisterForm.ShiftsMayAug,
// 		ShiftsApr:       employeeRegisterForm.ShiftsApr,
// 		ShiftsSepOct:    employeeRegisterForm.ShiftsSepOct,
// 		HourlyWage:      employeeRegisterForm.HourlyWage,
// 		Nationality:     employeeRegisterForm.Nationality,
// 		NidBlobLink:     employeeRegisterForm.NidBlobLink,
// 		CvBlobLink:      employeeRegisterForm.CvBlobLink,
// 		ProfilePhotoUrl: employeeRegisterForm.ProfilePhotoUrl,
// 		Date:            time.Time{},
// 	}
//
// 	var docRef, err = employeeCollection.Doc(id).Set(ctx, employee)
// 	if err != nil {
// 		log.Printf("err~~> %s\n", err)
// 		return err, http.StatusInternalServerError
// 	}
//
// 	log.Printf("INS: docRef~~> %s\n", docRef)
//
// 	return nil, 0
// }

func (r *Database) RemoveEmployee(uid string) (error, int) {
	var employeeToUpdate = (&auth.UserToUpdate{}).Disabled(true)
	var userRecord, err = r.auth.UpdateUser(context.Background(), uid, employeeToUpdate)
	if err != nil {
		log.Printf("err~~> %v\n", err)
		return err, http.StatusInternalServerError
	}

	log.Printf("userRecord~~> %v\n", userRecord)

	err = r.auth.DeleteUser(context.Background(), uid)
	if err != nil {
		log.Printf("err~~> %s\n", err)
		return err, http.StatusInternalServerError
	}

	// r.MarkEmployeeRemoved(uid)

	log.Printf("DEL: uid %v", uid)
	log.Printf("MARKED: uid %v as INACTIVE", uid)

	return nil, http.StatusOK
}

/*
We will not be having an `isEmployeeActive` field, rather we will have a seperate collection
that will hold all 'ex-employees'. CRUD operations on said collection will be done via the
client SDK, thus the following function is redundant.
*/
// func (r *Database) MarkEmployeeRemoved(uid string) (error, int) {
// 	var confirmedAppointmentsCollection = r.client.Collection(constants.CONFIRMED_APPOINTMENTS)
// 	var docRef, err = confirmedAppointmentsCollection.Doc(uid).Update(context.Background(), []firestore.Update{
// 		{
// 			Path: "isEmployeeActive",
// 			Value: struct {
// 				isActive    bool
// 				dateRemoved time.Time
// 			}{
// 				isActive:    false,
// 				dateRemoved: time.Now(),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		log.Printf("err~~> %s\n", err)
// 		return err, http.StatusInternalServerError
// 	}
//
// 	log.Printf("UPD: docRef~~> %s\n", docRef)
//
// 	return nil, 0
// }

/** // **/

/** @_ Clock in/out functionality **/

func (r *Database) ClockIn(tentId string, employeeId string, hourlyWage string) (error, int) {
	log.Println("In DB 1")
	ctx := context.Background()
	attendanceDataCol := r.client.Collection(constants.ATTENDANCE_DATA)
	dateObj := utils.TodaysDateObj()

	docRefIter := attendanceDataCol.Where("employeeId", "==", employeeId).Where("date", "==", dateObj).Where("tentId", "==", tentId).Documents(ctx)
	_, err := docRefIter.Next()
	// if err == nil {
	// 	log.Println("In DB ERR")
	// 	return fmt.Errorf("Document already exists"), http.StatusBadRequest
	// }

	log.Println("In DB 2")
	var employee schema.Employee
	r.DEBUG_GetEmployeeData(employeeId, &employee)

	log.Println("In DB 3")

	attendanceData := schema.EmployeeAttendanceData{
		Date:         dateObj,
		ClockIn:      time.Now(),
		ClockedIn:    true,
		ClockedOut:   false,
		HoursWorked:  0,
		EmployeeID:   employeeId,
		EmployeeName: employee.FirstName + employee.LastName,
		HourlyWage:   hourlyWage,
		TentID:       tentId,
	}

	log.Printf("ClockIn: attendanceData~~> %v\n", attendanceData)

	ref, result, err := attendanceDataCol.Add(ctx, attendanceData)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	log.Printf("Log: Ref~~> %v\n", ref)
	log.Printf("Log: Result~~> %v\n", result)
	return nil, http.StatusOK
}

func (r *Database) ClockOut(tentId string, employeeId string) (error, int) {
	ctx := context.Background()
	attendanceDataCol := r.client.Collection(constants.ATTENDANCE_DATA)
	// locationsCol := r.client.Collection(constants.LOCATIONS)
	dateObj := utils.TodaysDateObj()

	docRefIter := attendanceDataCol.Where("employeeId", "==", employeeId).Where("date", "==", dateObj).Where("tentId", "==", tentId).Where("clockedIn", "==", true).Documents(ctx)
	doc, err := docRefIter.Next()
	if err != nil {
		if err == iterator.Done {
			return fmt.Errorf("No such document exists, please check if you have entered the current Tent ID or Employee ID"), http.StatusBadRequest
		}

		return err, http.StatusInternalServerError
	}

	var attendanceData schema.EmployeeAttendanceData
	err = doc.DataTo(&attendanceData)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	hoursWorked := time.Now().Hour() - attendanceData.ClockIn.Hour()
	results, err := attendanceDataCol.Doc(doc.Ref.ID).Update(ctx, []firestore.Update{
		{
			Path:  "clockedIn",
			Value: false,
		},
		{
			Path:  "clockedOut",
			Value: true,
		},
		{
			Path:  "clockOut",
			Value: time.Now(),
		},
		{
			Path:  "hoursWorked",
			Value: hoursWorked,
		},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	/** DONE **/
	//  DONE Check if document (query for clocked-in == true) for current day exists
	//  DONE Check if document (query for clocked-in == true) for provided employee ID exist
	//  DONE Check if document (query for clocked-in == true) for provided tent ID exist
	//  DONE (This is taken care of during clock-in procedure) Clock out only if it's clocked in

	log.Printf("Log: results~~> %v\n", results)
	return nil, http.StatusOK
}

/** // **/

/** @_ Update employee info (for Auth) **/

func (r *Database) UpdateEmployeeInfo(info schema.EmployeeInfoForAuthCredChange) (error, int) {
	var infoToUpdate = (&auth.UserToUpdate{}).
		Email(info.Email).
		Password(info.Password).
		DisplayName(info.DisplayName).
		PhoneNumber(info.Contact).
		PhotoURL(info.ProfilePictureUrl)

	var ctx = context.Background()
	var updatedRecord, err = r.auth.UpdateUser(ctx, info.Id, infoToUpdate)
	if err != nil {
		log.Printf("err~~> %v\n", err)
		return err, http.StatusInternalServerError
	}

	log.Printf("UpdateEmployeeInfo: updatedRecord~~> %v\n", updatedRecord.UID)
	return nil, http.StatusOK
}

/** // **/

/** @_ Toggling tent states **/

func (r *Database) SetTentStateOpen(tentId string, employeeId string) (error, int) {
	var tent schema.TentStateRecord
	r.GetTentData(tentId, &tent)

	/// Append employeeId to tent data.
	var employeeData schema.TentStateEmployeeRecord
	employeeData.Id = employeeId
	tent.ClockedEmployees = append(tent.ClockedEmployees, employeeData)

	// for _, v := range tent.ClockedEmployees {
	// 	if v.Id == employeeId {
	// 		tent.ClockedEmployees = append(tent.ClockedEmployees[:i], tent.ClockedEmployees[i+1:]...)
	// 	}
	// }

	var ctx = context.Background()

	/// Toggle "open" state to true.
	var tentCollection = r.client.Collection(constants.TENT_STATE_RECORDS)
	var _, err = tentCollection.Doc(tentId).Update(ctx, []firestore.Update{
		{
			Path:  "clockedEmployees",
			Value: tent.ClockedEmployees,
		},
		{
			Path:  "open",
			Value: true,
		},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (r *Database) SetTentStateClose(tentId string, employeeId string) (error, int) {
	var tent schema.TentStateRecord
	r.GetTentData(tentId, &tent)

	/// Remove employeeId from tent data.
	for i, v := range tent.ClockedEmployees {
		if v.Id == employeeId {
			tent.ClockedEmployees = append(tent.ClockedEmployees[:i], tent.ClockedEmployees[i+1:]...)
		}
	}

	var ctx = context.Background()

	if len(tent.ClockedEmployees) != 0 {
		/// Some employees are still clocked in, just update "clockedEmployees".
		var tentCollection = r.client.Collection(constants.TENT_STATE_RECORDS)
		var _, err = tentCollection.Doc(tentId).Update(ctx, []firestore.Update{
			{
				Path:  "clockedEmployees",
				Value: tent.ClockedEmployees,
			},
		})
		if err != nil {
			return err, http.StatusInternalServerError
		}

		return nil, http.StatusOK
	} else {
		/// No clocked in employees, so set "open" state to false.
		var tentCollection = r.client.Collection(constants.TENT_STATE_RECORDS)
		var _, err = tentCollection.Doc(tentId).Update(ctx, []firestore.Update{
			{
				Path:  "clockedEmployees",
				Value: tent.ClockedEmployees,
			},
			{
				Path:  "open",
				Value: false,
			},
		})
		if err != nil {
			return err, http.StatusInternalServerError
		}

		return nil, http.StatusOK
	}
}

func (r *Database) GetTentData(uid string, t *schema.TentStateRecord) (error, int) {
	var ctx = context.Background()

	var tentCollection = r.client.Collection(constants.TENT_STATE_RECORDS)
	var docRef, err = tentCollection.Doc(uid).Get(ctx)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	err = docRef.DataTo(*t)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	log.Printf("API Log for GetTentData: docRef.Exists()~~> %v\n", docRef.Exists())
	return nil, http.StatusOK
}

/** // **/

func (r *Database) PostAppointment(uap schema.UnconfirmedAppointment, cell int) (error, int) {
	var ctx = context.Background()
	var docRef, results, err = r.client.Collection(constants.UNCONFIRMED_APPOINTMENTS).Add(ctx, uap)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	log.Printf("Log: docRef~~> %v\tresults~~> %v\n", docRef, results)
	return nil, http.StatusOK
}

func (r *Database) GetUnconfirmedAppointments(uaps *[]schema.UnconfirmedAppointment) (error, int) {
	var ctx = context.Background()
	var iter = r.client.Collection(constants.UNCONFIRMED_APPOINTMENTS).OrderBy("postDate", firestore.Asc).Limit(25).Documents(ctx)
	var counter int
	for {
		var doc, err = iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			return err, http.StatusInternalServerError
		}

		var uap schema.UnconfirmedAppointment
		err = doc.DataTo(&uap)
		if err != nil {
			return err, http.StatusInternalServerError
		}

		(*uaps) = append((*uaps), uap)
		counter++
	}

	log.Printf("Log: Retrieved~~> %v docs\n", counter)
	return nil, http.StatusOK
}

func (r *Database) GetConfirmedAppointments(aps *[]schema.ConfirmedAppointment) (error, int) {
	var ctx = context.Background()
	var iter = r.client.Collection(constants.CONFIRMED_APPOINTMENTS).OrderBy("date", firestore.Asc).Limit(25).Documents(ctx)
	var counter int
	for {
		var doc, err = iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			return err, http.StatusInternalServerError
		}

		var ap schema.ConfirmedAppointment
		err = doc.DataTo(&ap)
		if err != nil {
			return err, http.StatusInternalServerError
		}

		(*aps) = append((*aps), ap)
		counter++
	}

	log.Printf("Log: Retrieved~~> %v docs\n", counter)
	return nil, http.StatusOK
}

func (r *Database) PutEmployeeToAppointment(eId string, apId string) (error, int) {
	var ctx = context.Background()
	var result, err = r.client.Collection(constants.CONFIRMED_APPOINTMENTS).Doc(apId).Update(ctx, []firestore.Update{
		{
			Path:  "assignedTo",
			Value: eId,
		},
		{
			Path:  "assigned",
			Value: true,
		},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	log.Printf("Log: result~~> %v\n", result)
	return nil, http.StatusOK
}

func (r *Database) PutConfirmAppointment(uapId string) (error, int) {
	var ctx = context.Background()
	var docSnap, err_1 = r.client.Collection(constants.UNCONFIRMED_APPOINTMENTS).Doc(uapId).Get(ctx)
	if err_1 != nil {
		return err_1, http.StatusInternalServerError
	}

	var uap schema.UnconfirmedAppointment
	err_1 = docSnap.DataTo(&uap)
	if err_1 != nil {
		return err_1, http.StatusInternalServerError
	}

	var ap = schema.ConfirmedAppointment{
		Date:            time.Now().Unix(),
		CustomerCellNum: uap.CustomerCellNum,
		Assigned:        false,
		AssignedTo:      "",
		Job:             uap.Job,
		Completed:       false,
	}

	var docRef, results, err_2 = r.client.Collection(constants.CONFIRMED_APPOINTMENTS).Add(ctx, ap)
	if err_2 != nil {
		return err_2, http.StatusInternalServerError
	}

	log.Printf("Log: docRef~~> %v\tresults~~> %v\n", docRef, results)
	return nil, http.StatusOK
}

func (r *Database) PutAssignEmployeeToDate(date int64) (error, int) {
	/** TODO **/
	// Compute current day-month-year
	// Query to check if a document with this day-month-year is present
	// If not present make a document with that date
	// Get the locationId and employeeId
	// Check if employeeId is present in the given Location

	var ctx = context.Background()
	var today = utils.TodaysDateString()
	var workDay = &schema.WorkDay{
		Date:                today,
		LocationAssignments: []schema.LocationAssignment{},
	}
	var iter = r.client.Collection(constants.WORKDAYS).Where("date", "==", today).Documents(ctx)
	var counter int
	for {
		var _, err = iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

			return err, http.StatusInternalServerError
		}

		counter++
	}

	if counter == 0 {
		var docRef, results, err = r.client.Collection(constants.WORKDAYS).Add(ctx, workDay)
		if err != nil {
			return err, http.StatusInternalServerError
		}

		log.Printf("Log: docRef~~> %v\tresults~~> %v\n", docRef, results)
	}

	// var docSnap, err_1 = r.client.Collection(constants.WorkDays)

	return nil, http.StatusOK
}

func (r *Database) PutCantMakeDate() (error, int) {
	return nil, http.StatusOK
}

func (r *Database) PutVolunteer() (error, int) {
	return nil, http.StatusOK
}

func (r *Database) GetLocationStatus() (error, int) {
	return nil, http.StatusOK
}

func (r *Database) PostForJob() (error, int) {
	return nil, http.StatusOK
}
