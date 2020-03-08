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

// type userPage struct {
// 	ClientActive string
// 	OauthActive  string
// 	GwActive     string
// 	UserList     *[]services.User
// 	User         *services.User
// 	Client       *services.Client
// }

// user handlers-----------------------------------------------------

//HandleUsers HandleUsers
func (h *OauthHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("users Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {

			s.Values["userLoggenIn"] = true
			vars := mux.Vars(r)
			clientID := vars["clientId"]
			h.Log.Debug("users client: ", clientID)

			if clientID != "" {

				h.Service.SetToken(h.token.AccessToken)

				res, _ := h.Service.GetClient(clientID)
				h.Log.Debug("users client res: ", *res)

				res2, _ := h.Service.SearchUserList(clientID)

				res3, _ := h.Service.GetRoleList()
				h.Log.Debug("users client re3s: ", *res3)

				var page oauthPage
				page.OauthActive = "active"
				page.Client = res
				page.UserList = res2
				page.UserRoleList = res3
				var sm secSideMenu
				sm.UsersActive = "active teal"
				page.SecSideMenu = &sm
				h.Templates.ExecuteTemplate(w, "users.html", &page)
			}
		}
	}
}

// func handleAddUser(w http.ResponseWriter, r *http.Request) {
// 	s.InitSessionStore(w, r)
// 	session, err := s.GetSession(r)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	loggedIn := session.Values["userLoggenIn"]
// 	token := getToken(w, r)
// 	fmt.Print("Logged in: ")
// 	fmt.Println(loggedIn)
// 	//fmt.Println(token.AccessToken)
// 	//var res *[]services.Client
// 	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
// 		authorize(w, r)
// 	} else {
// 		templates.ExecuteTemplate(w, "addClient.html", nil)
// 	}

// }

//HandleNewUser HandleNewUser
func (h *OauthHandler) HandleNewUser(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("user new Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			var uu services.User
			clientID := r.FormValue("clientId")
			h.Log.Debug("user new client: ", clientID)

			clientIDD, _ := strconv.ParseInt(clientID, 10, 0)
			uu.ClientID = clientIDD

			username := r.FormValue("username")
			h.Log.Debug("user new username: ", username)

			uu.Username = username

			userRoleIDStr := r.FormValue("userRoleId")
			h.Log.Debug("user new userRoleIDStr: ", userRoleIDStr)
			userRoleID, _ := strconv.ParseInt(userRoleIDStr, 10, 0)

			uu.RoleID = userRoleID

			firstName := r.FormValue("firstName")
			h.Log.Debug("user new firstName: ", firstName)

			uu.FirstName = firstName

			lastName := r.FormValue("lastName")
			h.Log.Debug("user new lastName: ", lastName)

			uu.LastName = lastName

			emailAddress := r.FormValue("emailAddress")
			h.Log.Debug("user new emailAddress: ", emailAddress)

			uu.EmailAddress = emailAddress

			password := r.FormValue("password")
			h.Log.Debug("user new password: ", password)

			uu.Password = password

			enabled := r.FormValue("enabled")
			h.Log.Debug("user new enabled: ", enabled)

			if enabled == "yes" {
				uu.Enabled = true
			}

			h.Service.SetToken(h.token.AccessToken)
			res := h.Service.AddUser(&uu)
			h.Log.Debug("user new res: ", *res)
			http.Redirect(w, r, "/users/"+clientID, http.StatusFound)
		}
	}
}

//HandleEditUser HandleEditUser
func (h *OauthHandler) HandleEditUser(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("edit user Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			euvars := mux.Vars(r)
			username := euvars["username"]
			clientID := euvars["clientId"]
			h.Log.Debug("edit user client: ", clientID)
			h.Log.Debug("edit user username: ", username)

			if clientID != "" {

				h.Service.SetToken(h.token.AccessToken)

				res, _ := h.Service.GetClient(clientID)
				h.Log.Debug("edit user client res: ", *res)

				res2, _ := h.Service.GetUser(username, clientID)
				h.Log.Debug("edit user res2: ", *res2)

				res3, _ := h.Service.GetRoleList()
				h.Log.Debug("edit user role res3: ", *res3)

				var page oauthPage
				page.OauthActive = "active"
				page.Client = res
				page.User = res2
				page.UserRoleList = res3
				page.UserAssignedRole = res2.RoleID

				var sm secSideMenu
				//sm.UsersActive = "active teal"
				page.SecSideMenu = &sm
				h.Templates.ExecuteTemplate(w, "editUser.html", &page)
			}
		}
	}
}

//HandleUpdateUserInfo HandleUpdateUserInfo
func (h *OauthHandler) HandleUpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("user update info Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			var uu services.UserInfo
			clientID := r.FormValue("clientId")
			h.Log.Debug("user update info client: ", clientID)

			clientIDD, _ := strconv.ParseInt(clientID, 10, 0)
			uu.ClientID = clientIDD

			username := r.FormValue("username")
			h.Log.Debug("user update info username: ", username)
			uu.Username = username

			userRoleIDStr := r.FormValue("userRoleId")
			userRoleID, _ := strconv.ParseInt(userRoleIDStr, 10, 0)
			h.Log.Debug("user update info userRoleIDStr: ", userRoleIDStr)
			uu.RoleID = userRoleID

			firstName := r.FormValue("firstName")
			h.Log.Debug("user update info firstName: ", firstName)
			uu.FirstName = firstName

			lastName := r.FormValue("lastName")
			h.Log.Debug("user update info lastName: ", lastName)
			uu.LastName = lastName

			emailAddress := r.FormValue("emailAddress")
			h.Log.Debug("user update info emailAddress: ", emailAddress)
			uu.EmailAddress = emailAddress

			h.Service.SetToken(h.token.AccessToken)
			res := h.Service.UpdateUser(&uu)
			h.Log.Debug("user update info res: ", *res)

			http.Redirect(w, r, "/editUser/"+username+"/"+clientID, http.StatusFound)
		}
	}
}

//HandleUpdateUserEnable HandleUpdateUserEnable
func (h *OauthHandler) HandleUpdateUserEnable(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("user update enabled Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			var uu services.UserDis
			clientID := r.FormValue("clientId")
			h.Log.Debug("user update enabled client: ", clientID)
			clientIDD, _ := strconv.ParseInt(clientID, 10, 0)
			uu.ClientID = clientIDD

			username := r.FormValue("username")
			h.Log.Debug("user update enabled username: ", username)
			uu.Username = username

			enabled := r.FormValue("enabled")
			h.Log.Debug("user update enabled: ", enabled)
			if enabled == "yes" {
				uu.Enabled = true
			}

			h.Service.SetToken(h.token.AccessToken)

			res := h.Service.UpdateUser(&uu)
			h.Log.Debug("user update enbaled res: ", *res)

			http.Redirect(w, r, "/editUser/"+username+"/"+clientID, http.StatusFound)
		}

	}
}

//HandleUpdateUserPw HandleUpdateUserPw
func (h *OauthHandler) HandleUpdateUserPw(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("user update pw Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			var uu services.UserPW
			clientID := r.FormValue("clientId")
			h.Log.Debug("user update pw client: ", clientID)
			clientIDD, _ := strconv.ParseInt(clientID, 10, 0)
			uu.ClientID = clientIDD

			username := r.FormValue("username")
			h.Log.Debug("user update pw username: ", username)
			uu.Username = username

			password := r.FormValue("password")
			h.Log.Debug("user update pw password: ", password)
			uu.Password = password

			h.Service.SetToken(h.token.AccessToken)

			res := h.Service.UpdateUser(&uu)
			h.Log.Debug("user update pw res: ", *res)

			http.Redirect(w, r, "/editUser/"+username+"/"+clientID, http.StatusFound)
		}
	}
}
