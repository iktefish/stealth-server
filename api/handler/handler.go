package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iktefish/stealth-server/db"
	"github.com/iktefish/stealth-server/schema"
	"github.com/iktefish/stealth-server/serializer"
)

var seri serializer.Serializer

type Handler struct {
	db db.Database
}

func NewHandler(db db.Database) Handler {
	return Handler{
		db,
	}
}

/** @_ Degbugging utilities **/

func (h *Handler) DEBUG_GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, fmt.Errorf("UID is empty").Error(), http.StatusBadRequest)
		return
	}

	var employee schema.Employee
	var err, statusCode = seri.EmployeeToJson(w, employee)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed unmarshal").Error(), statusCode)
		return
	}
}

/** // **/

/** @_ Auth server **/

func (h *Handler) RegisterEmployee(w http.ResponseWriter, r *http.Request) {
	var employee schema.EmployeeRegisterForm
	err, statusCode := seri.JsonToEmployeeRegisterForm(r, &employee)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed marshal").Error(), statusCode)
		return
	}

	type NewlyCreatedEmployee struct {
		Uid string `json:"uid"`
	}
	err, statusCode, id := h.db.RegisterEmployee(employee)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to create Employee").Error(), statusCode)
		return
	}

	newlyCreatedEmployee := NewlyCreatedEmployee{Uid: id}
	log.Printf("newlyCreatedEmployee.uid~~> %v\n", newlyCreatedEmployee.Uid)
	err = json.NewEncoder(w).Encode(newlyCreatedEmployee)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed unmarshal").Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	var uid = r.URL.Query().Get("id")
	if uid == "" {
		http.Error(w, fmt.Errorf("UID is empty").Error(), http.StatusBadRequest)
		return
	}

	// var employee schema.Employee
	// var err, statusCode = seri.JsonToEmployee(r, &employee)
	// if err != nil {
	// 	http.Error(w, fmt.Errorf("Failed marshal").Error(), statusCode)
	// 	return
	// }

	var err, statusCode = h.db.RemoveEmployee(uid)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to remove Employee").Error(), statusCode)
		return
	}
}

/** // **/

/** @_ Clock in/out functionality **/

func (h *Handler) ClockIn(w http.ResponseWriter, r *http.Request) {
	var locId = r.URL.Query().Get("id")
	if locId == "" {
		http.Error(w, "Empty Location ID", http.StatusBadRequest)
	}

	h.db.ClockIn(locId)
	/* logic.CheckWorkDayOver(h.db, locId) */
	return
}

func (h *Handler) ClickOut(w http.ResponseWriter, r *http.Request) {
	var locId = r.URL.Query().Get("id")
	if locId == "" {
		http.Error(w, "Empty Location ID", http.StatusBadRequest)
	}

	h.db.ClockOut(locId)
	return
}

/** // **/

func (h *Handler) PostAppointment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) GetUnconfirmedAppointments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) GetConfirmedAppointments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) PutEmployeeToAppointment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) PutConfirmAppointment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) PutAssignEmployeeToDate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) PutCantMakeDate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) PutVolunteer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) GetLocationStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) PostForJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}
