//
// Copyright 2018, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cosmic

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type DeleteUserParams struct {
	p map[string]interface{}
}

func (p *DeleteUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDeleteUserParams(id string) *DeleteUserParams {
	p := &DeleteUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a user for an account
func (s *UserService) DeleteUser(p *DeleteUserParams) (*DeleteUserResponse, error) {
	resp, err := s.cs.newRequest("deleteUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteUserResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type DisableUserParams struct {
	p map[string]interface{}
}

func (p *DisableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DisableUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DisableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDisableUserParams(id string) *DisableUserParams {
	p := &DisableUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Disables a user account
func (s *UserService) DisableUser(p *DisableUserParams) (*DisableUserResponse, error) {
	resp, err := s.cs.newRequest("disableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DisableUserResponse struct {
	JobID               string `json:"jobid,omitempty"`
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Created             string `json:"created,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Email               string `json:"email,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	State               string `json:"state,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Username            string `json:"username,omitempty"`
}

type EnableUserParams struct {
	p map[string]interface{}
}

func (p *EnableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new EnableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewEnableUserParams(id string) *EnableUserParams {
	p := &EnableUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Enables a user account
func (s *UserService) EnableUser(p *EnableUserParams) (*EnableUserResponse, error) {
	resp, err := s.cs.newRequest("enableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type EnableUserResponse struct {
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Created             string `json:"created,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Email               string `json:"email,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	State               string `json:"state,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Username            string `json:"username,omitempty"`
}

type LockUserParams struct {
	p map[string]interface{}
}

func (p *LockUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *LockUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new LockUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewLockUserParams(id string) *LockUserParams {
	p := &LockUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Locks a user account
func (s *UserService) LockUser(p *LockUserParams) (*LockUserResponse, error) {
	resp, err := s.cs.newRequest("lockUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LockUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LockUserResponse struct {
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Created             string `json:"created,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Email               string `json:"email,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	State               string `json:"state,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Username            string `json:"username,omitempty"`
}

type UpdateUserParams struct {
	p map[string]interface{}
}

func (p *UpdateUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := p.p["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := p.p["userapikey"]; found {
		u.Set("userapikey", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := p.p["usersecretkey"]; found {
		u.Set("usersecretkey", v.(string))
	}
	return u
}

func (p *UpdateUserParams) SetEmail(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["email"] = v
}

func (p *UpdateUserParams) SetFirstname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["firstname"] = v
}

func (p *UpdateUserParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateUserParams) SetLastname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lastname"] = v
}

func (p *UpdateUserParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *UpdateUserParams) SetTimezone(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timezone"] = v
}

func (p *UpdateUserParams) SetUserapikey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["userapikey"] = v
}

func (p *UpdateUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

func (p *UpdateUserParams) SetUsersecretkey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["usersecretkey"] = v
}

// You should always use this function to get a new UpdateUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewUpdateUserParams(id string) *UpdateUserParams {
	p := &UpdateUserParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a user account
func (s *UserService) UpdateUser(p *UpdateUserParams) (*UpdateUserResponse, error) {
	resp, err := s.cs.newRequest("updateUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateUserResponse struct {
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Created             string `json:"created,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Email               string `json:"email,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	State               string `json:"state,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Username            string `json:"username,omitempty"`
}

type RegisterUserKeysParams struct {
	p map[string]interface{}
}

func (p *RegisterUserKeysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RegisterUserKeysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new RegisterUserKeysParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewRegisterUserKeysParams(id string) *RegisterUserKeysParams {
	p := &RegisterUserKeysParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This command allows a user to register for the developer API, returning a secret key and an API key. This request is made through the integration API port, so it is a privileged command and must be made on behalf of a user. It is up to the implementer just how the username and password are entered, and then how that translates to an integration API request. Both secret key and API key should be returned to the user
func (s *UserService) RegisterUserKeys(p *RegisterUserKeysParams) (*RegisterUserKeysResponse, error) {
	resp, err := s.cs.newRequest("registerUserKeys", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterUserKeysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type RegisterUserKeysResponse struct {
	Apikey    string `json:"apikey,omitempty"`
	Secretkey string `json:"secretkey,omitempty"`
}

type ListUsersParams struct {
	p map[string]interface{}
}

func (p *ListUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *ListUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
}

func (p *ListUsersParams) SetAccounttype(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
}

func (p *ListUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListUsersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListUsersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListUsersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListUsersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
}

func (p *ListUsersParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

// You should always use this function to get a new ListUsersParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewListUsersParams() *ListUsersParams {
	p := &ListUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UserService) GetUserByID(id string, opts ...OptionFunc) (*User, int, error) {
	p := &ListUsersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListUsers(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Users[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for User UUID: %s!", id)
}

// Lists user accounts
func (s *UserService) ListUsers(p *ListUsersParams) (*ListUsersResponse, error) {
	var r ListUsersResponse
	for page := 2; ; page++ {
		var l ListUsersResponse
		resp, err := s.cs.newRequest("listUsers", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.Users = append(r.Users, l.Users...)

		if r.Count == len(r.Users) {
			return &r, nil
		}

		p.SetPagesize(len(l.Users))
		p.SetPage(page)
	}
}

type ListUsersResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"user"`
}

type User struct {
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Created             string `json:"created,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Email               string `json:"email,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	State               string `json:"state,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Username            string `json:"username,omitempty"`
}

type GetVirtualMachineUserDataParams struct {
	p map[string]interface{}
}

func (p *GetVirtualMachineUserDataParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *GetVirtualMachineUserDataParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

// You should always use this function to get a new GetVirtualMachineUserDataParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetVirtualMachineUserDataParams(virtualmachineid string) *GetVirtualMachineUserDataParams {
	p := &GetVirtualMachineUserDataParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Returns user data associated with the VM
func (s *UserService) GetVirtualMachineUserData(p *GetVirtualMachineUserDataParams) (*GetVirtualMachineUserDataResponse, error) {
	resp, err := s.cs.newRequest("getVirtualMachineUserData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVirtualMachineUserDataResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type GetVirtualMachineUserDataResponse struct {
	Userdata         string `json:"userdata,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
}
