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

func TestOauthHandler_HandleUsers(t *testing.T) {
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
	ser.MockClientCode = 200

	var cr services.User
	cr.ClientID = 55
	cr.Username = "tester"
	//cr.Role = "tester"
	var crs []services.User
	crs = append(crs, cr)
	ser.MockUserList = &crs
	ser.MockUserListCode = 200

	var cru services.Role
	cru.ID = 2
	var crus []services.Role
	crus = append(crus, cru)
	ser.MockRoleList = &crus
	ser.MockRoleListCode = 200

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
	h.HandleUsers(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUsersAuth(t *testing.T) {
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
	ser.MockClientCode = 200

	var cr services.User
	cr.ClientID = 55
	cr.Username = "tester"
	//cr.Role = "tester"
	var crs []services.User
	crs = append(crs, cr)
	ser.MockUserList = &crs
	ser.MockUserListCode = 200

	var cru services.Role
	cru.ID = 2
	var crus []services.Role
	crus = append(crus, cru)
	ser.MockRoleList = &crus
	ser.MockRoleListCode = 200

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
	h.HandleUsers(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleNewUser(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleNewUser(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleNewUserAuth(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleNewUser(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleEditUser(t *testing.T) {
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
	ser.MockClientCode = 200

	var cr services.User
	cr.ClientID = 55
	cr.Username = "tester"
	ser.MockUser = &cr
	ser.MockUserCode = 200

	var cru services.Role
	cru.ID = 2
	var crus []services.Role
	crus = append(crus, cru)
	ser.MockRoleList = &crus
	ser.MockRoleListCode = 200

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
	h.HandleEditUser(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 200 {
		t.Fail()
	}
}

func TestOauthHandler_HandleEditUserAuth(t *testing.T) {
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
	ser.MockClientCode = 200

	var cr services.User
	cr.ClientID = 55
	cr.Username = "tester"
	ser.MockUser = &cr
	ser.MockUserCode = 200

	var cru services.Role
	cru.ID = 2
	var crus []services.Role
	crus = append(crus, cru)
	ser.MockRoleList = &crus
	ser.MockRoleListCode = 200

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
	h.HandleEditUser(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateUserInfo(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateUserInfo(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateUserInfoAuth(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateUserInfo(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateUserEnabled(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateUserEnable(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateUserEnabledAuth(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateUserEnable(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateUserPw(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateUserPw(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestOauthHandler_HandleUpdateUserPwAuth(t *testing.T) {
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

	var aures services.UserResponse
	aures.Code = 200
	aures.Success = true
	ser.MockUserResponse = &aures

	//ser.MockCode = 200
	h.Service = &ser

	r, _ := http.NewRequest("POST", "https://test.com?username=tester&clientId=5&userRoleId=5&firstName=test&lastName=er&emailAddress=test@t.com&password=pass&enabled=yes", nil)
	vars := map[string]string{
		"clientId": "1",
	}
	r = mux.SetURLVars(r, vars)
	s, suc := h.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["userLoggenIn"] = true

	w := httptest.NewRecorder()
	h.HandleUpdateUserPw(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}
