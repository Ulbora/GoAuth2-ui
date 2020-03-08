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

func TestOauthHandler_HandleRoles(t *testing.T) {
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
	ser.MockClientCode = 200

	var cr services.ClientRole
	cr.ClientID = 55
	cr.ID = 1
	cr.Role = "tester"
	var crs []services.ClientRole
	crs = append(crs, cr)
	ser.MockClientRoleList = &crs
	ser.MockRoleURIListCode = 200

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
	h.HandleRoles(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRolesAuth(t *testing.T) {
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
	ser.MockClientCode = 200

	var cr services.ClientRole
	cr.ClientID = 55
	cr.ID = 1
	cr.Role = "tester"
	var crs []services.ClientRole
	crs = append(crs, cr)
	ser.MockClientRoleList = &crs
	ser.MockRoleURIListCode = 200

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
	h.HandleRoles(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRoleAdd(t *testing.T) {
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

	var aures services.ClientRole
	aures.ID = 5
	aures.ClientID = 1
	aures.Role = "tester"
	var auress []services.ClientRole
	auress = append(auress, aures)
	ser.MockClientRoleList = &auress

	var rures services.ClientRoleResponse
	rures.Code = 200
	rures.Success = true
	ser.MockClientRoleResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?clientRole=tester2&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRoleAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRoleAddExists(t *testing.T) {
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

	var aures services.ClientRole
	aures.ID = 5
	aures.ClientID = 1
	aures.Role = "tester"
	var auress []services.ClientRole
	auress = append(auress, aures)
	ser.MockClientRoleList = &auress

	var rures services.ClientRoleResponse
	rures.Code = 200
	rures.Success = true
	ser.MockClientRoleResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?clientRole=tester&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRoleAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRoleAddAuth(t *testing.T) {
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

	var aures services.ClientRole
	aures.ID = 5
	aures.ClientID = 1
	aures.Role = "tester"
	var auress []services.ClientRole
	auress = append(auress, aures)
	ser.MockClientRoleList = &auress

	var rures services.ClientRoleResponse
	rures.Code = 200
	rures.Success = true
	ser.MockClientRoleResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?clientRole=tester2&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleRoleAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRoleDelete(t *testing.T) {
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

	var rures services.ClientRoleResponse
	rures.Code = 200
	rures.Success = true
	ser.MockClientRoleResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com", nil)
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
	h.HandleRoleDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleRoleDeleteAuth(t *testing.T) {
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

	var rures services.ClientRoleResponse
	rures.Code = 200
	rures.Success = true
	ser.MockClientRoleResponse = &rures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com", nil)
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
	h.HandleRoleDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}
