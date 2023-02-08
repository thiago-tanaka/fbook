package middlewares

import (
	"fmt"
	"net/http"
	"webapp/src/cookies"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.Read(r)

		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		next(w, r)
	}
}
