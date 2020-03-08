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

func TestOauthHandler_HandleRedirectURLs(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	var ser services.MockOauth2Service
	var c services.Client
	c.ClientID = 55
	ser.MockClient = &c
	var cl []services.Client
	cl = append(cl, c)
	ser.MockClientList = &cl
	ser.MockClientListCode = 200

	var cr services.RedirectURI
	cr.ClientID = 55
	cr.ID = 1
	cr.URI = "/tester"

	var cr2 services.RedirectURI
	cr2.ClientID = 55
	cr2.ID = 11
	cr2.URI = "/tester2"

	var crs []services.RedirectURI
	crs = append(crs, cr)
	crs = append(crs, cr2)
	ser.MockRedirectURIList = &crs
	ser.MockRedirectURIListCode = 200

	//ser.MockAllowedURIListCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRedirectURLs(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRedirectURLsAuth(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// h.token = &mTkn

	var ser services.MockOauth2Service
	var c services.Client
	c.ClientID = 55
	ser.MockClient = &c
	var cl []services.Client
	cl = append(cl, c)
	ser.MockClientList = &cl
	ser.MockClientListCode = 200

	var cr services.RedirectURI
	cr.ClientID = 55
	cr.ID = 1
	cr.URI = "/tester"

	var cr2 services.RedirectURI
	cr2.ClientID = 55
	cr2.ID = 11
	cr2.URI = "/tester2"

	var crs []services.RedirectURI
	crs = append(crs, cr)
	crs = append(crs, cr2)
	ser.MockRedirectURIList = &crs
	ser.MockRedirectURIListCode = 200

	//ser.MockAllowedURIListCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRedirectURLs(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRedirectURLAdd(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	var ser services.MockOauth2Service

	var aures services.RedirectURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockRedirectURIResponse = &aures

	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?redirectURL=tester&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRedirectURLAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRedirectURLAddAuth(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// h.token = &mTkn

	var ser services.MockOauth2Service

	var aures services.RedirectURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockRedirectURIResponse = &aures

	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?redirectURL=tester&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRedirectURLAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRedirectURLDelete(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	var ser services.MockOauth2Service

	var aures services.RedirectURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockRedirectURIResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"id":       "5",
		"clientId": "1",
		//"roleId":   "6",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRedirectURLDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRedirectURLDeleteAuth(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// h.token = &mTkn

	var ser services.MockOauth2Service

	var aures services.RedirectURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockRedirectURIResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"id":       "5",
		"clientId": "1",
		//"roleId":   "6",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRedirectURLDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}
