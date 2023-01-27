package config

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/iktefish/stealth-server/constants"
	"google.golang.org/api/option"
)

/*
Initiates a new Firebase SDK, from 'service-key.json' file, and a Firestore Client from
the SDK.
*/
func NewSdkAndClient() (*firebase.App, *firestore.Client, error) {
	var app, err_1 = InitFirebaseSdk()
	if err_1 != nil {
		return nil, nil, err_1
	}

	var ctx = context.Background()
	var client, err_2 = app.Firestore(ctx)
	if err_2 != nil {
		return nil, nil, err_2
	}

	return app, client, nil
}

/*
Initiates a new Firebase SDK from 'service-key.json' file. You can generate this file from
"Firebase Console" > "Project Overview" > "Service accounts" > "Generate new private key"
(select Go).
*/
func InitFirebaseSdk() (*firebase.App, error) {
	var ctx = context.Background()
	var opt = option.WithCredentialsFile(constants.ServiceKeyPath)
	var app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
