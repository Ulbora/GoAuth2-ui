//Package services ...
package services

import (
	"testing"
)

func TestMockOauth2Service_SetToken(t *testing.T) {
	var os MockOauth2Service
	s := os.GetNew()
	s.SetToken("12345")
	if os.Token != "12345" {
		t.Fail()
	}
}

func TestMockOauth2Service_AddClient(t *testing.T) {
	var os MockOauth2Service
	var cr ClientResponse
	cr.Success = true
	cr.Code = 200
	os.MockClientResponse = &cr
	s := os.GetNew()
	var c Client
	c.ClientID = 55
	res := s.AddClient(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_UpdateClient(t *testing.T) {
	var os MockOauth2Service
	var cr ClientResponse
	cr.Success = true
	cr.Code = 200
	os.MockClientResponse = &cr
	s := os.GetNew()
	var c Client
	c.ClientID = 55
	res := s.UpdateClient(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetClient(t *testing.T) {
	var os MockOauth2Service
	var c Client
	c.ClientID = 55
	os.MockClient = &c
	os.MockClientCode = 200
	s := os.GetNew()
	res, code := s.GetClient("55")
	if res.ClientID != 55 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_GetClientList(t *testing.T) {
	var os MockOauth2Service
	var c Client
	c.ClientID = 55
	var cl []Client
	cl = append(cl, c)
	os.MockClientList = &cl
	os.MockClientListCode = 200
	s := os.GetNew()
	res, code := s.GetClientList()
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_SearchClientList(t *testing.T) {
	var os MockOauth2Service
	var c Client
	c.ClientID = 55
	var cl []Client
	cl = append(cl, c)
	os.MockClientList = &cl
	os.MockClientListCode = 200
	s := os.GetNew()
	var cc Client
	res, code := s.SearchClient(&cc)
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_DeleteClient(t *testing.T) {
	var os MockOauth2Service
	var cr ClientResponse
	cr.Success = true
	cr.Code = 200
	os.MockClientResponse = &cr
	s := os.GetNew()
	var c Client
	c.ClientID = 55
	res := s.DeleteClient("123")
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_AddAllowedURI(t *testing.T) {
	var os MockOauth2Service
	var cr AllowedURIResponse
	cr.Success = true
	cr.Code = 200
	os.MockAllowedURIResponse = &cr
	s := os.GetNew()
	var c AllowedURI
	c.ClientID = 55
	res := s.AddAllowedURI(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_UpdateAllowedURI(t *testing.T) {
	var os MockOauth2Service
	var cr AllowedURIResponse
	cr.Success = true
	cr.Code = 200
	os.MockAllowedURIResponse = &cr
	s := os.GetNew()
	var c AllowedURI
	c.ClientID = 55
	res := s.UpdateAllowedURI(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetAllowedURI(t *testing.T) {
	var os MockOauth2Service
	var c AllowedURI
	c.ClientID = 55
	os.MockAllowedURI = &c
	os.MockAllowedURICode = 200
	s := os.GetNew()
	res, code := s.GetAllowedURI("55")
	if res.ClientID != 55 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_GetAllowedURIList(t *testing.T) {
	var os MockOauth2Service
	var c AllowedURI
	c.ClientID = 55
	var cl []AllowedURI
	cl = append(cl, c)
	os.MockAllowedURIList = &cl
	os.MockAllowedURIListCode = 200
	s := os.GetNew()
	res, code := s.GetAllowedURIList("55")
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_DeleteAllowedURI(t *testing.T) {
	var os MockOauth2Service
	var cr AllowedURIResponse
	cr.Success = true
	cr.Code = 200
	os.MockAllowedURIResponse = &cr
	s := os.GetNew()

	res := s.DeleteAllowedURI("123")
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_AddClientRole(t *testing.T) {
	var os MockOauth2Service
	var cr ClientRoleResponse
	cr.Success = true
	cr.Code = 200
	os.MockClientRoleResponse = &cr
	s := os.GetNew()
	var c ClientRole
	c.ClientID = 55
	res := s.AddClientRole(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetClientRoleList(t *testing.T) {
	var os MockOauth2Service
	var c ClientRole
	c.ClientID = 55
	var cl []ClientRole
	cl = append(cl, c)
	os.MockClientRoleList = &cl
	os.MockClientRoleListCode = 200
	s := os.GetNew()
	res, code := s.GetClientRoleList("55")
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_DeleteClientRole(t *testing.T) {
	var os MockOauth2Service
	var cr ClientRoleResponse
	cr.Success = true
	cr.Code = 200
	os.MockClientRoleResponse = &cr
	s := os.GetNew()

	res := s.DeleteClientRole("123")
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_AddGrantType(t *testing.T) {
	var os MockOauth2Service
	var cr GrantTypeResponse
	cr.Success = true
	cr.Code = 200
	os.MockGrantTypeResponse = &cr
	s := os.GetNew()
	var c GrantType
	c.ClientID = 55
	res := s.AddGrantType(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetGrantTypeList(t *testing.T) {
	var os MockOauth2Service
	var c GrantType
	c.ClientID = 55
	var cl []GrantType
	cl = append(cl, c)
	os.MockGrantTypeList = &cl
	os.MockGrantTypeListCode = 200
	s := os.GetNew()
	res, code := s.GetGrantTypeList("55")
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_DeleteGrantType(t *testing.T) {
	var os MockOauth2Service
	var cr GrantTypeResponse
	cr.Success = true
	cr.Code = 200
	os.MockGrantTypeResponse = &cr
	s := os.GetNew()

	res := s.DeleteGrantType("123")
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_AddRedirectURI(t *testing.T) {
	var os MockOauth2Service
	var cr RedirectURIResponse
	cr.Success = true
	cr.Code = 200
	os.MockRedirectURIResponse = &cr
	s := os.GetNew()
	var c RedirectURI
	c.ClientID = 55
	res := s.AddRedirectURI(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetRedirectURIList(t *testing.T) {
	var os MockOauth2Service
	var c RedirectURI
	c.ClientID = 55
	var cl []RedirectURI
	cl = append(cl, c)
	os.MockRedirectURIList = &cl
	os.MockRedirectURIListCode = 200
	s := os.GetNew()
	res, code := s.GetRedirectURIList("55")
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_DeleteRedirectURI(t *testing.T) {
	var os MockOauth2Service
	var cr RedirectURIResponse
	cr.Success = true
	cr.Code = 200
	os.MockRedirectURIResponse = &cr
	s := os.GetNew()

	res := s.DeleteRedirectURI("123")
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_AddRoleURI(t *testing.T) {
	var os MockOauth2Service
	var cr RoleURIResponse
	cr.Success = true
	cr.Code = 200
	os.MockRoleURIResponse = &cr
	s := os.GetNew()
	var c RoleURI
	c.ClientRoleID = 55
	res := s.AddRoleURI(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetRoleURIList(t *testing.T) {
	var os MockOauth2Service
	var c RoleURI
	c.ClientRoleID = 55
	var cl []RoleURI
	cl = append(cl, c)
	os.MockRoleURIList = &cl
	os.MockRoleURIListCode = 200
	s := os.GetNew()
	res, code := s.GetRoleURIList("55")
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_DeleteRoleURI(t *testing.T) {
	var os MockOauth2Service
	var cr RoleURIResponse
	cr.Success = true
	cr.Code = 200
	os.MockRoleURIResponse = &cr
	s := os.GetNew()
	var ru RoleURI

	res := s.DeleteRoleURI(&ru)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_AddUser(t *testing.T) {
	var os MockOauth2Service
	var cr UserResponse
	cr.Success = true
	cr.Code = 200
	os.MockUserResponse = &cr
	s := os.GetNew()
	var c User
	c.ClientID = 55
	res := s.AddUser(&c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_UpdateUser(t *testing.T) {
	var os MockOauth2Service
	var cr UserResponse
	cr.Success = true
	cr.Code = 200
	os.MockUserResponse = &cr
	s := os.GetNew()
	var c UpdateUser
	res := s.UpdateUser(c)
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetUser(t *testing.T) {
	var os MockOauth2Service
	var c User
	c.ClientID = 55
	os.MockUser = &c
	os.MockUserCode = 200
	s := os.GetNew()
	res, code := s.GetUser("tester", "55")
	if res.ClientID != 55 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_GetUserList(t *testing.T) {
	var os MockOauth2Service
	var c User
	c.ClientID = 55
	var cl []User
	cl = append(cl, c)
	os.MockUserList = &cl
	os.MockUserListCode = 200
	s := os.GetNew()
	res, code := s.GetUserList()
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_SearchUserList(t *testing.T) {
	var os MockOauth2Service
	var c User
	c.ClientID = 55
	var cl []User
	cl = append(cl, c)
	os.MockUserList = &cl
	os.MockUserListCode = 200
	s := os.GetNew()
	res, code := s.SearchUserList("55")
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}

func TestMockOauth2Service_DeleteUser(t *testing.T) {
	var os MockOauth2Service
	var cr UserResponse
	cr.Success = true
	cr.Code = 200
	os.MockUserResponse = &cr
	s := os.GetNew()

	res := s.DeleteUser("tester", "123")
	if !res.Success {
		t.Fail()
	}
}

func TestMockOauth2Service_GetRoleList(t *testing.T) {
	var os MockOauth2Service
	var c Role
	c.ID = 55
	var cl []Role
	cl = append(cl, c)
	os.MockRoleList = &cl
	os.MockRoleListCode = 200
	s := os.GetNew()
	res, code := s.GetRoleList()
	if len(*res) != 1 || code != 200 {
		t.Fail()
	}
}
