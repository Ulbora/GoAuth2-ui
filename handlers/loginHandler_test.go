//Package handlers ...
package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	oauth2 "github.com/Ulbora/go-oauth2-client"
)

func TestOauthHandler_authorize(t *testing.T) {
	var h OauthHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"
	h.OauthHost = "test.com"

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	suc := h.authorize(w, r)
	if !suc {
		t.Fail()
	}
}

func TestOauthHandler_HandleLogin(t *testing.T) {
	var h OauthHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"
	h.OauthHost = "test.com"

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	h.HandleLogin(w, r)
}

func TestOauthHandler_handleLogout(t *testing.T) {
	var h OauthHandler
	//h.TokenMap = make(map[string]*oauth2.Token)

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"

	// var mockAcTkn oauth2.MockAuthCodeToken
	// mockAcTkn.MockToken = &mTkn

	// h.AuthToken = &mockAcTkn

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["accessTokenKey"] = "123"

	w := httptest.NewRecorder()
	s.Save(r, w)
	h.HandleLogout(w, r)
}

func TestOauthHandler_handleToken(t *testing.T) {
	var h OauthHandler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"

	var mockAcTkn oauth2.MockAuthCodeToken
	mockAcTkn.MockToken = &mTkn

	h.Auth = &mockAcTkn

	var cc ClientCreds
	cc.AuthCodeState = "123"
	cc.AuthCodeClient = "2"
	cc.AuthCodeSecret = "12345"
	h.ClientCreds = &cc
	h.OauthHost = "http://test12.com"
	r, _ := http.NewRequest("POST", "https://test.com?code=555&state=123", nil)
	w := httptest.NewRecorder()
	h.HandleToken(w, r)

}
