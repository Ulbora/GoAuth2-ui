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
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	px "github.com/Ulbora/GoProxy"
)

var CID int64

func TestClientService_AddClient(t *testing.T) {
	var c Oauth2Service
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id": 2}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var uri RedirectURI
	uri.URI = "http://googole.com"
	var uris []RedirectURI
	uris = append(uris, uri)
	var cc Client
	cc.Email = "ken@ken.com"
	cc.Enabled = true
	cc.Name = "A Big Company"
	cc.RedirectURIs = uris
	res := c.AddClient(&cc)
	fmt.Println("res in add: ", res)

	CID = res.ClientID
	if res.Success != true {
		t.Fail()
	}
}

func TestClientService_UpdateClient(t *testing.T) {
	var c Oauth2Service
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id": 2}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var cc Client
	cc.Email = "ken@ken1.com"
	cc.Enabled = true
	cc.Name = "A Really Big Company"
	cc.WebSite = "http://www.ulbora.com"
	cc.ClientID = CID
	res := c.UpdateClient(&cc)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestClientService_GetClient(t *testing.T) {
	var c Oauth2Service
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"enabled":true, "clientId": 2}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	fmt.Print("CID: ")
	fmt.Println(CID)
	res := c.GetClient(strconv.FormatInt(CID, 10))
	fmt.Print("res mocked: ")
	fmt.Println(res)
	if res.Enabled != true {
		t.Fail()
	}
}

func TestClientService_SearchClient(t *testing.T) {
	var c Oauth2Service
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"enabled":true, "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var cc Client
	cc.Name = "Big"
	res := c.SearchClient(&cc)
	fmt.Print("searched res: ")
	fmt.Println(res)
	fmt.Println("res len: ", len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestClientService_GetClientList(t *testing.T) {
	var c Oauth2Service
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"enabled":true, "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.GetClientList()
	fmt.Print("res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) == 0 {
		t.Fail()
	}
}

func TestClientService_DeleteClient(t *testing.T) {
	var c Oauth2Service
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id": 2}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res := c.DeleteClient(strconv.FormatInt(CID, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
