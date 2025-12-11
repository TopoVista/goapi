package middleware

import (
	"errors"
	"net/http"

	"github.com/TopoVista/goapi/api"
	"github.com/TopoVista/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid usernames or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
	        log.Error("Unauthorized Error")
        	api.RequestErrorHandler(w, UnAuthorizedError)
	        return
        }

		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := database.GetUserLoginDetails(username)
		if loginDetails == nil || token != loginDetails.AuthToken {
			log.Error("Unauthorized Error")
			api.RequestErrorHandler(w, UnAuthorizedError)
			return	
		}

		next.ServeHTTP(w, r)
	})
}	