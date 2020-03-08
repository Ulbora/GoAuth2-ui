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

func TestOauthHandler_handleGrantType(t *testing.T) {
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

	var gt services.GrantType
	gt.ClientID = 55
	gt.ID = 1
	gt.GrantType = "authcode"
	var gts []services.GrantType
	gts = append(gts, gt)
	ser.MockGrantTypeList = &gts
	ser.MockGrantTypeListCode = 200

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
	h.handleGrantType(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_handleGrantTypeAuth(t *testing.T) {
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

	var gt services.GrantType
	gt.ClientID = 55
	gt.ID = 1
	gt.GrantType = "authcode"
	var gts []services.GrantType
	gts = append(gts, gt)
	ser.MockGrantTypeList = &gts
	ser.MockGrantTypeListCode = 200

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
	h.handleGrantType(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_handleGrantTypeAdd(t *testing.T) {
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

	var aures services.GrantTypeResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockGrantTypeResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?grantType=3&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.handleGrantTypeAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_handleGrantTypeAddAuth(t *testing.T) {
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

	var aures services.GrantTypeResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockGrantTypeResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?grantType=3&clientId=5", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.handleGrantTypeAdd(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_handleGrantTypeDelete(t *testing.T) {
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

	var aures services.GrantTypeResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockGrantTypeResponse = &aures

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
	h.handleGrantTypeDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_handleGrantTypeDeleteAuth(t *testing.T) {
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

	var aures services.GrantTypeResponse
	aures.Code = 200
	aures.ID = 5
	aures.Success = true
	ser.MockGrantTypeResponse = &aures

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
	h.handleGrantTypeDelete(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}
