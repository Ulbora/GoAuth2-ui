//Package handlers ...
package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

type clientPage struct {
	ClientActive string
	OauthActive  string
	GwActive     string
	ClientList   *[]services.Client
	Client       *services.Client
}

// user handlers-----------------------------------------------------

//HandleClients HandleClients
func (h *OauthHandler) HandleClients(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("clients Logged in: ", loggedIn)

		var res *[]services.Client
		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			clientName := r.FormValue("clientName")
			h.Log.Debug("client name in HandleClients: ", clientName)

			if clientName != "" {
				h.Service.SetToken(h.token.AccessToken)

				var cc services.Client
				cc.Name = clientName

				res, _ = h.Service.SearchClient(&cc)
			}
			var page clientPage
			page.ClientActive = "active"
			page.ClientList = res
			h.Templates.ExecuteTemplate(w, "clients.html", &page)
		}
	}
}

//HandleAddClient HandleAddClient
func (h *OauthHandler) HandleAddClient(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("Add client Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			h.Templates.ExecuteTemplate(w, "addClient.html", nil)
		}
	}
}

//HandleEditClient HandleEditClient
func (h *OauthHandler) HandleEditClient(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("edit client Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			s.Values["userLoggenIn"] = true
			vars := mux.Vars(r)
			clientID := vars["clientId"]

			if clientID != "" {
				h.Service.SetToken(h.token.AccessToken)
				res, _ := h.Service.GetClient(clientID)
				var page oauthPage
				page.ClientActive = "active"
				page.Client = res
				h.Log.Debug("AuthCodeClient: ", h.ClientCreds.AuthCodeClient)
				h.Log.Debug("clientID: ", clientID)
				if h.ClientCreds.AuthCodeClient == clientID {
					page.ClientIsSelf = true
				}
				h.Templates.ExecuteTemplate(w, "editClient.html", &page)
			}
		}
	}
}

//HandleNewClient HandleNewClient
func (h *OauthHandler) HandleNewClient(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("new client Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			clientName := r.FormValue("clientName")
			h.Log.Debug("clientName: ", clientName)

			webSite := r.FormValue("webSite")
			h.Log.Debug("webSite: ", webSite)

			emailAddress := r.FormValue("emailAddress")
			h.Log.Debug("emailAddress: ", emailAddress)

			redirectURLStr := r.FormValue("redirectURLs")
			redirectURLStr = strings.Replace(redirectURLStr, " ", "", -1)
			h.Log.Debug("redirectURLStr: ", redirectURLStr)

			redirectURLs := strings.Split(redirectURLStr, ",")
			h.Log.Debug("redirectURLs: ", redirectURLs)

			enabled := r.FormValue("enabled")
			h.Log.Debug("enabled: ", enabled)

			var cc services.Client
			cc.Name = clientName
			cc.Email = emailAddress
			if enabled == "yes" {
				cc.Enabled = true
			} else {
				cc.Enabled = false
			}
			cc.WebSite = webSite

			var uris []services.RedirectURI
			for i := range redirectURLs {
				var uri services.RedirectURI
				uri.URI = redirectURLs[i]
				uris = append(uris, uri)
			}
			cc.RedirectURIs = uris
			h.Service.SetToken(h.token.AccessToken)

			res := h.Service.AddClient(&cc)
			if res.Success == true {
				http.Redirect(w, r, "/clients", http.StatusFound)
			} else {
				http.Redirect(w, r, "/addClient", http.StatusFound)
			}
		}
	}
}

//HandleUpdateClient HandleUpdateClient
func (h *OauthHandler) HandleUpdateClient(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("update client Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			clientIDStr := r.FormValue("clientId")
			clientID, errID := strconv.ParseInt(clientIDStr, 10, 0)
			if errID != nil {
				h.Log.Debug("errID: ", errID)

			}
			h.Log.Debug("clientID: ", clientID)

			clientName := r.FormValue("clientName")
			h.Log.Debug("clientName: ", clientName)

			secret := r.FormValue("secret")
			h.Log.Debug("secret: ", secret)

			webSite := r.FormValue("webSite")
			h.Log.Debug("webSite: ", webSite)

			emailAddress := r.FormValue("emailAddress")
			h.Log.Debug("emailAddress: ", emailAddress)

			enabled := r.FormValue("enabled")
			h.Log.Debug("enabled: ", enabled)

			var cc services.Client
			cc.ClientID = clientID
			cc.Secret = secret
			cc.Name = clientName
			cc.Email = emailAddress
			if enabled == "yes" {
				cc.Enabled = true
			} else {
				cc.Enabled = false
			}
			cc.WebSite = webSite

			h.Service.SetToken(h.token.AccessToken)

			res := h.Service.UpdateClient(&cc)
			if res.Success == true {
				http.Redirect(w, r, "/clients", http.StatusFound)
			} else {
				fmt.Println(res)
				http.Redirect(w, r, "/editClient/"+clientIDStr, http.StatusFound)
			}
		}
	}
}
