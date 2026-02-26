// Package services ...
package services

import (
	"bytes"
	"fmt"
	"io"

	"net/http"
	"strconv"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
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

var CID int64

func TestClientService_AddClient(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`{"success":true, "id": 2}`))
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
	s := c.GetNew()
	res := s.AddClient(&cc)
	fmt.Println("res in add: ", res)

	CID = res.ClientID
	if res.Success != true {
		t.Fail()
	}
}

func TestClientService_UpdateClient(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`{"success":true, "id": 2}`))
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
	s := c.GetNew()
	res := s.UpdateClient(&cc)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestClientService_GetClient(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`{"enabled":true, "clientId": 2}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	fmt.Print("CID: ")
	fmt.Println(CID)
	s := c.GetNew()
	res, code := s.GetClient(strconv.FormatInt(CID, 10))
	fmt.Print("res mocked: ")
	fmt.Println(res)
	if res.Enabled != true || code != 200 {
		t.Fail()
	}
}

func TestClientService_SearchClient(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`[{"enabled":true, "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	var cc Client
	cc.Name = "Big"
	s := c.GetNew()
	res, code := s.SearchClient(&cc)
	fmt.Print("searched res: ")
	fmt.Println(res)
	fmt.Println("res len: ", len(*res))
	if res == nil || len(*res) == 0 || code != 200 {
		t.Fail()
	}
}

func TestClientService_GetClientList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`[{"enabled":true, "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	s := c.GetNew()
	res, code := s.GetClientList()
	fmt.Print("res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	fmt.Println("code in getClientList: ", code)
	if res == nil || len(*res) == 0 || code != 200 {
		t.Fail()
	}
}

func TestClientService_DeleteClient(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`{"success":true, "id": 2}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	s := c.GetNew()
	res := s.DeleteClient(strconv.FormatInt(CID, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
