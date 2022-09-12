package from_db

import (
	"context"

	"github.com/iktefish/stealth-wr-server/globals"
	"github.com/iktefish/stealth-wr-server/schema"
	"github.com/iktefish/stealth-wr-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Delete all documents in all collections. First it the names of all collections
from a database, and deleted every documents in each collection using an empty
filter.
WARNING: DONOT USE ON MAIN DATABASE CLUSTER.
*/
func Delete_All(client *mongo.Client, ctx context.Context) int64 {

	DummyDb := client.Database("DummyDb")
	allCols, err := DummyDb.ListCollectionNames(
		ctx,
		bson.D{},
	)
	utils.Handle_Error(err)

	var deletedCount int64
	for _, colName := range allCols {
		col := DummyDb.Collection(colName)
		deletedCol, err := col.DeleteMany(ctx, bson.M{})
		utils.Handle_Error(err)
		deletedCount += deletedCol.DeletedCount
	}

	return deletedCount

}

/*
Fetch all documents from "episodes" collection.
*/
func Fetch_All(client *mongo.Client, ctx context.Context) []primitive.M {

	DummyDb := client.Database("DummyDb")
	episodesCol := DummyDb.Collection("episodes")

	cur, err := episodesCol.Find(ctx, bson.M{})
	utils.Handle_Error(err)

	// It isn't wise to load an entire collection, possibly gigabytes in size, into memory
	/* var episodes []bson.M
	if err = cur.All(ctx, &episodes); err != nil {
		utils.Handle_Error(err)
	} */

	episodes := []primitive.M{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var episode bson.M
		err = cur.Decode(&episode)
		utils.Handle_Error(err)
		episodes = append(episodes, episode)
	}

	return episodes

}

/*
Fetch and sort all documents from "episodes" collection.
*/
func Fetch_All_And_Sort(client *mongo.Client, ctx context.Context,
) []primitive.M {

	DummyDb := client.Database("DummyDb")
	episodesCol := DummyDb.Collection("episodes")

	opts := options.Find()
	opts.SetSort(bson.D{{"rating", -1}}) // -1 is for descending order, 1 is ascending order

	cur, err := episodesCol.Find(ctx, bson.D{
		{"rating", bson.D{
			{"$gt", 4.0},
		}},
	}, opts)
	utils.Handle_Error(err)

	var episodesSorted []bson.M
	err = cur.All(ctx, &episodesSorted)
	utils.Handle_Error(err)

	return episodesSorted

}

/* ------------------------------***********************------------------------------ */
/* ------------------------------* P R O D U C T I O N *------------------------------ */
/* ------------------------------***********************------------------------------ */

/*
Take and email string and the clear password from client and query databse
to authenticate user login.
*/
func Auth_Success(
	client *mongo.Client,
	ctx context.Context,
	email string,
	clearPassword string,
) bool {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_CUSTOMERS)
	cur, err := col.Find(
		ctx,
		bson.M{"email": email},
	)

	var result []schema.Customer
	err = cur.All(ctx, &result)
	utils.Handle_Error(err)

	return utils.Verify_Hash(
		[]byte(result[0].Hashed_Password),
		clearPassword,
	)

}

/*
Fetch all `Location` documents from "locations" collection.
*/
func Fetch_Locations(client *mongo.Client, ctx context.Context,
) []schema.Location {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_LOCATIONS)
	cur, err := col.Find(ctx, bson.M{})
	utils.Handle_Error(err)

	locations := []schema.Location{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var location schema.Location
		err = cur.Decode(&location)
		utils.Handle_Error(err)
		locations = append(locations, location)
	}

	return locations

}

/*
Fetch all `Employee` documents from "employees" collection.
*/
func Fetch_Employees(client *mongo.Client, ctx context.Context,
) []schema.Employee {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_EMPLOYEES)
	cur, err := col.Find(ctx, bson.M{})
	utils.Handle_Error(err)

	employees := []schema.Employee{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var employee schema.Employee
		err = cur.Decode(&employee)
		utils.Handle_Error(err)
		employees = append(employees, employee)
	}

	return employees

}

/*
Fetch all `Job` documents from "jobs" collection.
*/
func Fetch_Jobs(client *mongo.Client, ctx context.Context,
) []schema.Job {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_JOBS)
	cur, err := col.Find(ctx, bson.M{})
	utils.Handle_Error(err)

	jobs := []schema.Job{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var job schema.Job
		err = cur.Decode(&job)
		utils.Handle_Error(err)
		jobs = append(jobs, job)
	}

	return jobs

}

/*
Fetch all `Service` documents from "services" collection.
*/
func Fetch_Services(client *mongo.Client, ctx context.Context,
) []schema.Service {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_SERVICES)
	cur, err := col.Find(ctx, bson.M{})
	utils.Handle_Error(err)

	services := []schema.Service{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var service schema.Service
		err = cur.Decode(&service)
		utils.Handle_Error(err)
		services = append(services, service)
	}

	return services

}

/*
Fetch all `ComVehicle` documents from "vehicles" collection.
*/
func Fetch_ComVehicles(client *mongo.Client, ctx context.Context,
) []schema.ComVehicle {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_VEHICLES)
	cur, err := col.Find(ctx, bson.M{})
	utils.Handle_Error(err)

	comVehicles := []schema.ComVehicle{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var comVehicle schema.ComVehicle
		err = cur.Decode(&comVehicle)
		utils.Handle_Error(err)
		comVehicles = append(comVehicles, comVehicle)
	}

	return comVehicles

}

/*
Fetch the `LocStats` of a `Location`. `id` for `LocStats` are referenced in
the `Location.Loc_Stats` field.
*/
func Fetch_LocStats(client *mongo.Client, ctx context.Context, id primitive.ObjectID,
) schema.LocStats {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_LOC_STATS)
	cur, err := col.Find(ctx, bson.M{
		"_id": id,
	})
	utils.Handle_Error(err)

	defer cur.Close(ctx)
	locStats := schema.LocStats{}
	err = cur.All(ctx, &locStats)
	utils.Handle_Error(err)

	return locStats

}

/*
Update `Loc_Investment` by scanning the `Loc_Rent` (from `Location`) field and
the `Hourly_Pay` (from `Employee`) field for investment.
TODO: For `Loc_Revenue` we want to intercept company email address to get the
reciepts of sales compute revenue.
*/
// func Create_LocStats(client *mongo.Client, ctx context.Context) {
//
// 	db := client.Database(globals.DB)
// 	col := db.Collection()
//
// }
