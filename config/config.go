package config

import (
	"context"
	"fmt"
	"os"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"

	// "github.com/iktefish/stealth-server/constants"
	"google.golang.org/api/option"
)

/* Initiates a new Firebase SDK (using 'service-key.json' file) and from it construct Firebase Auth client and a Firestore Client. */
func NewSdkAndClients() (*firebase.App, *auth.Client, *firestore.Client, error) {
	var app, err_1 = InitFirebaseSdk()
	if err_1 != nil {
		return nil, nil, nil, err_1
	}

	var ctx = context.Background()
	var authClient, err_2 = app.Auth(ctx)
	if err_2 != nil {
		return nil, nil, nil, err_2
	}

	var storeClient, err_3 = app.Firestore(ctx)
	if err_3 != nil {
		return nil, nil, nil, err_3
	}

	return app, authClient, storeClient, nil
}

/* Initiates a new Firebase SDK from 'service-key.json' file. You can generate this file from "Firebase Console" > "Project Overview" > "Service accounts" > "Generate new private key" (select Go). */
func InitFirebaseSdk() (*firebase.App, error) {
	var ctx = context.Background()

	var key = os.Getenv("KEY")

	fmt.Println("---")
	fmt.Println(key)
	fmt.Println("---")

	var opt = option.WithCredentialsJSON([]byte(key))
	// var opt = option.WithCredentialsFile(constants.ServiceKeyPath)

	var app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
