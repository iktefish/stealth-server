package middleware

import "net/http"

func AuthMiddleware(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	/*
	   Authenticate ...
	*/

	return h
}
