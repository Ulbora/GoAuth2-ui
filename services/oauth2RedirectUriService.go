//Package services ...
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
func (r *RedirectURIService) AddRedirectURI(rd *RedirectURI) *RedirectURIResponse {
	var rtn = new(RedirectURIResponse)
	var addURL = r.Host + "/rs/clientRedirectUri/add"
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
			req.Header.Set("Authorization", "Bearer "+r.Token)
			req.Header.Set("clientId", r.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", r.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Redirect URI Add err: ")
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

// GetRedirectURIList get GetRedirectURIList list
func (r *RedirectURIService) GetRedirectURIList(clientID string) *[]RedirectURI {
	var rtn = make([]RedirectURI, 0)
	var gURL = r.Host + "/rs/clientRedirectUri/list/" + clientID
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
		if cErr != nil {
			fmt.Print("Redirect URI list Service read err: ")
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

// DeleteRedirectURI delete DeleteRedirectURI
func (r *RedirectURIService) DeleteRedirectURI(id string) *RedirectURIResponse {
	var rtn = new(RedirectURIResponse)
	var gURL = r.Host + "/rs/clientRedirectUri/delete/" + id
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("clientId", r.ClientID)
		//req.Header.Set("userId", r.UserID)
		//req.Header.Set("hashed", r.Hashed)
		req.Header.Set("apiKey", r.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("redirect uri Service delete err: ")
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
