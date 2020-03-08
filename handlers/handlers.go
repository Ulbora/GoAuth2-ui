//Package handlers ...
package handlers

import (
	"html/template"
	"net/http"

	s "github.com/Ulbora/GoAuth2-ui/services"
	lg "github.com/Ulbora/Level_Logger"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	gs "github.com/Ulbora/go-sessions"
	"github.com/gorilla/sessions"
)

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//Handler Handler
type Handler interface {
	HandleClients(w http.ResponseWriter, r *http.Request)
	HandleAddClient(w http.ResponseWriter, r *http.Request)
	HandleEditClient(w http.ResponseWriter, r *http.Request)
	HandleNewClient(w http.ResponseWriter, r *http.Request)
	HandleUpdateClient(w http.ResponseWriter, r *http.Request)

	HandleAllowedUris(w http.ResponseWriter, r *http.Request)
	HandleAllowedUrisAdd(w http.ResponseWriter, r *http.Request)
	HandleAllowedUrisUpdate(w http.ResponseWriter, r *http.Request)
	HandleAllowedUrisDelete(w http.ResponseWriter, r *http.Request)

	// HandleGrantType(w http.ResponseWriter, r *http.Request)
	// HandleGrantTypeAdd(w http.ResponseWriter, r *http.Request)
	// HandleGrantTypeDelete(w http.ResponseWriter, r *http.Request)

	HandleOauth2(w http.ResponseWriter, r *http.Request)

	// HandleRedirectURLs(w http.ResponseWriter, r *http.Request)
	// HandleRedirectURLAdd(w http.ResponseWriter, r *http.Request)
	// HandleRedirectURLDelete(w http.ResponseWriter, r *http.Request)

	// HandleRoles(w http.ResponseWriter, r *http.Request)
	// HandleRoleAdd(w http.ResponseWriter, r *http.Request)
	// HandleRoleDelete(w http.ResponseWriter, r *http.Request)

	// HandleUsers(w http.ResponseWriter, r *http.Request)
	// HandleNewUser(w http.ResponseWriter, r *http.Request)
	// HandleEditUser(w http.ResponseWriter, r *http.Request)
	// HandleUpdateUserInfo(w http.ResponseWriter, r *http.Request)
	// HandleUpdateUserEnable(w http.ResponseWriter, r *http.Request)
	// HandleUpdateUserPw(w http.ResponseWriter, r *http.Request)

	HandleIndex(w http.ResponseWriter, r *http.Request)
}

const (
	authCodeRedirectURI = "/tokenHandler"
	authCodeState       = "ghh66555h"
)

//ClientCreds ClientCreds
type ClientCreds struct {
	AuthCodeClient string
	AuthCodeSecret string
	AuthCodeState  string
	// SchemeDefault  string // = "http://"
}

//OauthHandler OauthHandler
type OauthHandler struct {
	Service   s.Service
	Session   gs.GoSession
	Templates *template.Template
	Store     *sessions.CookieStore
	//TokenMap      map[string]*oauth2.Token
	ClientCreds   *ClientCreds
	OauthHost     string
	UserHost      string
	SchemeDefault string // = "http://"
	AuthToken     oauth2.AuthToken
	token         *oauth2.Token
	//Client          oa.Client
	//AssetControl    rc.AssetControl
	//TokenCompressed bool
	//JwtCompress     cp.JwtCompress
	Log *lg.Logger
}

//GetNew GetNew
func (h *OauthHandler) GetNew() Handler {
	var hh Handler
	hh = h
	return hh
}

// user handlers-----------------------------------------------------

//HandleIndex HandleIndex
func (h *OauthHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	h.Templates.ExecuteTemplate(w, "index.html", nil)
}

func (h *OauthHandler) getSession(r *http.Request) (*sessions.Session, bool) {
	//fmt.Println("getSession--------------------------------------------------")
	var suc bool
	var srtn *sessions.Session
	if h.Store == nil {
		h.Session.Name = "goauth2"
		h.Session.MaxAge = 3600
		h.Store = h.Session.InitSessionStore()
		//errors without this
		//-------gob.Register(&AuthorizeRequestInfo{})
	}
	if r != nil {
		// fmt.Println("secure in getSession", h.Session.Secure)
		// fmt.Println("name in getSession", h.Session.Name)
		// fmt.Println("MaxAge in getSession", h.Session.MaxAge)
		// fmt.Println("SessionKey in getSession", h.Session.SessionKey)

		//h.Session.HTTPOnly = true

		//h.Session.InitSessionStore()
		s, err := h.Store.Get(r, h.Session.Name)
		//s, err := store.Get(r, "temp-name")
		//s, err := store.Get(r, "goauth2")

		loggedInAuth := s.Values["loggedIn"]
		userAuth := s.Values["user"]
		h.Log.Debug("loggedIn: ", loggedInAuth)
		h.Log.Debug("user: ", userAuth)

		larii := s.Values["authReqInfo"]
		h.Log.Debug("arii-----login", larii)

		h.Log.Debug("session error in getSession: ", err)
		if err == nil {
			suc = true
			srtn = s
		}
	}
	//fmt.Println("exit getSession--------------------------------------------------")
	return srtn, suc
}
