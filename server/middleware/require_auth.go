package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ritwik310/my-website/server/db"
	"github.com/ritwik310/my-website/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Writes Admin Un-Authenticated on Response
func writeUnauth(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("{\"message\": \"admin unauthorized\"}"))
}

// AuthRequired - Middleware that checks authentication
func AuthRequired(handler http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var err error
		var admin models.Admin

		var eCookie *http.Cookie // "admin-email" Cookie
		var hCookie *http.Cookie // "admin-id" Cookie

		eCookie, err = r.Cookie("admin-email")
		hCookie, err = r.Cookie("admin-id")

		//  Cookie error handleing
		if err != nil {
			fmt.Println("Error: authentication error:", err.Error())
			writeUnauth(w)
			return
		}

		// Mongo Collection
		collection := db.Client.Database("dev_db").Collection("admins")

		// Query User from Database (by Email)
		err = nil
		err = collection.FindOne(context.TODO(), bson.D{bson.E{Key: "email", Value: eCookie.Value}}).Decode(&admin)
		if err != nil {
			fmt.Println("Error: authentication error:", err.Error())
			writeUnauth(w)
			return
		}

		// Compare Hashed ID
		err = nil
		err = bcrypt.CompareHashAndPassword([]byte(hCookie.Value), []byte(admin.GoogleID))
		if err != nil {
			fmt.Println("Error: authentication error:", err.Error())
			writeUnauth(w)
			return
		}

		handler.ServeHTTP(w, r)
		// handler.ServeHTTP(w, r, p)
	}
}
