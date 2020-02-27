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

//UserService service
type UserService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//User user
type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Enabled      bool   `json:"enabled"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	RoleID       int64  `json:"roleId"`
	ClientID     int64  `json:"clientId"`
}

//UpdateUser interface
type UpdateUser interface {
	GetType() string
}

//UserPW user
type UserPW struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID int64  `json:"clientId"`
}

//GetType type
func (u *UserPW) GetType() string {
	return "PW"
}

//UserDis user
type UserDis struct {
	Username string `json:"username"`
	Enabled  bool   `json:"enabled"`
	ClientID int64  `json:"clientId"`
}

//GetType type
func (u *UserDis) GetType() string {
	return "DIS"
}

//UserInfo user
type UserInfo struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	RoleID       int64  `json:"roleId"`
	ClientID     int64  `json:"clientId"`
}

//GetType type
func (u *UserInfo) GetType() string {
	return "INFO"
}

//Role user role
type Role struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

//UserResponse resp
type UserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//AddUser add
func (u *UserService) AddUser(user *User) *UserResponse {
	var rtn = new(UserResponse)
	var addURL = u.Host + "/rs/user/add"
	aJSON, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+u.Token)
			req.Header.Set("clientId", u.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", u.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("user Add err: ")
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

//UpdateUser update
func (u *UserService) UpdateUser(user UpdateUser) *UserResponse {
	var rtn = new(UserResponse)
	var upURL = u.Host + "/rs/user/update"

	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+u.Token)
			req.Header.Set("clientId", u.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", u.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("User Service Update err: ")
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

// GetUser get
func (u *UserService) GetUser(username string, clientID string) *User {
	var rtn = new(User)
	var gURL = u.Host + "/rs/user/get/" + username + "/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", u.ClientID)
		req.Header.Set("Authorization", "Bearer "+u.Token)
		req.Header.Set("apiKey", u.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("user Service read err: ")
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

// GetUserList get user list
func (u *UserService) GetUserList() *[]User {
	var rtn = make([]User, 0)
	var gURL = u.Host + "/rs/user/list"
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", u.ClientID)
		req.Header.Set("Authorization", "Bearer "+u.Token)
		req.Header.Set("apiKey", u.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("user list Service read err: ")
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

// SearchUserList search by client
func (u *UserService) SearchUserList(clientID string) *[]User {
	var rtn = make([]User, 0)
	var gURL = u.Host + "/rs/user/search/" + clientID
	fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", u.ClientID)
		req.Header.Set("Authorization", "Bearer "+u.Token)
		req.Header.Set("apiKey", u.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		fmt.Print("userList res: ")
		fmt.Println(resp)
		if cErr != nil {
			fmt.Print("user Service read err: ")
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

// DeleteUser delete
func (u *UserService) DeleteUser(username string, clientID string) *UserResponse {
	var rtn = new(UserResponse)
	var gURL = u.Host + "/rs/user/delete/" + username + "/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+u.Token)
		req.Header.Set("clientId", u.ClientID)
		//req.Header.Set("userId", r.UserID)
		//req.Header.Set("hashed", r.Hashed)
		req.Header.Set("apiKey", u.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("user delete err: ")
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

// GetRoleList get role list
func (u *UserService) GetRoleList() *[]Role {
	var rtn = make([]Role, 0)
	var gURL = u.Host + "/rs/role/list"
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	if rErr != nil {
		fmt.Print("request err: ")
		fmt.Println(rErr)
	} else {
		req.Header.Set("clientId", u.ClientID)
		req.Header.Set("Authorization", "Bearer "+u.Token)
		req.Header.Set("apiKey", u.APIKey)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("role list Service read err: ")
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
