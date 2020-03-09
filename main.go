package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	hand "github.com/Ulbora/GoAuth2-ui/handlers"
	services "github.com/Ulbora/GoAuth2-ui/services"
	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

//GO111MODULE=on go mod init github.com/Ulbora/GoAuth2-ui
func main() {
	var authCodeClientID string
	if os.Getenv("AUTH_CODE_CLIENT_ID") != "" {
		authCodeClientID = os.Getenv("AUTH_CODE_CLIENT_ID")
	} else {
		authCodeClientID = "10"
	}

	var goauth2Host string
	if os.Getenv("GOAUTH2_HOST") != "" {
		goauth2Host = os.Getenv("GOAUTH2_HOST")
	} else {
		goauth2Host = "http://localhost:3000"
	}

	var schemeDefault string
	if os.Getenv("SCHEME_DEFAULT") != "" {
		schemeDefault = os.Getenv("SCHEME_DEFAULT")
	} else {
		schemeDefault = "http://"
	}

	var oh hand.OauthHandler
	var logger lg.Logger
	logger.LogLevel = lg.InfoLevel
	oh.Log = &logger
	var cc hand.ClientCreds
	cc.AuthCodeState = "58dkhhhd"
	oh.ClientCreds = &cc
	oh.ClientCreds.AuthCodeClient = authCodeClientID
	oh.OauthHost = goauth2Host

	var ser services.Oauth2Service
	ser.Log = &logger
	var p px.GoProxy
	ser.Proxy = p.GetNewProxy()
	ser.Host = goauth2Host
	ser.ClientID = authCodeClientID
	oh.Service = ser.GetNew()
	oh.SchemeDefault = schemeDefault
	var act oauth2.AuthCodeToken
	oh.Auth = &act

	oh.Templates = template.Must(template.ParseFiles("./static/index.html", "./static/header.html", "./static/headerChart.html",
		"./static/footer.html", "./static/navbar.html", "./static/clients.html", "./static/addClient.html",
		"./static/editClient.html", "./static/oauth2.html", "./static/redirectUrls.html", "./static/grantTypes.html",
		"./static/roles.html", "./static/allowedUris.html", "./static/secSideMenu.html",
		"./static/users.html", "./static/editUser.html"))

	h := oh.GetNew()

	router := mux.NewRouter()
	port := "8091"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	//securety routes
	router.HandleFunc("/", h.HandleIndex).Methods("GET")
	router.HandleFunc("/clients", h.HandleClients).Methods("GET")
	router.HandleFunc("/addClient", h.HandleAddClient).Methods("GET")
	router.HandleFunc("/editClient/{clientId}", h.HandleEditClient).Methods("GET")
	router.HandleFunc("/newClient", h.HandleNewClient).Methods("POST")
	router.HandleFunc("/updateClient", h.HandleUpdateClient).Methods("POST")

	router.HandleFunc("/users/{clientId}", h.HandleUsers).Methods("GET")
	router.HandleFunc("/newUser", h.HandleNewUser).Methods("POST")
	router.HandleFunc("/editUser/{username}/{clientId}", h.HandleEditUser).Methods("GET")
	router.HandleFunc("/updateUserInfo", h.HandleUpdateUserInfo).Methods("POST")
	router.HandleFunc("/updateUserEnable", h.HandleUpdateUserEnable).Methods("POST")
	router.HandleFunc("/updateUserPw", h.HandleUpdateUserPw).Methods("POST")

	router.HandleFunc("/oauth2/{clientId}", h.HandleOauth2).Methods("GET")

	router.HandleFunc("/clientRedirectUrls/{clientId}", h.HandleRedirectURLs).Methods("GET")
	router.HandleFunc("/addRedirectUrl", h.HandleRedirectURLAdd).Methods("POST")
	router.HandleFunc("/deleteRedirectUri/{id}/{clientId}", h.HandleRedirectURLDelete).Methods("GET")

	router.HandleFunc("/clientGrantTypes/{clientId}", h.HandleGrantType).Methods("GET")
	router.HandleFunc("/addGrantType", h.HandleGrantTypeAdd).Methods("POST")
	router.HandleFunc("/deleteGrantType/{id}/{clientId}", h.HandleGrantTypeDelete).Methods("GET")

	router.HandleFunc("/clientRoles/{clientId}", h.HandleRoles).Methods("GET")
	router.HandleFunc("/addClientRole", h.HandleRoleAdd).Methods("POST")
	router.HandleFunc("/deleteClientRoles/{id}/{clientId}", h.HandleRoleDelete).Methods("GET")

	router.HandleFunc("/clientAllowedUris/{clientId}", h.HandleAllowedUris).Methods("GET")
	router.HandleFunc("/addAllowedUri", h.HandleAllowedUrisAdd).Methods("POST")
	router.HandleFunc("/editAllowedUri", h.HandleAllowedUrisUpdate).Methods("POST")
	router.HandleFunc("/deleteAllowedUri/{id}/{roleId}/{clientId}", h.HandleAllowedUrisDelete).Methods("GET")

	router.HandleFunc("/tokenHandler", h.HandleToken)
	router.HandleFunc("/login", h.HandleLogin)
	router.HandleFunc("/logout", h.HandleLogout)

	// admin resources
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	//http.Handle("/js", fs)

	fmt.Println("GoAuth2-ui is running on port:" + port)
	log.Println("Listening on port: " + port)
	http.ListenAndServe(":"+port, router)

}
