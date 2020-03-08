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

type allowedURIDisplay struct {
	ID           int64
	URI          string
	ClientID     int64
	AssignedRole int64
}

//HandleAllowedUris HandleAllowedUris
func (h *OauthHandler) HandleAllowedUris(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("allowed uri Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			vars := mux.Vars(r)
			clientID := vars["clientId"]
			h.Log.Debug("clientID: ", clientID)
			if clientID != "" {

				res, _ := h.Service.GetClient(clientID)
				h.Log.Debug("client res: ", *res)

				var page oauthPage
				page.OauthActive = "active"
				page.Client = res
				rr, _ := h.Service.GetClientRoleList(clientID)
				h.Log.Debug("rr: ", *rr)

				page.ClientRoles = rr

				rMap := make(map[int64]int64)

				for _, rrr := range *rr {
					ruu, _ := h.Service.GetRoleURIList(strconv.FormatInt(rrr.ID, 10))
					h.Log.Debug("ruu: ", *ruu)

					for _, rui := range *ruu {
						h.Log.Debug("rui: ", rui)
						rMap[rui.ClientAllowedURIID] = rui.ClientRoleID
					}
				}
				ares, _ := h.Service.GetAllowedURIList(clientID)
				h.Log.Debug("ares: ", *ares)

				var newAres []allowedURIDisplay

				for _, ar := range *ares {
					var ard allowedURIDisplay
					ard.AssignedRole = rMap[ar.ID]
					ard.ClientID = ar.ClientID
					ard.ID = ar.ID
					ard.URI = ar.URI
					newAres = append(newAres, ard)
					h.Log.Debug("role: ", ar)
				}
				h.Log.Debug("roles: ", newAres)

				page.AllowedURIs = &newAres
				var sm secSideMenu
				sm.AllowedURIActive = "active teal"
				page.SecSideMenu = &sm

				h.Templates.ExecuteTemplate(w, "allowedUris.html", &page)
			}
		}
	}
}

//HandleAllowedUrisAdd HandleAllowedUrisAdd
func (h *OauthHandler) HandleAllowedUrisAdd(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("allowed uri add Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			uri := r.FormValue("uri")
			h.Log.Debug("uri: ", uri)

			roleIDStr := r.FormValue("roleId")
			h.Log.Debug("roleIDStr: ", roleIDStr)

			roleID, _ := strconv.ParseInt(roleIDStr, 10, 0)

			clientIDStr := r.FormValue("clientId")
			clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
			h.Log.Debug("clientId: ", clientID)

			if roleIDStr != "" && clientIDStr != "" {

				var auu services.AllowedURI
				auu.ClientID = clientID
				auu.URI = uri
				aures := h.Service.AddAllowedURI(&auu)
				h.Log.Debug("aures: ", *aures)
				if aures.Success == true {
					var crr services.RoleURI
					crr.ClientRoleID = roleID
					crr.ClientAllowedURIID = aures.ID
					crres := h.Service.AddRoleURI(&crr)
					h.Log.Debug("crres: ", *crres)
				}
			}
			http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
		}
	}
}

//HandleAllowedUrisUpdate HandleAllowedUrisUpdate
func (h *OauthHandler) HandleAllowedUrisUpdate(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("allowed uri update Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true

			IDStr := r.FormValue("id")
			h.Log.Debug("IDStr: ", IDStr)

			ID, _ := strconv.ParseInt(IDStr, 10, 0)
			h.Log.Debug("ID: ", ID)

			uri := r.FormValue("uri")
			h.Log.Debug("uri: ", uri)

			roleIDStr := r.FormValue("roleId")
			h.Log.Debug("roleIDStr: ", roleIDStr)

			roleID, _ := strconv.ParseInt(roleIDStr, 10, 0)
			h.Log.Debug("roleID: ", roleID)

			originalRoleIDStr := r.FormValue("originalRoleId")
			h.Log.Debug("originalRoleIDStr: ", originalRoleIDStr)

			originalRoleID, _ := strconv.ParseInt(originalRoleIDStr, 10, 0)
			h.Log.Debug("originalRoleID: ", originalRoleID)

			var updateRole = false
			if roleID != originalRoleID {
				updateRole = true
			}
			h.Log.Debug("updateRole: ", updateRole)

			clientIDStr := r.FormValue("clientId")
			clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
			h.Log.Debug("clientIDStr: ", clientIDStr)
			h.Log.Debug("clientID: ", clientID)

			var auu services.AllowedURI
			auu.ID = ID
			auu.ClientID = clientID
			auu.URI = uri
			aures := h.Service.UpdateAllowedURI(&auu)
			h.Log.Debug("aures: ", *aures)
			if aures.Success == true {
				if updateRole == true {

					var crr services.RoleURI
					crr.ClientRoleID = originalRoleID
					crr.ClientAllowedURIID = ID
					h.Service.DeleteRoleURI(&crr)

					crr.ClientRoleID = roleID
					crr.ClientAllowedURIID = ID
					crres := h.Service.AddRoleURI(&crr)
					h.Log.Debug("crres in update: ", *crres)
				}
			}
			http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
		}
	}
}

//HandleAllowedUrisDelete HandleAllowedUrisDelete
func (h *OauthHandler) HandleAllowedUrisDelete(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("allowed uri delete Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true

			vars := mux.Vars(r)

			IDStr := vars["id"]
			h.Log.Debug("IDStr: ", IDStr)

			ID, _ := strconv.ParseInt(IDStr, 10, 0)
			h.Log.Debug("ID: ", ID)

			clientIDStr := vars["clientId"]
			clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
			h.Log.Debug("clientIDStr: ", clientIDStr)
			h.Log.Debug("clientID: ", clientID)

			roleIDStr := vars["roleId"]
			h.Log.Debug("roleIDStr: ", roleIDStr)

			roleID, _ := strconv.ParseInt(roleIDStr, 10, 0)
			h.Log.Debug("roleID: ", roleID)

			if IDStr != "" && clientIDStr != "" && roleIDStr != "" {

				var crr services.RoleURI
				crr.ClientRoleID = roleID
				crr.ClientAllowedURIID = ID
				h.Service.DeleteRoleURI(&crr)

				aures := h.Service.DeleteAllowedURI(IDStr)
				h.Log.Debug("aures: ", *aures)

				http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
			}
		}
	}
}
