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

func (h *OauthHandler) handleGrantType(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("grant type Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			vars := mux.Vars(r)
			clientID := vars["clientId"]
			h.Log.Debug("grant type clientid: ", clientID)
			if clientID != "" {
				h.Service.SetToken(h.token.AccessToken)
				res, _ := h.Service.GetClient(clientID)
				h.Log.Debug("grant type client res: ", *res)
				//fmt.Println(res)
				var page oauthPage
				page.OauthActive = "active"
				page.Client = res
				res2, _ := h.Service.GetGrantTypeList(clientID)
				h.Log.Debug("grant type gt res: ", *res2)
				page.GrantTypes = res2
				var sm secSideMenu
				sm.GrantTypeActive = "active teal"
				page.SecSideMenu = &sm

				h.Templates.ExecuteTemplate(w, "grantTypes.html", &page)
			}
		}
	}
}

func (h *OauthHandler) handleGrantTypeAdd(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("clients add Logged in: ", loggedIn)
		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			grantType := r.FormValue("grantType")
			h.Log.Debug("grant type add granttype: ", grantType)

			clientIDStr := r.FormValue("clientId")
			clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
			h.Log.Debug("grant type add clientid: ", clientID)

			s.Values["userLoggenIn"] = true

			h.Service.SetToken(h.token.AccessToken)

			if grantType != "" {
				var gg services.GrantType
				gg.ClientID = clientID
				gg.GrantType = grantType
				gres := h.Service.AddGrantType(&gg)
				h.Log.Debug("grant type add granttype res: ", *gres)
			}
			http.Redirect(w, r, "/clientGrantTypes/"+clientIDStr, http.StatusFound)
		}
	}
}

func (h *OauthHandler) handleGrantTypeDelete(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("grantType delete Logged in: ", loggedIn)
		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			vars := mux.Vars(r)
			ID := vars["id"]
			clientID := vars["clientId"]

			h.Log.Debug("grantType delete id: ", ID)

			if ID != "" && clientID != "" {
				h.Service.SetToken(h.token.AccessToken)
				gres := h.Service.DeleteGrantType(ID)
				h.Log.Debug("grantType delete res: ", *gres)
			}
			http.Redirect(w, r, "/clientGrantTypes/"+clientID, http.StatusFound)
		}
	}
}
