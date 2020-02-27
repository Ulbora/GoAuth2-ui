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
func (g *GrantTypeService) AddGrantType(rd *GrantType) *GrantTypeResponse {
	var rtn = new(GrantTypeResponse)
	var addURL = g.Host + "/rs/clientGrantType/add"
	aJSON, err := json.Marshal(rd)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+g.Token)
			req.Header.Set("clientId", g.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", g.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("GrantType Add err: ")
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

// GetGrantTypeList get GetGrantTypeList list
func (g *GrantTypeService) GetGrantTypeList(clientID string) *[]GrantType {
	var rtn = make([]GrantType, 0)
	var gURL = g.Host + "/rs/clientGrantType/list/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", g.ClientID)
		req.Header.Set("Authorization", "Bearer "+g.Token)
		req.Header.Set("apiKey", g.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("GrantType list Service read err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
		}
	}
	return &rtn
}

// DeleteGrantType delete DeleteGrantType
func (g *GrantTypeService) DeleteGrantType(id string) *GrantTypeResponse {
	var rtn = new(GrantTypeResponse)
	var gURL = g.Host + "/rs/clientGrantType/delete/" + id
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+g.Token)
		req.Header.Set("clientId", g.ClientID)
		//req.Header.Set("userId", r.UserID)
		//req.Header.Set("hashed", r.Hashed)
		req.Header.Set("apiKey", g.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("GrantType Service delete err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&rtn)
			if error != nil {
				log.Println(error.Error())
			}
			rtn.Code = resp.StatusCode
		}
	}
	return rtn
}
