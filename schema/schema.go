package schema

import "time"

type Location struct {
	Name string `json:"name" firestore:"name"`
	Open bool   `json:"open" firestore:"open"`
}

type Employee struct {
	Name               string    `json:"name" firestore:"name"`
	Email              string    `json:"email" firestore:"email"`
	Password           string    `json:"password,omitempty" firestore:"password,omitempty"`
	PhotoUrl           string    `json:"photoUrl" firestore:"photoUrl"`
	PhoneNumber        string    `json:"phoneNumber" firestore:"phoneNumber"`
	DisplayName        string    `json:"displayName" firestore:"displayName"`
	Address            string    `json:"address" firestore:"address"`
	Birthday           time.Time `json:"birthday" firestore:"birthday"`
	Nationality        string    `json:"nationality" firestore:"nationality"`
	ApplicationFormUrl string    `json:"applicationFormUrl" firestore:"applicationFormUrl"`
}

type WorkDay struct {
	/* HERE: [DONE] ... change time.Time to be timestamps instead of date-times */
	Date                string               `json:"date" firestore:"date,omitempty,serverTimestamp"`
	LocationAssignments []LocationAssignment `json:"locationAssignments" firestore:"locationAssignments"`
}

type LocationAssignment struct {
	Location  string     `json:"location" firestore:"location"`
	Assignees []Assignee `json:"assignees" firestore:"assignees"`
}

type Assignee struct {
	EmployeeId  string        `json:"employeeId" firestore:"employeeId"`
	HoursWorked time.Duration `json:"hoursWorked" firestore:"hoursWorked"`
}

type UnconfirmedAppointment struct {
	PostDate        int64  `json:"postDate" firestore:"postDate,serverTimestamp"`
	CustomerCellNum string `json:"customerCellNum" firestore:"customerCellNum"`
	Job             string `json:"job" firestore:"job"`
	Confirmed       bool   `json:"confirmed" firestore:"confirmed"`
	PreferredDate   int64  `json:"preferredDate" firestore:"preferredDate,serverTimestamp"`
}

type ConfirmedAppointment struct {
	Date            int64  `json:"date" firestore:"date,omitempty,serverTimestamp"`
	CustomerCellNum string `json:"customerCellNum" firestore:"customerCellNum"`
	Assigned        bool   `json:"assigned" firestore:"assigned"`
	AssignedTo      string `json:"assignedTo" firestore:"assignedTo"`
	Job             string `json:"job" firestore:"job"`
	Completed       bool   `json:"completed" firestore:"completed"`
}
