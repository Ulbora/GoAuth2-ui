//Package services ...
package services

import (
	"bytes"
	"encoding/json"
	"net/http"
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

//ClientRoleService ClientRoleService
type ClientRoleService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//ClientRole ClientRole
type ClientRole struct {
	ID       int64  `json:"id"`
	ClientID int64  `json:"clientId"`
	Role     string `json:"role"`
}

//ClientRoleResponse resp
type ClientRoleResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddClientRole AddClientRole
func (r *Oauth2Service) AddClientRole(cr *ClientRole) *ClientRoleResponse {
	var rtn = new(ClientRoleResponse)
	var addURL = r.Host + "/rs/clientRoleSuper/add"
	aJSON, err := json.Marshal(cr)
	r.Log.Debug("Add ClientRole: ", err)
	if err == nil {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		r.Log.Debug("Add client Role req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			//req.Header.Set("apiKey", r.APIKey)
			_, code := r.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetClientRoleList get GetClientRoleList list
func (r *Oauth2Service) GetClientRoleList(clientID string) (*[]ClientRole, int) {
	var rtn = make([]ClientRole, 0)
	var code int
	var gURL = r.Host + "/rs/clientRole/list/" + clientID
	req, rErr := http.NewRequest("GET", gURL, nil)
	r.Log.Debug("Get client role list req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", r.ClientID)
		req.Header.Set("Authorization", "Bearer "+r.Token)
		//req.Header.Set("apiKey", r.APIKey)
		_, code = r.Proxy.Do(req, &rtn)
	}
	return &rtn, code
}

// DeleteClientRole delete DeleteClientRole
func (r *Oauth2Service) DeleteClientRole(id string) *ClientRoleResponse {
	var rtn = new(ClientRoleResponse)
	var gURL = r.Host + "/rs/clientRole/delete/" + id
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	r.Log.Debug("Delete client: ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("clientId", r.ClientID)
		//req.Header.Set("apiKey", r.APIKey)
		_, code := r.Proxy.Do(req, &rtn)
		rtn.Code = code
	}
	return rtn
}
