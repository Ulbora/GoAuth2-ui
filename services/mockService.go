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

//MockOauth2Service MockOauth2Service
type MockOauth2Service struct {
	Token    string
	ClientID string
	//APIKey   string
	//UserID string
	//Hashed string
	Host  string
	Proxy px.Proxy
	Log   *lg.Logger

	//MockCode int

	MockClientResponse *ClientResponse
	MockClient         *Client
	MockClientCode     int
	MockClientList     *[]Client
	MockClientListCode int

	MockAllowedURIResponse *AllowedURIResponse
	MockAllowedURI         *AllowedURI
	MockAllowedURICode     int
	MockAllowedURIList     *[]AllowedURI
	MockAllowedURIListCode int

	MockClientRoleResponse *ClientRoleResponse
	MockClientRoleList     *[]ClientRole
	MockClientRoleListCode int

	MockGrantTypeResponse *GrantTypeResponse
	MockGrantTypeList     *[]GrantType
	MockGrantTypeListCode int

	MockRedirectURIResponse *RedirectURIResponse
	MockRedirectURIList     *[]RedirectURI
	MockRedirectURIListCode int

	MockRoleURIResponse *RoleURIResponse
	MockRoleURIList     *[]RoleURI
	MockRoleURIListCode int

	MockUserResponse *UserResponse
	MockUser         *User
	MockUserCode     int
	MockUserList     *[]User
	MockUserListCode int
	MockRoleList     *[]Role
	MockRoleListCode int
}

//GetNew GetNew
func (s *MockOauth2Service) GetNew() Service {
	var sr Service
	sr = s
	return sr
}

//SetToken SetToken
func (s *MockOauth2Service) SetToken(token string) {
	s.Token = token
}

//AddClient AddClient
func (s *MockOauth2Service) AddClient(client *Client) *ClientResponse {
	return s.MockClientResponse
}

//UpdateClient UpdateClient
func (s *MockOauth2Service) UpdateClient(client *Client) *ClientResponse {
	return s.MockClientResponse
}

//GetClient GetClient
func (s *MockOauth2Service) GetClient(clientID string) (*Client, int) {
	return s.MockClient, s.MockClientCode
}

//GetClientList GetClientList
func (s *MockOauth2Service) GetClientList() (*[]Client, int) {
	return s.MockClientList, s.MockClientListCode
}

//SearchClient SearchClient
func (s *MockOauth2Service) SearchClient(client *Client) (*[]Client, int) {
	return s.MockClientList, s.MockClientListCode
}

//DeleteClient DeleteClient
func (s *MockOauth2Service) DeleteClient(id string) *ClientResponse {
	return s.MockClientResponse
}

//AddAllowedURI AddAllowedURI
func (s *MockOauth2Service) AddAllowedURI(au *AllowedURI) *AllowedURIResponse {
	return s.MockAllowedURIResponse
}

//UpdateAllowedURI UpdateAllowedURI
func (s *MockOauth2Service) UpdateAllowedURI(au *AllowedURI) *AllowedURIResponse {
	return s.MockAllowedURIResponse
}

//GetAllowedURI GetAllowedURI
func (s *MockOauth2Service) GetAllowedURI(id string) (*AllowedURI, int) {
	return s.MockAllowedURI, s.MockAllowedURICode
}

//GetAllowedURIList GetAllowedURIList
func (s *MockOauth2Service) GetAllowedURIList(clientID string) (*[]AllowedURI, int) {
	return s.MockAllowedURIList, s.MockAllowedURIListCode
}

//DeleteAllowedURI DeleteAllowedURI
func (s *MockOauth2Service) DeleteAllowedURI(id string) *AllowedURIResponse {
	return s.MockAllowedURIResponse
}

//AddClientRole AddClientRole
func (s *MockOauth2Service) AddClientRole(cr *ClientRole) *ClientRoleResponse {
	return s.MockClientRoleResponse
}

//GetClientRoleList GetClientRoleList
func (s *MockOauth2Service) GetClientRoleList(clientID string) (*[]ClientRole, int) {
	return s.MockClientRoleList, s.MockClientRoleListCode
}

//DeleteClientRole DeleteClientRole
func (s *MockOauth2Service) DeleteClientRole(id string) *ClientRoleResponse {
	return s.MockClientRoleResponse
}

//AddGrantType AddGrantType
func (s *MockOauth2Service) AddGrantType(rd *GrantType) *GrantTypeResponse {
	return s.MockGrantTypeResponse
}

//GetGrantTypeList GetGrantTypeList
func (s *MockOauth2Service) GetGrantTypeList(clientID string) (*[]GrantType, int) {
	return s.MockGrantTypeList, s.MockGrantTypeListCode
}

//DeleteGrantType DeleteGrantType
func (s *MockOauth2Service) DeleteGrantType(id string) *GrantTypeResponse {
	return s.MockGrantTypeResponse
}

//AddRedirectURI AddRedirectURI
func (s *MockOauth2Service) AddRedirectURI(rd *RedirectURI) *RedirectURIResponse {
	return s.MockRedirectURIResponse
}

//GetRedirectURIList GetRedirectURIList
func (s *MockOauth2Service) GetRedirectURIList(clientID string) (*[]RedirectURI, int) {
	return s.MockRedirectURIList, s.MockRedirectURIListCode
}

//DeleteRedirectURI DeleteRedirectURI
func (s *MockOauth2Service) DeleteRedirectURI(id string) *RedirectURIResponse {
	return s.MockRedirectURIResponse
}

//AddRoleURI AddRoleURI
func (s *MockOauth2Service) AddRoleURI(ru *RoleURI) *RoleURIResponse {
	return s.MockRoleURIResponse
}

//GetRoleURIList GetRoleURIList
func (s *MockOauth2Service) GetRoleURIList(uID string) (*[]RoleURI, int) {
	return s.MockRoleURIList, s.MockRoleURIListCode
}

//DeleteRoleURI DeleteRoleURI
func (s *MockOauth2Service) DeleteRoleURI(ru *RoleURI) *RoleURIResponse {
	return s.MockRoleURIResponse
}

//AddUser AddUser
func (s *MockOauth2Service) AddUser(user *User) *UserResponse {
	return s.MockUserResponse
}

//UpdateUser UpdateUser
func (s *MockOauth2Service) UpdateUser(user UpdateUser) *UserResponse {
	return s.MockUserResponse
}

//GetUser GetUser
func (s *MockOauth2Service) GetUser(username string, clientID string) (*User, int) {
	return s.MockUser, s.MockUserCode
}

//GetUserList GetUserList
func (s *MockOauth2Service) GetUserList() (*[]User, int) {
	return s.MockUserList, s.MockUserListCode
}

//SearchUserList SearchUserList
func (s *MockOauth2Service) SearchUserList(clientID string) (*[]User, int) {
	return s.MockUserList, s.MockUserListCode
}

//DeleteUser DeleteUser
func (s *MockOauth2Service) DeleteUser(username string, clientID string) *UserResponse {
	return s.MockUserResponse
}

//GetRoleList GetRoleList
func (s *MockOauth2Service) GetRoleList() (*[]Role, int) {
	return s.MockRoleList, s.MockRoleListCode
}
