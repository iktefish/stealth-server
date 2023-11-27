package schema

import "time"

type Location struct {
	Name string `json:"name" firestore:"name"`
	Open bool   `json:"open" firestore:"open"`
}

// type Employee struct {
// 	Name               string    `json:"name" firestore:"name"`
// 	Email              string    `json:"email" firestore:"email"`
// 	Password           string    `json:"password,omitempty" firestore:"password,omitempty"`
// 	PhotoUrl           string    `json:"photoUrl" firestore:"photoUrl"`
// 	PhoneNumber        string    `json:"phoneNumber" firestore:"phoneNumber"`
// 	DisplayName        string    `json:"displayName" firestore:"displayName"`
// 	Address            string    `json:"address" firestore:"address"`
// 	Birthday           time.Time `json:"birthday" firestore:"birthday"`
// 	Nationality        string    `json:"nationality" firestore:"nationality"`
// 	Position           string    `json:"position" firestore:"position"`
// 	Gender             string    `json:"gender" firestore:"gender"`
// 	Eligible           string    `json:"eligible" firestore:"eligible"`
// 	Car                string    `json:"car" firestore:"car"`
// 	Offense            string    `json:"offense" firestore:"offense"`
// 	ShiftsMayAug       string    `json:"shiftsMayAug" firestore:"shiftsMayAug"`
// 	ShiftsApr          string    `json:"shiftsApr" firestore:"shiftsApr"`
// 	ShiftsSepOct       string    `json:"shiftsSepOct" firestore:"shiftsSepOct"`
// 	HourlyWage         string    `json:"hourlyWage" firestore:"hourlyWage"`
// 	NidBlobLink        string    `json:"nidBlobLink" firestore:"nidBlobLink"`
// 	CvBlobLink         string    `json:"cvBlobUrl" firestore:"cvBlobUrl"`
// 	DateCreated        time.Time `json:"dateCreated" firestore:"dateCreated"`
// 	ApplicationFormUrl string    `json:"applicationFormUrl" firestore:"applicationFormUrl"`
// }

type EmployeeRegisterForm struct {
	FirstName       string      `json:"firstName" firestore:"firstName"`
	LastName        string      `json:"lastName" firestore:"lastName"`
	Phone           string      `json:"phone" firestore:"phone"`
	Position        string      `json:"position" firestore:"position"`
	Gender          string      `json:"gender" firestore:"gender"`
	Email           string      `json:"email" firestore:"email"`
	Password        string      `json:"password" firestore:"password"`
	Birthday        string      `json:"bday" firestore:"bday"`
	Address         string      `json:"address" firestore:"address"`
	Eligibility     bool        `json:"eligibility" firestore:"eligibility"`
	Car             bool        `json:"car" firestore:"car"`
	CriminalOffense bool        `json:"criminalOffense" firestore:"criminalOffense"`
	ShiftsMayAug    string      `json:"shiftsMayAug" firestore:"shiftsMayAug"`
	ShiftsApr       string      `json:"shiftsApr" firestore:"shiftsApr"`
	ShiftsSepOct    string      `json:"shiftsSepOct" firestore:"shiftsSepOct"`
	HourlyWage      string      `json:"hourlyWage" firestore:"hourlyWage"`
	Nationality     string      `json:"nationality" firestore:"nationality"`
	NidBlobLink     LinkAndPath `json:"nidBlobLink" firestore:"nidBlobLink"`
	CvBlobLink      LinkAndPath `json:"cvBlobLink" firestore:"cvBlobLink"`
	ProfilePhotoUrl LinkAndPath `json:"profilePhotoUrl" firestore:"profilePhotoUrl"`
	Date            time.Time   `json:"date" firestore:"date"`
}

type Employee struct {
	FirstName       string      `json:"firstName" firestore:"firstName"`
	LastName        string      `json:"lastName" firestore:"lastName"`
	Phone           string      `json:"phone" firestore:"phone"`
	Position        string      `json:"position" firestore:"position"`
	Gender          string      `json:"gender" firestore:"gender"`
	Email           string      `json:"email" firestore:"email"`
	Birthday        string      `json:"bday" firestore:"bday"`
	Address         string      `json:"address" firestore:"address"`
	Eligibility     bool        `json:"eligibility" firestore:"eligibility"`
	Car             bool        `json:"car" firestore:"car"`
	CriminalOffense bool        `json:"criminalOffense" firestore:"criminalOffense"`
	ShiftsMayAug    string      `json:"shiftsMayAug" firestore:"shiftsMayAug"`
	ShiftsApr       string      `json:"shiftsApr" firestore:"shiftsApr"`
	ShiftsSepOct    string      `json:"shiftsSepOct" firestore:"shiftsSepOct"`
	HourlyWage      string      `json:"hourlyWage" firestore:"hourlyWage"`
	Nationality     string      `json:"nationality" firestore:"nationality"`
	NidBlobLink     LinkAndPath `json:"nidBlobLink" firestore:"nidBlobLink"`
	CvBlobLink      LinkAndPath `json:"cvBlobLink" firestore:"cvBlobLink"`
	ProfilePhotoUrl LinkAndPath `json:"profilePhotoUrl" firestore:"profilePhotoUrl"`
	Date            time.Time   `json:"date" firestore:"date"`
}

type LinkAndPath struct {
	FilePath string `json:"filePath" firestore:"filePath"`
	Url      string `json:"url" firestore:"url"`
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
