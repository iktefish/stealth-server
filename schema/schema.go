package schema

import "time"

type Location struct {
	Name string `json:"name" firestore:"name"`
	Open bool   `json:"open" firestore:"open"`
}

type Employee struct {
	Name            string `json:"name" firestore:"name"`
	ApplicationForm string `json:"applicationForm" firestore:"applicationForm"`
}

type WorkDay struct {
	Date                time.Time            `json:"date" firestore:"date,omitempty,serverTimestamp"`
	LocationAssignments []locationAssignment `json:"locationAssignments" firestore:"locationAssignments"`
}

type locationAssignment struct {
	Location  string     `json:"location" firestore:"location"`
	Assignees []assignee `json:"assignees" firestore:"assignees"`
}

type assignee struct {
	EmployeeId  string        `json:"employeeId" firestore:"employeeId"`
	HoursWorked time.Duration `json:"hoursWorked" firestore:"hoursWorked"`
}

type UnconfirmedAppointment struct {
	CustomerCellNum string    `json:"customerCellNum" firestore:"customerCellNum"`
	Job             string    `json:"job" firestore:"job"`
	Confirmed       bool      `json:"confirmed" firestore:"confirmed"`
	PreferredDate   time.Time `json:"preferredDate" firestore:"preferredDate,serverTimestamp"`
}

type ConfirmedAppointment struct {
	Date      time.Time `json:"date" firestore:"date,omitempty,serverTimestamp"`
	Job       string    `json:"job" firestore:"job"`
	Completed bool      `json:"completed" firestore:"completed"`
}
