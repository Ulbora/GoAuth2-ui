/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//RoleURIService service
type RoleURIService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

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
func (r *RoleURIService) AddRoleURI(ru *RoleURI) *RoleURIResponse {
	var rtn = new(RoleURIResponse)
	var addURL = r.Host + "/rs/clientRoleUri/add"
	aJSON, err := json.Marshal(ru)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", r.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("RoleURI Add err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				//fmt.Print("resp: ")
				//fmt.Println(resp)
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}

// GetRoleURIList get GetRoleURIList list
func (r *RoleURIService) GetRoleURIList(uID string) *[]RoleURI {
	var rtn = make([]RoleURI, 0)
	var gURL = r.Host + "/rs/clientRoleUri/list/" + uID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", r.ClientID)
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("apiKey", r.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		//fmt.Println(gURL)
		if cErr != nil {
			fmt.Print("RoleURI list Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			//fmt.Println(rtn)
			if error != nil {
				log.Println(error.Error())
			}
		}
	}
	return &rtn
}

//DeleteRoleURI add
func (r *RoleURIService) DeleteRoleURI(ru *RoleURI) *RoleURIResponse {
	var rtn = new(RoleURIResponse)
	var uURL = r.Host + "/rs/clientRoleUri/delete"
	aJSON, err := json.Marshal(ru)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", uURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", r.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("RoleURI delete err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				//fmt.Print("resp: ")
				//fmt.Println(resp)
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}
