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

//RedirectURIService RedirectURIService
type RedirectURIService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

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
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		r.Log.Debug("Add redirect req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			req.Header.Set("apiKey", r.APIKey)
			_, code := r.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetRedirectURIList get GetRedirectURIList list
func (r *Oauth2Service) GetRedirectURIList(clientID string) *[]RedirectURI {
	var rtn = make([]RedirectURI, 0)
	var gURL = r.Host + "/rs/clientRedirectUri/list/" + clientID
	req, rErr := http.NewRequest("GET", gURL, nil)
	r.Log.Debug("Get redirect req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", r.ClientID)
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("apiKey", r.APIKey)
		r.Proxy.Do(req, &rtn)
	}
	return &rtn
}

// DeleteRedirectURI delete DeleteRedirectURI
func (r *Oauth2Service) DeleteRedirectURI(id string) *RedirectURIResponse {
	var rtn = new(RedirectURIResponse)
	var gURL = r.Host + "/rs/clientRedirectUri/delete/" + id
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	r.Log.Debug("delete redirect req: ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("clientId", r.ClientID)
		req.Header.Set("apiKey", r.APIKey)
		_, code := r.Proxy.Do(req, &rtn)
		rtn.Code = code
	}
	return rtn
}
