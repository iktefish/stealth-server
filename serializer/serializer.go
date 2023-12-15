package serializer

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/iktefish/stealth-server/schema"
)

type Serializer struct {
}

func (s *Serializer) JsonToLocation(r *http.Request, l *schema.Location) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(l)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) JsonToEmployee(r *http.Request, e *schema.Employee) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(e)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) JsonToEmployeeRegisterForm(r *http.Request, e *schema.EmployeeRegisterForm) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(e)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) JsonToWorkDay(r *http.Request, w *schema.WorkDay) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(w)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) JsonToUnconfirmedAppointment(r *http.Request, a *schema.UnconfirmedAppointment) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) JsonToConfirmedAppointment(r *http.Request, a *schema.ConfirmedAppointment) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) JsonToEmployeeAttendanceData(r *http.Request, a *schema.EmployeeAttendanceData) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(a)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) JsonToTent(r *http.Request, t *schema.Tent) (error, int) {
	var err = json.NewDecoder(r.Body).Decode(t)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) LocationToJson(w io.Writer, l schema.Location) (error, int) {
	var err = json.NewEncoder(w).Encode(l)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) EmployeeToJson(w io.Writer, e schema.Employee) (error, int) {
	var err = json.NewEncoder(w).Encode(e)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) EmployeeAttendanceDataToJson(w io.Writer, a schema.EmployeeAttendanceData) (error, int) {
	var err = json.NewEncoder(w).Encode(a)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) ListOfEmployeeAttendanceDataToJson(w io.Writer, a []schema.EmployeeAttendanceData) (error, int) {
	var err = json.NewEncoder(w).Encode(a)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) TentToJson(w io.Writer, t schema.Tent) (error, int) {
	var err = json.NewEncoder(w).Encode(t)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) WorkDayToJson(w io.Writer, d schema.WorkDay) (error, int) {
	var err = json.NewEncoder(w).Encode(d)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) UnconfirmedAppointmentToJson(w io.Writer, a schema.UnconfirmedAppointment) (error, int) {
	var err = json.NewEncoder(w).Encode(a)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (s *Serializer) ConfirmedAppointmentToJson(w io.Writer, a schema.ConfirmedAppointment) (error, int) {
	var err = json.NewEncoder(w).Encode(a)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
