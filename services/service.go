//Package services ...
package services

import (
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

//Service Service
type Service interface {
	AddClient(client *Client) *ClientResponse
	UpdateClient(client *Client) *ClientResponse
	GetClient(clientID string) (*Client, int)
	GetClientList() (*[]Client, int)
	SearchClient(client *Client) (*[]Client, int)
	DeleteClient(id string) *ClientResponse

	AddAllowedURI(au *AllowedURI) *AllowedURIResponse
	UpdateAllowedURI(au *AllowedURI) *AllowedURIResponse
	GetAllowedURI(id string) (*AllowedURI, int)
	GetAllowedURIList(clientID string) (*[]AllowedURI, int)
	DeleteAllowedURI(id string) *AllowedURIResponse

	AddClientRole(cr *ClientRole) *ClientRoleResponse
	GetClientRoleList(clientID string) (*[]ClientRole, int)
	DeleteClientRole(id string) *ClientRoleResponse

	AddGrantType(rd *GrantType) *GrantTypeResponse
	GetGrantTypeList(clientID string) (*[]GrantType, int)
	DeleteGrantType(id string) *GrantTypeResponse

	AddRedirectURI(rd *RedirectURI) *RedirectURIResponse
	GetRedirectURIList(clientID string) (*[]RedirectURI, int)
	DeleteRedirectURI(id string) *RedirectURIResponse

	AddRoleURI(ru *RoleURI) *RoleURIResponse
	GetRoleURIList(uID string) (*[]RoleURI, int)
	DeleteRoleURI(ru *RoleURI) *RoleURIResponse

	AddUser(user *User) *UserResponse
	UpdateUser(user UpdateUser) *UserResponse
	GetUser(username string, clientID string) (*User, int)
	GetUserList() (*[]User, int)
	SearchUserList(clientID string) (*[]User, int)
	DeleteUser(username string, clientID string) *UserResponse
	GetRoleList() (*[]Role, int)

	SetToken(token string)
}

//Oauth2Service Oauth2Service
type Oauth2Service struct {
	Token    string
	ClientID string
	//APIKey   string
	//UserID string
	//Hashed string
	Host  string
	Proxy px.Proxy
	Log   *lg.Logger
}

//GetNew GetNew
func (s *Oauth2Service) GetNew() Service {
	var sr Service
	sr = s
	return sr
}

//SetToken SetToken
func (s *Oauth2Service) SetToken(token string) {
	s.Token = token
}
