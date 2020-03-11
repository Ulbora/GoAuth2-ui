//Package handlers ...
package handlers

import (
	"net/http"

	oauth2 "github.com/Ulbora/go-oauth2-client"
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

//HandleLogin HandleLogin
func (h *OauthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	//s.InitSessionStore(w, r)
	h.authorize(w, r)
}

// HandleLogout handler
func (h *OauthHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	h.token = nil
	cookie := &http.Cookie{
		Name:   "goauth2-ui",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	cookie2 := &http.Cookie{
		Name:   "goauth2",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie2)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *OauthHandler) authorize(w http.ResponseWriter, r *http.Request) bool {
	h.Log.Debug("in authorize")

	var a oauth2.AuthCodeAuthorize
	a.ClientID = h.ClientCreds.AuthCodeClient // h.getAuthCodeClient()
	a.OauthHost = h.OauthHost                 // getOauthRedirectHost()
	a.RedirectURI = h.getRedirectURI(r, authCodeRedirectURI)
	a.Scope = "write"
	a.State = h.ClientCreds.AuthCodeState // authCodeState
	a.Res = w
	a.Req = r

	h.Log.Debug("a: ", a)
	resp := a.AuthCodeAuthorizeUser()

	h.Log.Debug("Resp: ", resp)
	return resp
}

//HandleToken HandleToken
func (h *OauthHandler) HandleToken(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	h.Log.Debug("handle token")
	if state == h.ClientCreds.AuthCodeState {

		h.Auth.SetOauthHost(h.OauthHost)
		h.Auth.SetClientID(h.ClientCreds.AuthCodeClient)
		h.Auth.SetSecret(h.ClientCreds.AuthCodeSecret)
		h.Auth.SetCode(code)
		h.Auth.SetRedirectURI(h.getRedirectURI(r, authCodeRedirectURI))

		h.Log.Debug("getting token")

		resp := h.Auth.AuthCodeToken()
		h.Log.Debug("token resp: ", *resp)

		h.Log.Debug("token len: ", len(resp.AccessToken))

		h.Log.Debug("token : ", resp.AccessToken)
		if resp != nil && resp.AccessToken != "" {

			s, suc := h.getSession(r)
			if suc {
				h.Log.Debug("userLoggenIn : ", true)
				s.Values["userLoggenIn"] = true

				h.token = resp

				err := s.Save(r, w)
				h.Log.Debug(err)
				http.Redirect(w, r, "/clients", http.StatusFound)
			}
		}
	}
}
