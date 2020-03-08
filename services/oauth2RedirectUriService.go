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

//RedirectURI RedirectURI
type RedirectURI struct {
	ID       int64  `json:"id"`
	ClientID int64  `json:"clientId"`
	URI      string `json:"uri"`
}

//RedirectURIResponse resp
type RedirectURIResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddRedirectURI AddRedirectURI
func (r *Oauth2Service) AddRedirectURI(rd *RedirectURI) *RedirectURIResponse {
	var rtn = new(RedirectURIResponse)
	var addURL = r.Host + "/rs/clientRedirectUri/add"
	aJSON, err := json.Marshal(rd)
	r.Log.Debug("Add redirect: ", err)
	if err == nil {
		reqr, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		r.Log.Debug("Add redirect req: ", rErr)
		if rErr == nil {
			reqr.Header.Set("Content-Type", "application/json")
			reqr.Header.Set("Authorization", "Bearer "+r.Token)
			reqr.Header.Set("clientId", r.ClientID)
			//reqr.Header.Set("apiKey", r.APIKey)
			_, code := r.Proxy.Do(reqr, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetRedirectURIList get GetRedirectURIList list
func (r *Oauth2Service) GetRedirectURIList(clientID string) (*[]RedirectURI, int) {
	var rtn = make([]RedirectURI, 0)
	var code int
	var gURL = r.Host + "/rs/clientRedirectUri/list/" + clientID
	reqrg, rErr := http.NewRequest("GET", gURL, nil)
	r.Log.Debug("Get redirect req: ", rErr)
	if rErr == nil {
		reqrg.Header.Set("clientId", r.ClientID)
		reqrg.Header.Set("Authorization", "Bearer "+r.Token)
		//reqrg.Header.Set("apiKey", r.APIKey)
		_, code = r.Proxy.Do(reqrg, &rtn)
	}
	return &rtn, code
}

// DeleteRedirectURI delete DeleteRedirectURI
func (r *Oauth2Service) DeleteRedirectURI(id string) *RedirectURIResponse {
	var rtn = new(RedirectURIResponse)
	var gURL = r.Host + "/rs/clientRedirectUri/delete/" + id
	reqrd, rErr := http.NewRequest("DELETE", gURL, nil)
	r.Log.Debug("delete redirect req: ", rErr)
	if rErr == nil {
		reqrd.Header.Set("Content-Type", "application/json")
		reqrd.Header.Set("Authorization", "Bearer "+r.Token)
		reqrd.Header.Set("clientId", r.ClientID)
		//reqrd.Header.Set("apiKey", r.APIKey)
		_, code := r.Proxy.Do(reqrd, &rtn)
		rtn.Code = code
	}
	return rtn
}
