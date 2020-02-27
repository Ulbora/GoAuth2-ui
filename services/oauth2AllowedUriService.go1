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

//AllowedURIService service
type AllowedURIService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

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
func (a *AllowedURIService) AddAllowedURI(au *AllowedURI) *AllowedURIResponse {
	var rtn = new(AllowedURIResponse)
	var addURL = a.Host + "/rs/clientAllowedUriSuper/add"
	aJSON, err := json.Marshal(au)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+a.Token)
			req.Header.Set("clientId", a.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", a.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Client Add err: ")
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

//UpdateAllowedURI update UpdateAllowedURI
func (a *AllowedURIService) UpdateAllowedURI(au *AllowedURI) *AllowedURIResponse {
	var rtn = new(AllowedURIResponse)
	var upURL = a.Host + "/rs/clientAllowedUriSuper/update"

	fmt.Println("in update uri service------------------")
	aJSON, err := json.Marshal(au)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+a.Token)
			req.Header.Set("clientId", a.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", a.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("AllowedURI Service Update err: ")
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
	}
	return rtn
}

// GetAllowedURI get GetAllowedURI
func (a *AllowedURIService) GetAllowedURI(id string) *AllowedURI {
	var rtn = new(AllowedURI)
	var gURL = a.Host + "/rs/clientAllowedUri/get/" + id
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", a.ClientID)
		req.Header.Set("Authorization", "Bearer "+a.Token)
		req.Header.Set("apiKey", a.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("AllowedURI Service read err: ")
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
	return rtn
}

// GetAllowedURIList get GetAllowedURIList list
func (a *AllowedURIService) GetAllowedURIList(clientID string) *[]AllowedURI {
	var rtn = make([]AllowedURI, 0)
	var gURL = a.Host + "/rs/clientAllowedUri/list/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", a.ClientID)
		req.Header.Set("Authorization", "Bearer "+a.Token)
		req.Header.Set("apiKey", a.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("AllowedURI list Service read err: ")
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

// DeleteAllowedURI delete DeleteAllowedURI
func (a *AllowedURIService) DeleteAllowedURI(id string) *AllowedURIResponse {
	var rtn = new(AllowedURIResponse)
	var gURL = a.Host + "/rs/clientAllowedUri/delete/" + id
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+a.Token)
		req.Header.Set("clientId", a.ClientID)
		//req.Header.Set("userId", r.UserID)
		//req.Header.Set("hashed", r.Hashed)
		req.Header.Set("apiKey", a.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("AllowedURI Service delete err: ")
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
