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

var CID2 int64
var rdID int64

func TestRedirectURIService_AddClient(t *testing.T) {
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
	res := c.AddClient(&cc)
	fmt.Print("add client res: ")
	fmt.Println(res)
	CID2 = res.ClientID
	if res.Success != true {
		t.Fail()
	}
}

func TestRedirectURIService_AddRedirectURI(t *testing.T) {
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
	uri.URI = "http://yahoooo.com"
	uri.ClientID = CID2
	cc := c.GetNew()
	res := cc.AddRedirectURI(&uri)

	fmt.Print("add uri res: ")
	fmt.Println(res)
	rdID = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestRedirectURIService_GetRedirectURIList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`[{"id":3, "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	c.ClientID = "10"
	c.Host = "http://localhost:3000"
	c.Token = tempToken
	res, code := c.GetRedirectURIList(strconv.FormatInt(CID2, 10))
	fmt.Print("uri res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestRedirectURIService_DeleteRedirectURI(t *testing.T) {
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
	res := c.DeleteRedirectURI(strconv.FormatInt(rdID, 10))
	fmt.Print("res deleted uri: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestRedirectURIService_DeleteClient(t *testing.T) {
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
	res := c.DeleteClient(strconv.FormatInt(CID2, 10))
	fmt.Print("res deleted client: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
