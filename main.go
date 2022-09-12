package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iktefish/stealth-wr-server/api"
	"github.com/iktefish/stealth-wr-server/dummy"
	"github.com/iktefish/stealth-wr-server/from_db"
	"github.com/iktefish/stealth-wr-server/to_db"
	"github.com/iktefish/stealth-wr-server/utils"
)

func main() {

	/* ------------------------------***********************------------------------------ */
	/* ------------------------------*    C O N N E C T    *------------------------------ */
	/* ------------------------------***********************------------------------------ */

	/*
	   Connect to database and make sure to detach connection.
	*/
	client, ctx, cancel := utils.Connect_Db()
	defer client.Disconnect(ctx)
	defer cancel()

	/* ------------------------------***********************------------------------------ */
	/* ------------------------------*   N O T   P R O D   *------------------------------ */
	/* ------------------------------***********************------------------------------ */

	/*
	   Fetch all.
	*/
	fmt.Println("\nFetch all episodes:")
	for _, episode := range from_db.Fetch_All(client, ctx) {
		fmt.Println("\tepisode: ", episode)
		fmt.Println("\tepisode['ep_num']: ", episode["ep_num"])
	}

	/*
		Insert one doc.
	*/
	fmt.Println("\nInsert_One_Doc() ~~> ", to_db.Insert_One_Doc(
		client,
		ctx,
	))

	/*
	   Insert many docs.
	*/
	fmt.Println("\nInsert_Many_Docs() ~~> ", to_db.Insert_Many_Docs(client, ctx))

	/*
	   Fetch all and sort.
	*/
	fmt.Println("\nFetch all episodes and sort:")
	for _, episode := range from_db.Fetch_All_And_Sort(client, ctx) {
		fmt.Println("\tepisode: ", episode)
		fmt.Println("\tepisode['ep_num']: ", episode["ep_num"])
	}

	/*
		Update one doc.
	*/
	fmt.Println("\nUpdated document: ", to_db.Update_Doc(
		client,
		ctx,
		"62bc86825e403311de401c84",
	))

	/*
		Update multiple documents where rating is > 4.0, and add "good_quality" tag.
	*/
	fmt.Printf("\nUpdate %v documents\n", to_db.Update_Many_Docs(
		client,
		ctx,
	))

	/*
	   Insert from native Go data structure.
	*/
	fmt.Println("\nInsert_From_Go_Struct() ~~> ", to_db.Insert_From_Go_Struct(
		client, ctx,
	))

	/*
	   Delete all documents in all collections.
	*/
	// fmt.Println("Deleted count: ", from_db.Delete_All(
	// 	client,
	// 	ctx,
	// ))

	/* ------------------------------***********************------------------------------ */
	/* ------------------------------*  T E S T   P R O D  *------------------------------ */
	/* ------------------------------***********************------------------------------ */

	/*
	   Hash a clear password.
	*/
	clear_pass := "123"
	hashed_pass := utils.Hash_Password(clear_pass)
	fmt.Println("\nClear password ~~> ", clear_pass)
	fmt.Println("Hashed password ~~>", hashed_pass)

	/*
	   Verify the aforementioned hash.
	*/
	verification := utils.Verify_Hash(hashed_pass, clear_pass)
	fmt.Println("\nVerification success ~~> ", verification)

	/*
	   Insert `Customer`.
	*/
	clearPassword := "12345"
	customer := dummy.Customer(clearPassword)
	fmt.Println("\nInsert_Customer() ~~> ", to_db.Insert_Customer(
		client,
		ctx,
		customer,
	))

	/*
	   Insert `Employee`.
	*/
	employee := dummy.Employee()
	fmt.Println("\nInsert_Employee() ~~> ", to_db.Insert_Employee(
		client,
		ctx,
		employee,
	))

	/*
	   Insert `ComVehicle`.
	*/
	comVehicle := dummy.ComVehicle()
	fmt.Println("\nInsert_ComVehicle() ~~> ", to_db.Insert_ComVehicle(
		client,
		ctx,
		comVehicle,
	))

	/*
	   Insert `Service`.
	*/
	service := dummy.Service()
	fmt.Println("\nInsert_Service() ~~> ", to_db.Insert_Service(
		client,
		ctx,
		service,
	))

	/*
	   Authenticate user login.
	*/
	email := "hb@bil.hil"
	success := from_db.Auth_Success(
		client,
		ctx,
		email,
		clearPassword,
	)
	fmt.Println("\nAuth_Success() ~~> ", success)

	/*
	   Check validity of input email.
	*/
	fmt.Println("\nEmail_Is_Valid() ~~> ", utils.Email_Is_Valid(
		"coffeemaker500@tea.spoon",
	))

	/*
	   Fetch all `Location`.
	*/
	fmt.Println("\nFetch_Locations() ~~> ", from_db.Fetch_Locations(
		client,
		ctx,
	))

	/*
	   Fetch all `Employee`.
	*/
	fmt.Println("\nFetch_Employees() ~~> ", from_db.Fetch_Employees(
		client,
		ctx,
	))

	/*
	   Fetch all `Job`.
	*/
	fmt.Println("\nFetch_Jobs() ~~> ", from_db.Fetch_Jobs(
		client,
		ctx,
	))

	/*
	   Fetch all `Service`.
	*/
	fmt.Println("\nFetch_Services() ~~> ", from_db.Fetch_Services(
		client,
		ctx,
	))

	/*
	   Fetch all `ComVehicle`.
	*/
	fmt.Println("\nFetch_ComVehicles() ~~> ", from_db.Fetch_ComVehicles(
		client,
		ctx,
	))

	/* ------------------------------***********************------------------------------ */
	/* ------------------------------*        A P I        *------------------------------ */
	/* ------------------------------***********************------------------------------ */

	/*
	   Start API router in main loop.
	*/
	router := gin.Default()

	router.GET("/locations", api.Get_Locations)
	router.GET("/employees", api.Get_Employees)
	router.GET("/jobs", api.Get_Jobs)
	router.GET("/services", api.Get_Services)
	router.GET("/vehicles", api.Get_ComVehicles)

	router.POST("/locations", api.Post_Location)

	router.Run("localhost:5000")

}
