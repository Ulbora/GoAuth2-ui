//Package handlers ...
package handlers

import (
	"net/http"

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

type secSideMenu struct {
	RedirectURLActive string
	GrantTypeActive   string
	RolesActive       string
	AllowedURIActive  string
	ClientActive      string
	UlboraURIsActive  string
	UsersActive       string
}

type oauthPage struct {
	ClientActive         string
	OauthActive          string
	GwActive             string
	CanDeleteRedirectURI bool
	ClientIsSelf         bool
	SecSideMenu          *secSideMenu
	ClientList           *[]services.Client
	Client               *services.Client
	RedirectURLs         *[]services.RedirectURI
	GrantTypes           *[]services.GrantType
	ClientRoles          *[]services.ClientRole
	AllowedURIs          *[]allowedURIDisplay
	RoleURIs             *[]services.RoleURI
	UserList             *[]services.User
	User                 *services.User
	UserRoleList         *[]services.Role
	UserAssignedRole     int64
}

//HandleOauth2 HandleOauth2
func (h *OauthHandler) HandleOauth2(w http.ResponseWriter, r *http.Request) {

	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token

		h.Log.Debug("in main page. Logged in: ", loggedIn)
		h.Log.Debug("in main page. token: ", token)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			h.Log.Debug("userLoggenIn in oauth2: ", true)
			vars := mux.Vars(r)
			clientID := vars["clientId"]

			if clientID != "" {
				h.Log.Debug("ClientId in HandleOauth2: ", clientID)
				h.Service.SetToken(h.token.AccessToken)
				res, _ := h.Service.GetClient(clientID)

				var page oauthPage
				page.OauthActive = "active"
				page.Client = res
				var sm secSideMenu
				sm.ClientActive = "active teal"
				page.SecSideMenu = &sm

				h.Templates.ExecuteTemplate(w, "oauth2.html", &page)
			}
		}
	}
}
