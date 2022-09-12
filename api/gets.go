/*
Module to handle GET and POST requests.
*/
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iktefish/stealth-wr-server/from_db"
	"github.com/iktefish/stealth-wr-server/utils"
)

/*
Fetch and serve list of all location.
*/
func Get_Locations(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	gctx.IndentedJSON(http.StatusOK, from_db.Fetch_Locations(client, ctx))

}

/*
Fetch and serve list of all employess.
*/
func Get_Employees(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	gctx.IndentedJSON(http.StatusOK, from_db.Fetch_Employees(client, ctx))

}

/*
Fetch and serve list of all jobs.
*/
func Get_Jobs(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	gctx.IndentedJSON(http.StatusOK, from_db.Fetch_Jobs(client, ctx))

}

/*
Fetch and serve list of all services.
*/
func Get_Services(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	gctx.IndentedJSON(http.StatusOK, from_db.Fetch_Services(client, ctx))

}

/*
Fetch and serve list of all company vehicles.
*/
func Get_ComVehicles(gctx *gin.Context) {

	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	gctx.IndentedJSON(http.StatusOK, from_db.Fetch_ComVehicles(client, ctx))

}
