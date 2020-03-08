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

func TestOauthHandler_HandleClients(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleClients(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleClients2(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.Client
	c.ClientID = 55
	var cl []services.Client
	cl = append(cl, c)
	ser.MockClientList = &cl
	ser.MockClientListCode = 200
	h.Service = &ser

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester", nil)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleClients(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAddClient(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	//var ser services.MockOauth2Service
	//var c services.Client
	//c.ClientID = 55
	//var cl []services.Client
	//cl = append(cl, c)
	//ser.MockClientList = &cl
	//ser.MockCode = 200
	//h.Service = &ser

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester", nil)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAddClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleAddClientAuth(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	//var ser services.MockOauth2Service
	//var c services.Client
	//c.ClientID = 55
	//var cl []services.Client
	//cl = append(cl, c)
	//ser.MockClientList = &cl
	//ser.MockCode = 200
	//h.Service = &ser

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester", nil)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleAddClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleEditClient(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = true
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleEditClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleEditClientAuth(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = true
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleEditClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleNewClient(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = true
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester&webSite=www.test.com&redirectURLs=testurl1,testurl2", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleNewClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleNewClient2(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = false
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester&webSite=www.test.com&redirectURLs=testurl1,testurl2&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleNewClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleNewClientAuth(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = false
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester&webSite=www.test.com&redirectURLs=testurl1,testurl2&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleNewClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateClient(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = true
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester&webSite=www.test.com&redirectURLs=testurl1,testurl2", nil)
	// vars := map[string]string{
	// 	"clientId": "1",
	// }
	// r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateClient2(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = false
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester&webSite=www.test.com&redirectURLs=testurl1,testurl2&enabled=yes", nil)
	// vars := map[string]string{
	// 	"clientId": "1",
	// }
	// r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateClientAuth(t *testing.T) {
	var h OauthHandler
	h.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"

	var ser services.MockOauth2Service
	var c services.ClientResponse
	c.ClientID = 55
	c.Success = true
	c.Code = 200
	//var cl []services.Client
	//cl = append(cl, c)
	ser.MockClientResponse = &c
	//ser.MockCode = 200
	h.Service = &ser

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// h.token = &mTkn

	r, _ := http.NewRequest("POST", "https://test.com?clientName=tester&webSite=www.test.com&redirectURLs=testurl1,testurl2", nil)
	// vars := map[string]string{
	// 	"clientId": "1",
	// }
	// r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateClient(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}
