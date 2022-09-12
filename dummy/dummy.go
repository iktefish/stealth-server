/*
Generate gummy data for exploration/debugging.
*/
package dummy

import (
	"time"

	"github.com/iktefish/stealth-wr-server/schema"
	"github.com/iktefish/stealth-wr-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Return dummy `LocStats`
*/
func LocStats() schema.LocStats {

	return schema.LocStats{
		Id:             primitive.NewObjectID(),
		Loc_Revenue:    10000,
		Loc_Investment: 13000,
		Loc_Due:        0,
	}

}

/*
Return dummy `Job`
*/
func Job() schema.Job {

	return schema.Job{
		Title:  "Tire mechanic",
		Salary: 30,
	}

}

/*
Return dummy `Employee`
*/
func Employee() schema.Employee {

	return schema.Employee{
		Email:       "coffeemaker500@tea.spoon",
		First_Name:  "Tea",
		Last_Name:   "Coffeeborn",
		Role:        Job(),
		Date_Hired:  primitive.NewDateTimeFromTime(time.Now()),
		Pt_Ft:       false,
		Hourly_Pay:  30,
		Payment_Due: 0,
	}

}

/*
Return dummy `ComVechicle`
*/
func ComVehicle() schema.ComVehicle {

	return schema.ComVehicle{
		License_Plate: "FACTS-JS-SUX",
	}

}

/*
Return dummy `Location`
*/
func Location() schema.Location {

	return schema.Location{
		Label:        "Northeast Road",
		Loc_Rent:     700,
		Loc_Stats:    LocStats().Id,
		Employees:    []schema.Employee{Employee()},
		Com_Vehicles: []schema.ComVehicle{ComVehicle()},
	}

}

/*
Return dummy `Customer`
*/
func Customer(clearPassword string) schema.Customer {

	return schema.Customer{
		Email:           "hb@bil.hil",
		First_Name:      "Hilly",
		Last_Name:       "Bill",
		Hashed_Password: string(utils.Hash_Password(clearPassword)),
	}

}

/*
Return dummy `Service`
*/
func Service() schema.Service {

	return schema.Service{
		Name:      "Small chip repair",
		Price:     10,
		Frequency: 384,
	}

}
