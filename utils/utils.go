/*
Module consisting of utility functions.
*/
package utils

import (
	"context"
	"fmt"
	"net/mail"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

/*
Check for error, handle error and terminate program if present.
*/
func Handle_Error(err error) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		os.Exit(1)
	}

}

/*
Connects to MongoDB on local machine or Atlas.
*/
func Connect_Db() (*mongo.Client, context.Context, context.CancelFunc) {

	client, err := mongo.NewClient(
		options.Client().ApplyURI(
			"mongodb://172.17.0.2:27017",
		),
	)
	Handle_Error(err)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	err = client.Connect(ctx)
	Handle_Error(err)

	return client, ctx, cancel

}

/*
Hash the password recieved from the client.
NOTE:
* Hashing password on the client is prone to manipulation by the user
* Password must be sent from client to server via HTTPS (POST method) using TLS, otherwise an eavesdropper can easily get access to it
*/
func Hash_Password(clearPassword string) []byte {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(clearPassword),
		bcrypt.DefaultCost,
	)
	Handle_Error(err)

	return hashedPassword

}

/*
Compare hashed password to clear text input by client for a match.
Return value of `true` means verification success.
NOTE: `clearPassword` will NOT be stored on any non-volatile media.
*/
func Verify_Hash(hashedPassword []byte, clearPassword string) bool {

	err := bcrypt.CompareHashAndPassword(
		hashedPassword,
		[]byte(clearPassword),
	)

	return err == nil

}

/*
Check whether input string is a valid email address.
*/
func Email_Is_Valid(email string) bool {

	_, err := mail.ParseAddress(email)

	return err == nil

}
