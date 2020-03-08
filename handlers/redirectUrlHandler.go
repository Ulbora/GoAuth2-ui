//Package handlers ...
package handlers

import (
	"net/http"
	"strconv"

	services "github.com/Ulbora/GoAuth2-ui/services"

	"github.com/gorilla/mux"
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

//HandleRedirectURLs HandleRedirectURLs
func (h *OauthHandler) HandleRedirectURLs(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("redurect uri Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			ruvars := mux.Vars(r)
			clientID := ruvars["clientId"]
			h.Log.Debug("redirect url client: ", clientID)

			if clientID != "" {
				h.Service.SetToken(h.token.AccessToken)
				res, _ := h.Service.GetClient(clientID)
				h.Log.Debug("redirect url client res: ", *res)
				var page oauthPage
				page.OauthActive = "active"
				page.Client = res

				res2, _ := h.Service.GetRedirectURIList(clientID)
				h.Log.Debug("redirect url list res: ", *res2)
				page.RedirectURLs = res2
				if len(*res2) > 1 {
					page.CanDeleteRedirectURI = true
				}
				var sm secSideMenu
				sm.RedirectURLActive = "active teal"
				page.SecSideMenu = &sm

				h.Templates.ExecuteTemplate(w, "redirectUrls.html", &page)
			}
		}
	}
}

//HandleRedirectURLAdd HandleRedirectURLAdd
func (h *OauthHandler) HandleRedirectURLAdd(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("redurect uri add Logged in: ", loggedIn)
		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			redirectURL := r.FormValue("redirectURL")
			h.Log.Debug("redurect uri add: ", redirectURL)

			clientIDStr := r.FormValue("clientId")
			clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
			h.Log.Debug("redurect uri add clientId: ", clientID)

			s.Values["userLoggenIn"] = true

			h.Service.SetToken(h.token.AccessToken)

			var rr services.RedirectURI
			rr.ClientID = clientID
			rr.URI = redirectURL
			ares := h.Service.AddRedirectURI(&rr)
			h.Log.Debug("redurect uri add ares: ", *ares)

			http.Redirect(w, r, "/clientRedirectUrls/"+clientIDStr, http.StatusFound)
		}
	}
}

//HandleRedirectURLDelete HandleRedirectURLDelete
func (h *OauthHandler) HandleRedirectURLDelete(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("redurect uri delete Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			rudvars := mux.Vars(r)
			ID := rudvars["id"]
			clientID := rudvars["clientId"]

			h.Log.Debug("redurect uri delete id: ", ID)
			h.Log.Debug("redurect uri delete clientID: ", clientID)

			if ID != "" && clientID != "" {
				h.Service.SetToken(h.token.AccessToken)

				dres := h.Service.DeleteRedirectURI(ID)
				h.Log.Debug("redurect uri delete dres: ", *dres)
			}
			http.Redirect(w, r, "/clientRedirectUrls/"+clientID, http.StatusFound)
		}
	}
}
