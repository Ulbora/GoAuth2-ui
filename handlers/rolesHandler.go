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

//HandleRoles HandleRoles
func (h *OauthHandler) HandleRoles(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("roles Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			rvars := mux.Vars(r)
			clientID := rvars["clientId"]
			h.Log.Debug("roles clientID: ", clientID)
			if clientID != "" {
				h.Service.SetToken(h.token.AccessToken)

				res, _ := h.Service.GetClient(clientID)
				h.Log.Debug("roles client res: ", *res)

				//fmt.Println(res)
				var page oauthPage
				page.OauthActive = "active"
				page.Client = res

				res2, _ := h.Service.GetClientRoleList(clientID)
				h.Log.Debug("roles client res2: ", *res2)
				for rrri, rrr := range *res2 {
					ruu, _ := h.Service.GetRoleURIList(strconv.FormatInt(rrr.ID, 10))
					h.Log.Debug("role urls: ", *ruu)
					for _, rru := range *ruu {
						if rru.ClientRoleID == rrr.ID {
							(*res2)[rrri].Used = true
							h.Log.Debug("role used: ", (*res2)[rrri])
						}
					}
				}
				// for _, rrr := range *res2 {
				// 	h.Log.Debug("test role: ", rrr)
				// }

				page.ClientRoles = res2
				if h.ClientCreds.AuthCodeClient == clientID {
					page.ClientIsSelf = true
				}

				var sm secSideMenu
				sm.RolesActive = "active teal"
				page.SecSideMenu = &sm

				h.Templates.ExecuteTemplate(w, "roles.html", &page)
			}
		}
	}
}

//HandleRoleAdd HandleRoleAdd
func (h *OauthHandler) HandleRoleAdd(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("roles add Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			clientRole := r.FormValue("clientRole")
			h.Log.Debug("roles add clientrole: ", clientRole)

			clientIDStr := r.FormValue("clientId")
			clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
			h.Log.Debug("roles add clientID: ", clientID)

			s.Values["userLoggenIn"] = true

			h.Service.SetToken(h.token.AccessToken)

			resTest, _ := h.Service.GetClientRoleList(clientIDStr)
			h.Log.Debug("roles add resTest: ", *resTest)
			var roleExists = false
			for _, rl := range *resTest {
				if rl.Role == clientRole {
					roleExists = true
					break
				}
			}
			h.Log.Debug("roles add roleExists: ", roleExists)
			if clientRole != "" && roleExists != true {
				var rr services.ClientRole
				rr.ClientID = clientID
				rr.Role = clientRole
				rres := h.Service.AddClientRole(&rr)
				h.Log.Debug("roles add rres: ", *rres)
			}
			http.Redirect(w, r, "/clientRoles/"+clientIDStr, http.StatusFound)
		}
	}
}

//HandleRoleDelete HandleRoleDelete
func (h *OauthHandler) HandleRoleDelete(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("roles delete Logged in: ", loggedIn)
		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			rdvars := mux.Vars(r)
			ID := rdvars["id"]
			clientID := rdvars["clientId"]
			h.Log.Debug("roles delete id: ", ID)
			h.Log.Debug("roles delete clientID: ", clientID)

			if ID != "" && clientID != "" {
				h.Service.SetToken(h.token.AccessToken)

				rres := h.Service.DeleteClientRole(ID)
				h.Log.Debug("roles delete rres: ", *rres)
			}
			http.Redirect(w, r, "/clientRoles/"+clientID, http.StatusFound)
		}
	}
}
