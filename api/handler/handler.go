package handler

import (
	"fmt"
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

func (h *Handler) RegisterEmployee(w http.ResponseWriter, r *http.Request) {
	var employee schema.Employee
	var err, statusCode = seri.JsonToEmployee(r, &employee)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed marshal").Error(), statusCode)
		return
	}

	err, statusCode = h.db.RegisterEmployee(employee)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to create Employee").Error(), statusCode)
		return
	}
}

func (h *Handler) RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented")
	return
}

func (h *Handler) PutCheckIn(w http.ResponseWriter, r *http.Request) {
	var locId = r.URL.Query().Get("id")
	if locId == "" {
		http.Error(w, "Empty Location ID", http.StatusBadRequest)
	}

	h.db.PutCheckIn(locId)
	/* logic.CheckWorkDayOver(h.db, locId) */
	return
}

func (h *Handler) PutCheckOut(w http.ResponseWriter, r *http.Request) {
	var locId = r.URL.Query().Get("id")
	if locId == "" {
		http.Error(w, "Empty Location ID", http.StatusBadRequest)
	}

	h.db.PutCheckOut(locId)
	return
}

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
