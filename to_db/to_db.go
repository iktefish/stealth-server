package to_db

import (
	"context"

	"github.com/iktefish/stealth-wr-server/dummy"
	"github.com/iktefish/stealth-wr-server/globals"
	"github.com/iktefish/stealth-wr-server/schema"
	"github.com/iktefish/stealth-wr-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Insert one document in "episodes" collection.
*/
func Insert_One_Doc(client *mongo.Client, ctx context.Context,
) interface{} {

	DummyDb := client.Database("DummyDb")
	episodesCol := DummyDb.Collection("episodes")

	getEpisode, err := episodesCol.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Lord of Flies"},
		{Key: "ep_num", Value: 25},
		{"rating", 5}, // This is shorthand for { Key: "rating", Value: 5 }
		{"tags", bson.A{"thriller", "psychological", "gore"}},
	})
	utils.Handle_Error(err)

	return getEpisode.InsertedID

}

/*
Insert many documents in "episodes" collection.
*/
func Insert_Many_Docs(client *mongo.Client, ctx context.Context,
) interface{} {

	DummyDb := client.Database("DummyDb")
	episodesCol := DummyDb.Collection("episodes")

	getEpisode, err := episodesCol.InsertMany(ctx, []interface{}{ // We will be providing more than 1 `bson.D`, hence the slice.
		bson.D{
			{"name", "Escape from Paradise"},
			{"ep_num", 8},
			{"rating", 4.3},
		},
		bson.D{
			{"name", "Unspeakable Tongue"},
			{"ep_num", 9},
			{"rating", 4.6},
		},
		bson.D{
			{"name", "Cruel Compassion"},
			{"ep_num", 10},
			{"rating", 4.9},
		},
	})
	utils.Handle_Error(err)

	return getEpisode.InsertedIDs

}

/*
Insert in database from Go data structure.
*/
func Insert_From_Go_Struct(client *mongo.Client, ctx context.Context,
) interface{} {

	DummyDb := client.Database("DummyDb")
	locationCol := DummyDb.Collection("locations")

	result, err := locationCol.InsertOne(ctx, dummy.Location())
	utils.Handle_Error(err)

	return result.InsertedID

}

/*
Update document with provided "_id" in "episodes" collection.
*/
func Update_Doc(client *mongo.Client, ctx context.Context, id_string string,
) interface{} {

	DummyDb := client.Database("DummyDb")
	episodesCol := DummyDb.Collection("episodes")

	id, _ := primitive.ObjectIDFromHex(id_string)
	result, err := episodesCol.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"rating", 4.95}}},
		},
	)
	utils.Handle_Error(err)

	return result.UpsertedID

}

/*
Update "name" of documents, where "rating" is > 4.5, to "Money maker".
WARNING: Doesn't seem to work!
*/
func Update_Many_Docs(client *mongo.Client, ctx context.Context,
) int64 {

	DummyDb := client.Database("DummyDb")
	episodesCol := DummyDb.Collection("episodes")

	result, err := episodesCol.UpdateMany(
		ctx,
		bson.D{{"rating", bson.D{{"$gt", 4.5}}}},
		bson.D{
			{"$set", bson.D{{"name", "Money maker"}}},
		},
	)
	utils.Handle_Error(err)

	return result.ModifiedCount

}

/* ------------------------------***********************------------------------------ */
/* ------------------------------* P R O D U C T I O N *------------------------------ */
/* ------------------------------***********************------------------------------ */

/*
Push `Location` to database.
*/
func Insert_Location(client *mongo.Client, ctx context.Context, location schema.Location,
) interface{} {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_LOCATIONS)

	result, err := col.InsertOne(
		ctx,
		location,
	)
	utils.Handle_Error(err)

	return result.InsertedID

}

/*
Push `Employee` to database.
*/
func Insert_Employee(client *mongo.Client, ctx context.Context, employee schema.Employee,
) interface{} {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_EMPLOYEES)

	result, err := col.InsertOne(
		ctx,
		employee,
	)
	utils.Handle_Error(err)

	return result.InsertedID

}

/*
Push `Vehicle` to database.
*/
func Insert_ComVehicle(client *mongo.Client, ctx context.Context, comVehicle schema.ComVehicle,
) interface{} {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_VEHICLES)

	result, err := col.InsertOne(
		ctx,
		comVehicle,
	)
	utils.Handle_Error(err)

	return result.InsertedID

}

/*
Push `Customer` to database.
*/
func Insert_Customer(client *mongo.Client, ctx context.Context, customer schema.Customer,
) interface{} {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_CUSTOMERS)

	result, err := col.InsertOne(
		ctx,
		customer,
	)
	utils.Handle_Error(err)

	return result.InsertedID

}

/*
Push `Service` to database.
*/
func Insert_Service(client *mongo.Client, ctx context.Context, service schema.Service,
) interface{} {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_SERVICES)

	result, err := col.InsertOne(
		ctx,
		service,
	)
	utils.Handle_Error(err)

	return result.InsertedID

}

/*
Push `Job` to database.
*/
func Insert_Job(client *mongo.Client, ctx context.Context, job schema.Job,
) interface{} {

	db := client.Database(globals.DB)
	col := db.Collection(globals.COL_JOBS)

	result, err := col.InsertOne(
		ctx,
		job,
	)
	utils.Handle_Error(err)

	return result.InsertedID

}
