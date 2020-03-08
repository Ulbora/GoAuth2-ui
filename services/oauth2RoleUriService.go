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

//RoleURI RoleURI
type RoleURI struct {
	ClientRoleID       int64 `json:"clientRoleId"`
	ClientAllowedURIID int64 `json:"clientAllowedUriId"`
}

//RoleURIResponse resp
type RoleURIResponse struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
}

//AddRoleURI add
func (r *Oauth2Service) AddRoleURI(ru *RoleURI) *RoleURIResponse {
	var rtn = new(RoleURIResponse)
	var addURL = r.Host + "/rs/clientRoleUri/add"
	aJSON, err := json.Marshal(ru)
	r.Log.Debug("Add role uri: ", err)
	if err == nil {
		requr, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		r.Log.Debug("Add role uri req: ", rErr)
		if rErr == nil {
			requr.Header.Set("Content-Type", "application/json")
			requr.Header.Set("Authorization", "Bearer "+r.Token)
			requr.Header.Set("clientId", r.ClientID)
			//requr.Header.Set("apiKey", r.APIKey)
			_, code := r.Proxy.Do(requr, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetRoleURIList get GetRoleURIList list
func (r *Oauth2Service) GetRoleURIList(uID string) (*[]RoleURI, int) {
	var rtn = make([]RoleURI, 0)
	var code int
	var gURL = r.Host + "/rs/clientRoleUri/list/" + uID
	requrl, rErr := http.NewRequest("GET", gURL, nil)
	r.Log.Debug("get role uri list req: ", rErr)
	if rErr == nil {
		requrl.Header.Set("clientId", r.ClientID)
		requrl.Header.Set("Authorization", "Bearer "+r.Token)
		//requrl.Header.Set("apiKey", r.APIKey)
		_, code = r.Proxy.Do(requrl, &rtn)
	}
	return &rtn, code
}

//DeleteRoleURI add
func (r *Oauth2Service) DeleteRoleURI(ru *RoleURI) *RoleURIResponse {
	var rtn = new(RoleURIResponse)
	var uURL = r.Host + "/rs/clientRoleUri/delete"
	aJSON, err := json.Marshal(ru)
	r.Log.Debug("delete role uri: ", err)
	if err == nil {
		requrd, rErr := http.NewRequest("POST", uURL, bytes.NewBuffer(aJSON))
		r.Log.Debug("delete role uri req: ", rErr)
		if rErr == nil {
			requrd.Header.Set("Content-Type", "application/json")
			requrd.Header.Set("Authorization", "Bearer "+r.Token)
			requrd.Header.Set("clientId", r.ClientID)
			//requrd.Header.Set("apiKey", r.APIKey)
			_, code := r.Proxy.Do(requrd, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}
