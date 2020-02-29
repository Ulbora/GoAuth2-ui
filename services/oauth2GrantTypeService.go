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

//GrantTypeService GrantTypeService
type GrantTypeService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//GrantType GrantType
type GrantType struct {
	ID        int64  `json:"id"`
	ClientID  int64  `json:"clientId"`
	GrantType string `json:"grantType"`
}

//GrantTypeResponse resp
type GrantTypeResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddGrantType AddGrantType
func (g *Oauth2Service) AddGrantType(rd *GrantType) *GrantTypeResponse {
	var rtn = new(GrantTypeResponse)
	var addURL = g.Host + "/rs/clientGrantType/add"
	aJSON, err := json.Marshal(rd)
	g.Log.Debug("Add Grant Type: ", err)
	if err == nil {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		g.Log.Debug("Add grant type req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+g.Token)
			req.Header.Set("clientId", g.ClientID)
			req.Header.Set("apiKey", g.APIKey)
			_, code := g.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetGrantTypeList get GetGrantTypeList list
func (g *Oauth2Service) GetGrantTypeList(clientID string) *[]GrantType {
	var rtn = make([]GrantType, 0)
	var gURL = g.Host + "/rs/clientGrantType/list/" + clientID
	req, rErr := http.NewRequest("GET", gURL, nil)
	g.Log.Debug("get grant type list req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", g.ClientID)
		req.Header.Set("Authorization", "Bearer "+g.Token)
		req.Header.Set("apiKey", g.APIKey)
		g.Proxy.Do(req, &rtn)
	}
	return &rtn
}

// DeleteGrantType delete DeleteGrantType
func (g *Oauth2Service) DeleteGrantType(id string) *GrantTypeResponse {
	var rtn = new(GrantTypeResponse)
	var gURL = g.Host + "/rs/clientGrantType/delete/" + id
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	g.Log.Debug("delete grant type req: ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+g.Token)
		req.Header.Set("clientId", g.ClientID)
		req.Header.Set("apiKey", g.APIKey)
		_, code := g.Proxy.Do(req, &rtn)
		rtn.Code = code
	}
	return rtn
}
