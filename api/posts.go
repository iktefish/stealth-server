/*
Module to handle GET and POST requests.
*/
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iktefish/stealth-wr-server/schema"
	"github.com/iktefish/stealth-wr-server/to_db"
	"github.com/iktefish/stealth-wr-server/utils"
)

/*
Post a location object as json and insert into database.
*/
func Post_Location(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	location := schema.Location{}
	err := gctx.BindJSON(&location)
	if err != nil {
		utils.Handle_Error(err)
	}

	to_db.Insert_Location(client, ctx, location)
	gctx.IndentedJSON(http.StatusCreated, location)

}

/*
Post an employee object as json and insert into database.
*/
func Post_Employee(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	employee := schema.Employee{}
	err := gctx.BindJSON(&employee)
	if err != nil {
		utils.Handle_Error(err)
	}

	to_db.Insert_Employee(client, ctx, employee)
	gctx.IndentedJSON(http.StatusCreated, employee)

}

/*
Post a job object as json and insert into database.
*/
func Post_Job(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	job := schema.Job{}
	err := gctx.BindJSON(&job)
	if err != nil {
		utils.Handle_Error(err)
	}

	to_db.Insert_Job(client, ctx, job)
	gctx.IndentedJSON(http.StatusCreated, job)

}

/*
Post a service object as json and insert into database.
*/
func Post_Service(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	service := schema.Service{}
	err := gctx.BindJSON(&service)
	if err != nil {
		utils.Handle_Error(err)
	}

	to_db.Insert_Service(client, ctx, service)
	gctx.IndentedJSON(http.StatusCreated, service)

}

/*
Post a company vehicle object as json and insert into database.
*/
func Post_ComVehicle(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	comVehicle := schema.ComVehicle{}
	err := gctx.BindJSON(&comVehicle)
	if err != nil {
		utils.Handle_Error(err)
	}

	to_db.Insert_ComVehicle(client, ctx, comVehicle)
	gctx.IndentedJSON(http.StatusCreated, comVehicle)

}
