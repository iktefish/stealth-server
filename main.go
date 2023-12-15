package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iktefish/stealth-server/api/handler"
	clockChecker "github.com/iktefish/stealth-server/clock-checker"
	"github.com/iktefish/stealth-server/config"
	"github.com/iktefish/stealth-server/db"
	"github.com/iktefish/stealth-server/middleware"
)

func main() {
	var router = chi.NewRouter()
	const port string = ":8080"

	/** @_ Firebase SDK and clients setup **/

	var app, authClient, storeClient, err = config.NewSdkAndClients()
	if err != nil {
		log.Println("Fatal crash during setup")
		log.Fatalln("Error: ", err)
	}

	var db = db.NewDatabase(app, storeClient, authClient)
	var handler = handler.NewHandler(db)

	/** // **/

	log.Printf("app~~> %s\n", app)
	log.Printf("storeClient~~> %s\n", storeClient)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Stealth ...")
	})

	/** @_ Apply Auth middleware **/

	// router.Put("/check-in", middleware.AuthMiddleware(handler.ClockIn))
	// router.Put("/check-out", middleware.AuthMiddleware(handler.ClockOut))
	// router.Post("/appointment", middleware.AuthMiddleware(handler.PostAppointment))
	// router.Get("/confirmed-appointments", middleware.AuthMiddleware(handler.GetConfirmedAppointments))
	// router.Put("/employee-to-appointment", middleware.AuthMiddleware(handler.PutEmployeeToAppointment))
	// router.Put("/confirm-appointment", middleware.AuthMiddleware(handler.PutConfirmAppointment))
	// router.Put("/assign-employee-to-date", middleware.AuthMiddleware(handler.PutAssignEmployeeToDate))
	// router.Put("/cant-make-date", middleware.AuthMiddleware(handler.PutCantMakeDate))
	// router.Put("/volunteer", middleware.AuthMiddleware(handler.PutVolunteer))

	/** // **/

	/** @_ No Auth middleware applied **/

	router.Post("/register-employee", handler.RegisterEmployee)
	router.Post("/remove-employee", handler.RemoveEmployee)
	router.Get("/location-info", handler.GetLocationStatus)
	router.Post("/appointment", middleware.AuthMiddleware(handler.PostAppointment))
	router.Get("/unconfirmed-appointments", handler.GetUnconfirmedAppointments)
	router.Post("/for-job", handler.PostForJob)

	/** // **/

	/** @_ Debugging utilities **/

	router.Post("/get-employee", handler.DEBUG_GetEmployeeData)
	// router.Post("/get-all-attendance-data", handler.DEBUG_GetAllAttendanceData)
	router.Post("/get-clocked-in-employees", handler.DEBUG_GetClockedInEmployees)
	router.Post("/get-clocked-out-employees", handler.DEBUG_GetClockedOutEmployees)

	/** // **/

	/** @_ Clock in/out **/

	router.Put("/clock-in", handler.ClockIn)
	router.Put("/clock-out", handler.ClockOut)

	/** // **/

	/** @_ Priodically check time **/

	go clockChecker.CheckerLoop(&db)

	/** // **/

	/** @_ Server loop **/

	log.Println("Server listening on port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))

	/** // **/
}
