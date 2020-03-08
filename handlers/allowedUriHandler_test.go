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

func TestOauthHandler_HandleAllowedUris(t *testing.T) {
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

	var cr services.ClientRole
	cr.ClientID = 55
	cr.ID = 1
	cr.Role = "tester"
	var crs []services.ClientRole
	crs = append(crs, cr)
	ser.MockClientRoleList = &crs

	var cru services.RoleURI
	cru.ClientAllowedURIID = 3
	cru.ClientRoleID = 1
	var crus []services.RoleURI
	crus = append(crus, cru)
	ser.MockRoleURIList = &crus

	var au services.AllowedURI
	au.ClientID = 55
	au.ID = 4
	au.URI = "someurl"
	var aus []services.AllowedURI
	aus = append(aus, au)
	ser.MockAllowedURIList = &aus

	ser.MockAllowedURIListCode = 200
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
	h.HandleAllowedUris(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAllowedUrisAuth(t *testing.T) {
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
	var cl []services.Client
	cl = append(cl, c)
	ser.MockClientList = &cl

	var cr services.ClientRole
	cr.ClientID = 55
	cr.ID = 1
	cr.Role = "tester"
	var crs []services.ClientRole
	crs = append(crs, cr)
	ser.MockClientRoleList = &crs

	var cru services.RoleURI
	cru.ClientAllowedURIID = 3
	cru.ClientRoleID = 1
	var crus []services.RoleURI
	crus = append(crus, cru)
	ser.MockRoleURIList = &crus

	var au services.AllowedURI
	au.ClientID = 55
	au.ID = 4
	au.URI = "someurl"
	var aus []services.AllowedURI
	aus = append(aus, au)
	ser.MockAllowedURIList = &aus

	ser.MockAllowedURIListCode = 200
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
	h.HandleAllowedUris(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAllowedUrisAdd(t *testing.T) {
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

	var aures services.AllowedURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockAllowedURIResponse = &aures

	var rures services.RoleURIResponse
	rures.Code = 200
	rures.Success = true
	ser.MockRoleURIResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAllowedUrisAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAllowedUrisAddAuth(t *testing.T) {
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

	var aures services.AllowedURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockAllowedURIResponse = &aures

	var rures services.RoleURIResponse
	rures.Code = 200
	rures.Success = true
	ser.MockRoleURIResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAllowedUrisAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAllowedUrisUpdate(t *testing.T) {
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

	var aures services.AllowedURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockAllowedURIResponse = &aures

	var rures services.RoleURIResponse
	rures.Code = 200
	rures.Success = true
	ser.MockRoleURIResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAllowedUrisUpdate(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAllowedUrisUpdateAuth(t *testing.T) {
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

	var aures services.AllowedURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockAllowedURIResponse = &aures

	var rures services.RoleURIResponse
	rures.Code = 200
	rures.Success = true
	ser.MockRoleURIResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAllowedUrisUpdate(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAllowedUrisDelete(t *testing.T) {
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

	var aures services.AllowedURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockAllowedURIResponse = &aures

	var rures services.RoleURIResponse
	rures.Code = 200
	rures.Success = true
	ser.MockRoleURIResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"id":       "5",
		"clientId": "1",
		"roleId":   "6",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAllowedUrisDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAllowedUrisDeleteAuth(t *testing.T) {
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

	var aures services.AllowedURIResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockAllowedURIResponse = &aures

	var rures services.RoleURIResponse
	rures.Code = 200
	rures.Success = true
	ser.MockRoleURIResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?roleId=3&clientId=5", nil)
	vars := map[string]string{
		"id":       "5",
		"clientId": "1",
		"roleId":   "6",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAllowedUrisDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}
