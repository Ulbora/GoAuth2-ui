// Package services ...
package services

import (
	"bytes"
	"fmt"
	"io"

	"net/http"
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

var UID = "bob123456789"
var CLID = "555589999999922222"
var CLIDINT int64 = 555589999999922222

func TestUserService_AddUser(t *testing.T) {
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
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken
	var user User
	user.ClientID = CLIDINT
	user.EmailAddress = "bob@bob.com"
	user.Enabled = true
	user.FirstName = "bob"
	user.LastName = "bob"
	user.Password = "bob"
	user.RoleID = 1
	user.Username = UID
	cc := c.GetNew()
	res := cc.AddUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserPassword(t *testing.T) {
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
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken
	var user UserPW
	user.Username = UID
	user.ClientID = CLIDINT
	user.Password = "bobbby"

	res := c.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserDisable(t *testing.T) {
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
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken
	var user UserDis
	user.Username = UID
	user.ClientID = CLIDINT
	user.Enabled = false

	res := c.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserInfo(t *testing.T) {
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
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken
	var user UserInfo
	user.Username = UID
	user.ClientID = CLIDINT
	user.EmailAddress = "bobbby@bob.com"
	user.FirstName = "bobby"
	user.RoleID = 1
	user.LastName = "williams"

	res := c.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_UpdateUserDisable2(t *testing.T) {
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
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken
	var user UserDis
	user.Username = UID
	user.ClientID = CLIDINT
	user.Enabled = true

	res := c.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_GetUser(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`{"username":"bob123456789", "enabled": true}`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken

	res, code := c.GetUser(UID, CLID)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Username != UID || res.Enabled == false || code != 200 {
		t.Fail()
	}
}

func TestUserService_GetUserList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`[{"username":"bob123456789", "enabled": true}]`))
	p.MockResp = &ress
	c.Proxy = p.GetNewProxy()
	p.MockRespCode = 200
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken

	res, code := c.GetUserList()
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) == 0 || code != 200 {
		t.Fail()
	}
}

func TestUserService_SearchUserList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`[{"username":"bob123456789", "enabled": true}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken

	res, code := c.SearchUserList(CLID)
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) == 0 || code != 200 {
		t.Fail()
	}
}

func TestUserService_DeleteUser(t *testing.T) {
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
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken

	res := c.DeleteUser(UID, CLID)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserService_GetRoleList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = io.NopCloser(bytes.NewBufferString(`[{"role":"test", "id": 3}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.UserHost = "http://localhost:3001"
	c.Token = tempToken

	res, code := c.GetRoleList()
	fmt.Print("role res: ")
	fmt.Println(res)
	if len(*res) == 0 || code != 200 {
		t.Fail()
	}
}

func TestUserService_GetUserType(t *testing.T) {
	var u UserPW

	res := u.GetType()
	fmt.Println("role res: ", res)
	if res != "PW" {
		t.Fail()
	}
}

func TestUserService_GetUserDescType(t *testing.T) {
	var u UserDis

	res := u.GetType()
	fmt.Println("role res: ", res)
	if res != "DIS" {
		t.Fail()
	}
}

func TestUserService_GetUserInfoType(t *testing.T) {
	var u UserInfo

	res := u.GetType()
	fmt.Println("role res: ", res)
	if res != "INFO" {
		t.Fail()
	}
}
