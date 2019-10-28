package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/context"
	"github.com/ritcrap/my-website/server/db"
	"github.com/ritcrap/my-website/server/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthStateStr string         // pseudo-random string
	oauthConfig   *oauth2.Config // contains configuration details for google oauth login
)

func init() {
	oauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_AUTH_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	oauthStateStr = "pseudo-random"
}

// GoogleLogin redirects login requests to google's oauth api
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateStr)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirecting to Google
}

// GoogleCallback handles googles user info response while oauth login
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// reading the user's data from google apis
	info, err := userInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	var data db.Admin
	err = json.Unmarshal(info, &data)
	if err != nil {
		writeErr(w, 500, err)
		return
	}

	// checking if that email is allowed to login or not, basically if admin's email or not
	switch data.Email {
	case os.Getenv("AUTHORIZED_EMAIL_I"):
		logrus.Infof("admin login - %v\n", data.Email)
	case os.Getenv("AUTHORIZED_EMAIL_II"):
		logrus.Infof("admin login - %v\n", data.Email)
	default:
		writeErr(w, 500, fmt.Errorf("email %v is not authorized for login", data.Email))
		return
	}

	// querying the admin data from database
	var admin db.Admin
	err = admin.Read(bson.M{"email": data.Email, "google_id": data.GoogleID})

	switch err {
	case mgo.ErrNotFound:
		// no record (row) found with the admin's email
		logrus.Infof("creating new admin - %v\n", admin.Email)

		// Inserting Document
		admin = data
		err = admin.Create()
		if err != nil {
			logrus.Warnf("couldn't create a new admin, %v\n", err)
			writeErr(w, 500, err)
			return
		}
	case nil:
		// everything's fine, user exists and allowed to login
	default:
		// some internal error
		writeErr(w, 500, err)
		return
	}

	// setting up user's cookie for authentication
	session, err := middleware.CookieStore.Get(r, "session")
	session.Values["admin_email"] = admin.Email
	session.Save(r, w)

	// writing the admin json to the client
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(admin)

	// redirecting to current_user route
	http.Redirect(w, r, os.Getenv("ADMIN_ORIGIN"), http.StatusSeeOther)
}

// userInfo sends a request to google apis and gets the user's data for us
func userInfo(state, code string) ([]byte, error) {
	if state != oauthStateStr {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// CurrentUser writes currenntly logged in user's info to teh client
func CurrentUser(w http.ResponseWriter, r *http.Request) {
	// reading session info passed via middleware.CheckAuth middleware
	// that this function (handler) is going to be passed on in the router
	email := context.Get(r, "admin_e")

	if email == nil {
		writeErr(w, http.StatusBadRequest, fmt.Errorf("unable to read auth cookie"))
		return
	}

	var admin db.Admin
	err := admin.Read(bson.M{"email": email.(string)})
	if err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, admin)

	// enabling C.O.R.S.
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(admin)
}
