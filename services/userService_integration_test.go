// +build integration move to top

//Package services ...
package services

import (
	"fmt"
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

var UIDi = "bob123456789"
var CLIDi = "555589999999922222"
var CLIDINTi int64 = 555589999999922222

func TestUserServicei_AddUser(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken
	var user User
	user.ClientID = CLIDINTi
	user.EmailAddress = "bob@bob.com"
	user.Enabled = true
	user.FirstName = "bob"
	user.LastName = "bob"
	user.Password = "bob"
	user.RoleID = 1
	user.Username = UIDi
	cc := c.GetNew()
	res := cc.AddUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserServicei_UpdateUserPassword(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken
	var user UserPW
	user.Username = UIDi
	user.ClientID = CLIDINTi
	user.Password = "bobbby"

	res := c.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserServicei_UpdateUserDisable(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken
	var user UserDis
	user.Username = UIDi
	user.ClientID = CLIDINTi
	user.Enabled = false

	res := c.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserServicei_UpdateUserInfo(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken
	var user UserInfo
	user.Username = UIDi
	user.ClientID = CLIDINTi
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

func TestUserServicei_UpdateUserDisable2(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken
	var user UserDis
	user.Username = UIDi
	user.ClientID = CLIDINTi
	user.Enabled = true

	res := c.UpdateUser(&user)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserServicei_GetUser(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken

	res := c.GetUser(UIDi, CLIDi)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Username != UIDi || res.Enabled == false {
		t.Fail()
	}
}

func TestUserServicei_GetUserList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken

	res := c.GetUserList()
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) == 0 {
		t.Fail()
	}
}

func TestUserServicei_SearchUserList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken

	res := c.SearchUserList(CLIDi)
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) == 0 {
		t.Fail()
	}
}

func TestUserServicei_DeleteUser(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken

	res := c.DeleteUser(UIDi, CLIDi)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestUserServicei_GetRoleList(t *testing.T) {
	var c Oauth2Service
	var l lg.Logger
	c.Log = &l
	var p px.GoProxy
	c.Proxy = p.GetNewProxy()
	fmt.Println("c.Proxy in test: ", c.Proxy)
	c.ClientID = "10"
	c.Host = "http://localhost:3001"
	c.Token = tempToken

	res := c.GetRoleList()
	fmt.Print("role res: ")
	fmt.Println(res)
	if len(*res) == 0 {
		t.Fail()
	}
}

func TestUserServicei_GetUserType(t *testing.T) {
	var u UserPW

	res := u.GetType()
	fmt.Println("role res: ", res)
	if res != "PW" {
		t.Fail()
	}
}

func TestUserServicei_GetUserDescType(t *testing.T) {
	var u UserDis

	res := u.GetType()
	fmt.Println("role res: ", res)
	if res != "DIS" {
		t.Fail()
	}
}

func TestUserServicei_GetUserInfoType(t *testing.T) {
	var u UserInfo

	res := u.GetType()
	fmt.Println("role res: ", res)
	if res != "INFO" {
		t.Fail()
	}
}
