//Package services ...
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// //ClientService service
// type ClientService struct {
// 	Token    string
// 	ClientID string
// 	APIKey   string
// 	UserID   string
// 	Hashed   string
// 	Host     string
// }

//Client client
type Client struct {
	ClientID     int64         `json:"clientId"`
	Secret       string        `json:"secret"`
	Name         string        `json:"name"`
	WebSite      string        `json:"webSite"`
	Email        string        `json:"email"`
	Enabled      bool          `json:"enabled"`
	Paid         bool          `json:"paid"`
	RedirectURIs []RedirectURI `json:"redirectUrls"`
}

//ClientResponse resp
type ClientResponse struct {
	Success  bool  `json:"success"`
	ClientID int64 `json:"id"`
	Code     int   `json:"code"`
}

//AddClient add template
func (c *Oauth2Service) AddClient(client *Client) *ClientResponse {
	var rtn = new(ClientResponse)
	var addURL = c.Host + "/rs/client/add"
	aJSON, err := json.Marshal(client)
	c.Log.Debug("Add Client: ", err)
	if err == nil {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		c.Log.Debug("Add client req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+c.Token)
			req.Header.Set("clientId", c.ClientID)
			//req.Header.Set("apiKey", c.APIKey)
			fmt.Println("c.Proxy: ", c.Proxy)
			_, code := c.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

//UpdateClient update UpdateClient
func (c *Oauth2Service) UpdateClient(client *Client) *ClientResponse {
	var rtn = new(ClientResponse)
	var upURL = c.Host + "/rs/client/update"

	//fmt.Println(content.Text)
	aJSON, err := json.Marshal(client)
	c.Log.Debug("Update client: ", err)
	if err == nil {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		c.Log.Debug("Update client req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+c.Token)
			req.Header.Set("clientId", c.ClientID)
			//req.Header.Set("apiKey", c.APIKey)
			_, code := c.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetClient get GetClient
func (c *Oauth2Service) GetClient(clientID string) (*Client, int) {
	var rtn = new(Client)
	var code int
	var gURL = c.Host + "/rs/client/get/" + clientID
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	c.Log.Debug("Get client req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("Authorization", "Bearer "+c.Token)
		//req.Header.Set("apiKey", c.APIKey)
		_, code = c.Proxy.Do(req, &rtn)
	}
	return rtn, code
}

// GetClientList get client list
func (c *Oauth2Service) GetClientList() (*[]Client, int) {
	var rtn = make([]Client, 0)
	var code int
	var gURL = c.Host + "/rs/client/list"
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("GET", gURL, nil)
	c.Log.Debug("Get client list req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", c.ClientID)
		req.Header.Set("Authorization", "Bearer "+c.Token)
		//req.Header.Set("apiKey", c.APIKey)
		_, code = c.Proxy.Do(req, &rtn)
	}
	return &rtn, code
}

//SearchClient SearchClient
func (c *Oauth2Service) SearchClient(client *Client) (*[]Client, int) {
	var rtn = make([]Client, 0)
	var code int
	var addURL = c.Host + "/rs/client/search"
	aJSON, err := json.Marshal(client)
	c.Log.Debug("search client: ", err)
	if err == nil {
		reqcs, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		c.Log.Debug("search client req: ", rErr)
		if rErr == nil {
			reqcs.Header.Set("Content-Type", "application/json")
			reqcs.Header.Set("Authorization", "Bearer "+c.Token)
			reqcs.Header.Set("clientId", c.ClientID)
			//reqcs.Header.Set("apiKey", c.APIKey)
			_, code = c.Proxy.Do(reqcs, &rtn)
		}
	}
	return &rtn, code
}

// DeleteClient delete DeleteClient
func (c *Oauth2Service) DeleteClient(id string) *ClientResponse {
	var rtn = new(ClientResponse)
	var gURL = c.Host + "/rs/client/delete/" + id
	//fmt.Println(gURL)
	req, rErr := http.NewRequest("DELETE", gURL, nil)
	c.Log.Debug("Delete client: ", rErr)
	if rErr == nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.Token)
		req.Header.Set("clientId", c.ClientID)
		//req.Header.Set("apiKey", c.APIKey)
		_, code := c.Proxy.Do(req, &rtn)
		rtn.Code = code
	}
	return rtn
}
