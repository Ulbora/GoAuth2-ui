//Package handlers ...
package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	services "github.com/Ulbora/GoAuth2-ui/services"
	lg "github.com/Ulbora/Level_Logger"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

func TestOauthHandler_HandleOauth2(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"
	h.OauthHost = "test.com"

	var ser services.MockOauth2Service
	var c services.Client
	c.ClientID = 55
	var cl []services.Client
	cl = append(cl, c)
	ser.MockClientList = &cl
	ser.MockClientListCode = 200
	h.Service = &ser

	// var serv services.Oauth2Service
	// serv.ClientID = h.ClientCreds.AuthCodeClient
	// serv.Host = h.OauthHost

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleOauth2(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleOauth2_2(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	// var mockAcTkn oauth2.MockAuthCodeToken
	// mockAcTkn.MockToken = &mTkn

	// h.AuthToken = &mockAcTkn
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"
	h.OauthHost = "test.com"

	var ser services.MockOauth2Service
	var c services.Client
	c.ClientID = 55
	var cl []services.Client
	cl = append(cl, c)
	ser.MockClientList = &cl
	ser.MockClientListCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"clientId": "5",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleOauth2(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}
