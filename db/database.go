package db

import (
	"context"
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

/** @_ Auth server **/

func (r *Database) RegisterEmployee(employee schema.Employee) (error, int) {
	var employeeToCreate = (&auth.UserToCreate{}).
		Email(employee.Email).
		EmailVerified(false).
		PhoneNumber(employee.PhoneNumber).
		Password(employee.Password).
		DisplayName(employee.DisplayName).
		PhotoURL(employee.PhotoUrl).
		Disabled(false)

	/* var userRecord, err = r.auth.GetUserByEmail(context.Background(), employee.Email) */
	var userRecord, err = r.auth.CreateUser(context.Background(), employeeToCreate)
	if err != nil {
		log.Printf("err~~> %v\n", err)
		return err, http.StatusInternalServerError
	}

	log.Printf("userRecord~~> %v\n", userRecord)

	var id = userRecord.UID
	err, statusCode := r.CloneEmployeeDataToFirestore(id, employee)
	if err != nil {
		log.Printf("err~~> %v\n", err)
		return err, statusCode
	}

	return nil, 0
}

func (r *Database) CloneEmployeeDataToFirestore(id string, employee schema.Employee) (error, int) {
	var ctx = context.Background()
	var employeeCollection = r.client.Collection(constants.EMPLOYEES)
	var docRef, err = employeeCollection.Doc(id).Set(ctx, employee)
	if err != nil {
		log.Printf("err~~> %s\n", err)
		return err, http.StatusInternalServerError
	}

	log.Printf("INS: docRef~~> %s\n", docRef)

	return nil, 0
}

func (r *Database) RemoveEmployee(uid string, employee schema.Employee) (error, int) {
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

	r.MarkEmployeeRemoved(uid)

	log.Printf("DEL: uid %v", uid)
	log.Printf("MARKED: uid %v as INACTIVE", uid)

	return nil, 0
}

func (r *Database) MarkEmployeeRemoved(uid string) (error, int) {
	var confirmedAppointmentsCollection = r.client.Collection(constants.CONFIRMED_APPOINTMENTS)
	var docRef, err = confirmedAppointmentsCollection.Doc(uid).Update(context.Background(), []firestore.Update{
		{
			Path: "isEmployeeActive",
			Value: struct {
				isActive    bool
				dateRemoved time.Time
			}{
				isActive:    false,
				dateRemoved: time.Now(),
			},
		},
	})
	if err != nil {
		log.Printf("err~~> %s\n", err)
		return err, http.StatusInternalServerError
	}

	log.Printf("UPD: docRef~~> %s\n", docRef)

	return nil, 0
}

/** // **/

func (r *Database) PutCheckIn(locId string) (error, int) {
	var ctx = context.Background()
	var results, err = r.client.Collection(constants.LOCATIONS).Doc(locId).Update(ctx, []firestore.Update{
		{
			Path:  "open",
			Value: true,
		},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	log.Printf("Log: results~~> %v\n", results)
	return nil, http.StatusOK
}

func (r *Database) PutCheckOut(locId string) (error, int) {
	var ctx = context.Background()
	var results, err = r.client.Collection(constants.LOCATIONS).Doc(locId).Update(ctx, []firestore.Update{
		{
			Path:  "open",
			Value: false,
		},
	})
	if err != nil {
		return err, http.StatusInternalServerError
	}

	log.Printf("Log: results~~> %v\n", results)
	return nil, http.StatusOK
}

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
	/*

	   Compute current day-month-year
	   Query to check if a document with this day-month-year is present
	           If not present make a document with that date
	   Get the locationId and employeeId
	   Check if employeeId is present in the given Location

	*/
	var ctx = context.Background()
	var today = utils.DateToday()
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
