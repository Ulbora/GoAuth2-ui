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

// //AllowedURIService service
// type AllowedURIService struct {
// 	Token    string
// 	ClientID string
// 	APIKey   string
// 	UserID   string
// 	Hashed   string
// 	Host     string
// }

//AllowedURI AllowedURI
type AllowedURI struct {
	ID       int64  `json:"id"`
	URI      string `json:"uri"`
	ClientID int64  `json:"clientId"`
}

//AllowedURIResponse resp
type AllowedURIResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddAllowedURI add
func (a *Oauth2Service) AddAllowedURI(au *AllowedURI) *AllowedURIResponse {
	var rtn = new(AllowedURIResponse)
	var addURL = a.Host + "/rs/clientAllowedUriSuper/add"
	aJSON, err := json.Marshal(au)
	a.Log.Debug("Add  allowed uri: ", err)
	if err == nil {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		a.Log.Debug("Add allowed uri req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+a.Token)
			req.Header.Set("clientId", a.ClientID)
			//req.Header.Set("apiKey", a.APIKey)
			_, code := a.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

//UpdateAllowedURI update UpdateAllowedURI
func (a *Oauth2Service) UpdateAllowedURI(au *AllowedURI) *AllowedURIResponse {
	var rtn = new(AllowedURIResponse)
	var upURL = a.Host + "/rs/clientAllowedUriSuper/update"
	aJSON, err := json.Marshal(au)
	a.Log.Debug("update allowed uri: ", err)
	if err == nil {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		a.Log.Debug("update allowed uri req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+a.Token)
			req.Header.Set("clientId", a.ClientID)
			//req.Header.Set("apiKey", a.APIKey)
			_, code := a.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetAllowedURI get GetAllowedURI
func (a *Oauth2Service) GetAllowedURI(id string) (*AllowedURI, int) {
	var rtn = new(AllowedURI)
	var code int
	var gURL = a.Host + "/rs/clientAllowedUri/get/" + id
	req, rErr := http.NewRequest("GET", gURL, nil)
	a.Log.Debug("get allowed uri req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", a.ClientID)
		req.Header.Set("Authorization", "Bearer "+a.Token)
		//req.Header.Set("apiKey", a.APIKey)
		_, code = a.Proxy.Do(req, &rtn)
	}
	return rtn, code
}

// GetAllowedURIList get GetAllowedURIList list
func (a *Oauth2Service) GetAllowedURIList(clientID string) (*[]AllowedURI, int) {
	var rtn = make([]AllowedURI, 0)
	var code int
	var gURL = a.Host + "/rs/clientAllowedUri/list/" + clientID
	req, rErr := http.NewRequest("GET", gURL, nil)
	a.Log.Debug("get allowed uri list req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", a.ClientID)
		req.Header.Set("Authorization", "Bearer "+a.Token)
		//req.Header.Set("apiKey", a.APIKey)
		_, code = a.Proxy.Do(req, &rtn)
	}
	return &rtn, code
}

// DeleteAllowedURI delete DeleteAllowedURI
func (a *Oauth2Service) DeleteAllowedURI(id string) *AllowedURIResponse {
	var rtn = new(AllowedURIResponse)
	var gURL = a.Host + "/rs/clientAllowedUri/delete/" + id
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	a.Log.Debug("delete allowed uri req: ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+a.Token)
		req.Header.Set("clientId", a.ClientID)
		//req.Header.Set("apiKey", a.APIKey)
		_, code := a.Proxy.Do(req, &rtn)
		rtn.Code = code
	}
	return rtn
}
