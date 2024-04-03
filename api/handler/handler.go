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
	err, statusCode := h.db.DEBUG_GetEmployeeData(id, &employee)
	if err != nil {
		http.Error(w, fmt.Errorf("Internal server error").Error(), statusCode)
		return
	}

	err, statusCode = seri.EmployeeToJson(w, employee)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed unmarshal").Error(), statusCode)
		return
	}
}

// func (h *Handler) DEBUG_GetAllAttendanceData(w http.ResponseWriter, r *http.Request) {
// 	var attendenceDate schema.EmployeeAttendanceData
// 	err, statusCode := h.db.DEBUG_GetClockedInEmployees(&attendenceDate)
// 	if err != nil {
// 		http.Error(w, fmt.Errorf("Internal server error").Error(), statusCode)
// 		return
// 	}
//
// 	err, statusCode = seri.EmployeeAttendanceDataToJson(w, attendenceDate)
// 	if err != nil {
// 		http.Error(w, fmt.Errorf("Failed unmarshal").Error(), statusCode)
// 		return
// 	}
// }

func (h *Handler) DEBUG_GetClockedInEmployees(w http.ResponseWriter, r *http.Request) {
	var attendanceData []schema.EmployeeAttendanceData
	err, statusCode := h.db.DEBUG_GetClockedInEmployeesAttendanceData(&attendanceData)
	if err != nil {
		http.Error(w, fmt.Errorf("Internal server error").Error(), statusCode)
		return
	}

	err, statusCode = seri.ListOfEmployeeAttendanceDataToJson(w, attendanceData)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed unmarshal").Error(), statusCode)
		return
	}
}

func (h *Handler) DEBUG_GetClockedOutEmployees(w http.ResponseWriter, r *http.Request) {
}

/** // **/

/** @_ Auth server **/

type NewlyCreatedEmployee struct {
	Uid string `json:"uid"`
}

func (h *Handler) RegisterEmployee(w http.ResponseWriter, r *http.Request) {
	log.Print("RegisterEmployee 1")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
	// w.Header().Set("Content-Type", "application/json")

	var employee schema.EmployeeRegisterForm

	err, statusCode := seri.JsonToEmployeeRegisterForm(r, &employee)
	log.Printf("RegisterEmployee 2 : statusCode~~> %v err~~> \n", statusCode, err)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed marshal").Error(), statusCode)
		return
	}
	err, statusCode, id := h.db.RegisterEmployee(employee)
	log.Printf("RegisterEmployee 3 : statusCode~~> %v err~~> \n", statusCode, err)
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to create Employee").Error(), statusCode)
		return
	}

	newlyCreatedEmployee := NewlyCreatedEmployee{Uid: id}
	log.Printf("RegisterEmployee 4 : newlyCreatedEmployee.uid~~> %v\n", newlyCreatedEmployee.Uid)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(newlyCreatedEmployee)
	log.Printf("RegisterEmployee 5 : err~~> %v\n", err)
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

/** @_ Clock in/out **/

func (h *Handler) ClockIn(w http.ResponseWriter, r *http.Request) {
	var employeeId = r.URL.Query().Get("employeeId")
	var tentId = r.URL.Query().Get("tentId")

	log.Println("employeeId~~> ", employeeId)
	log.Println("tentId~~> ", tentId)

	if tentId == "" {
		http.Error(w, "Empty tent ID", http.StatusBadRequest)
		return
	}
	if employeeId == "" {
		http.Error(w, "Empty employee ID", http.StatusBadRequest)
		return
	}

	// var tent schema.Tent
	// var err, statusCode = seri.JsonToTent(r, &tent)
	// if err != nil {
	// 	http.Error(w, "Faulty JSON provided", statusCode)
	// 	return
	// }

	log.Println("Contecting DB")
	h.db.ClockIn(tentId, employeeId)
	// h.db.ClockIn(tentId, employeeId, tent)
	/* logic.CheckWorkDayOver(h.db, locId) */
	return
}

func (h *Handler) ClockOut(w http.ResponseWriter, r *http.Request) {
	// var locId = r.URL.Query().Get("id")
	// if locId == "" {
	// 	http.Error(w, "Empty tent ID", http.StatusBadRequest)
	// }

	// h.db.ClockOut(locId)

	var employeeId = r.URL.Query().Get("employeeId")
	var tentId = r.URL.Query().Get("tentId")
	if tentId == "" {
		http.Error(w, "Empty tent ID", http.StatusBadRequest)
		return
	}
	if employeeId == "" {
		http.Error(w, "Empty employee ID", http.StatusBadRequest)
		return
	}

	h.db.ClockOut(tentId, employeeId)
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
