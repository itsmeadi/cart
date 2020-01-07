package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/templatego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

const (
	SessionKey  = "ntuc"
	errorFormat = "[sessions] ERROR! %s\n"
)

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

// User is a retrieved and authentiacted user.
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Gender        string `json:"gender"`
}

var cred Credentials
var conf *oauth2.Config
var state string
var store = sessions.NewCookieStore([]byte("secret"))

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func init() {
	//file, err := ioutil.ReadFile("./creds.json")
	//if err != nil {
	//	log.Printf("File error: %v\n", err)
	//	os.Exit(1)
	//}
	//json.Unmarshal(file, &cred)
	//

	cred.Cid = `505592410251-3c6j54kt3u9ksop1eoe1vsffdhd3egur.apps.googleusercontent.com`
	cred.Csecret = `gHREtDeHSPfeehqON08Cig1W`
	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:9090/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func AbortWithError(code int, err error, w *http.ResponseWriter) {
	res := *w
	res.WriteHeader(code)

	errStr := err.Error()
	_, _ = res.Write([]byte(errStr))

}

func (api *API) Auth(hand Handler) Handler {

	return Handler(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			session, err := store.Get(r, SessionKey)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//Handle the exchange code to initiate a transport.
			userId := session.Values["user-id"]
			if userId == nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			userEmail := session.Values["user-email"]
			if userEmail == nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			//w.Header().Set("Content-Type", "text/html")
			//http.Redirect(w, r, r.Host+"/login.html", http.StatusSeeOther)

			ctx := r.Context()

			ctx = context.WithValue(ctx, "user_id", userId)
			ctx = context.WithValue(ctx, "user_email", userEmail)
			r = r.WithContext(ctx)
			hand(w, r)
		})

}

func (api *API) logOut(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SessionKey)

	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		log.Println("[API][logOut][session.Save()]Error=", err)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return

}

func (api *API) authHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	session, err := store.Get(r, SessionKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Handle the exchange code to initiate a transport.
	retrievedState := session.Values["state"]
	if retrievedState != r.FormValue("state") {
		w.WriteHeader(http.StatusUnauthorized)
		//log.Printf("Invalid session state: %+v %+v", retrievedState, r.FormValue("state"))
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tok, err := conf.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		AbortWithError(http.StatusBadRequest, err, &w)
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		AbortWithError(http.StatusBadRequest, err, &w)
		return
	}
	defer email.Body.Close()
	data, _ := ioutil.ReadAll(email.Body)
	var userG models.UserGoogle
	err = json.Unmarshal(data, &userG)
	if err != nil {
		log.Println(string(data))
		AbortWithError(http.StatusBadRequest, err, &w)
		return
	}

	user, err := api.Interactor.User.GetUserBySub(ctx, userG.Sub)
	uid := user.ID
	if err != nil {
		AbortWithError(http.StatusBadRequest, err, &w)
		return
	}
	if uid == 0 {
		uid, err = api.Interactor.User.AddUser(ctx, models.User{Sub: userG.Sub})
		if err != nil {
			AbortWithError(http.StatusBadRequest, err, &w)
			return
		}
	}
	session.Values["user-id"] = uid
	session.Values["user-email"] = userG.Email
	err = session.Save(r, w)
	http.Redirect(w, r, "/products?category_id=3", http.StatusSeeOther)

}

func (api *API) googleLoginHandler(w http.ResponseWriter, r *http.Request) {

	state = randToken()
	session, err := store.Get(r, SessionKey)
	if err != nil {
		log.Println(err)
		AbortWithError(http.StatusBadRequest, err, &w)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		AbortWithError(http.StatusInternalServerError, err, &w)
		return
	}
	qtemplate := struct {
		GoogleLoginUrl string
	}{
		GoogleLoginUrl: getLoginURL(state),
	}
	if err := templatego.TemplateMap["login"].Execute(w, qtemplate); err != nil {
		log.Printf("[ERROR] [Question] Render page error: %s\n", err)

	}
}

//
//func (api *API) loginHandler(w http.ResponseWriter, r *http.Request) {
//
//	state = randToken()
//	session, err := store.Get(r, SessionKey)
//	if err != nil {
//		log.Println(err)
//		AbortWithError(http.StatusBadRequest, err, &w)
//		return
//	}
//	session.Values["state"] = state
//	err = session.Save(r, w)
//	if err != nil {
//		AbortWithError(http.StatusInternalServerError, err, &w)
//		return
//	}
//	w.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + getLoginURL(state) + "'><button>Login with Google!</button> </a> </body></html>"))
//}

//
//func main() {
//	router := gin.Default()
//	router.Use(sessions.Sessions("goquestsession", store))
//	router.Static("/css", "./static/css")
//	router.Static("/img", "./static/img")
//	router.LoadHTMLGlob("./templates/*")
//
//	router.GET("/", indexHandler)
//	router.GET("/login", loginHandler)
//	router.GET("/auth", authHandler)
//
//	router.Run("127.0.0.1:9090")
//}
