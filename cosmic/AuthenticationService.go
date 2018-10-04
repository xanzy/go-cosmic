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
	"net/url"
	"strconv"
)

type ListDomainLdapLinkParams struct {
	p map[string]interface{}
}

func (p *ListDomainLdapLinkParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	return u
}

func (p *ListDomainLdapLinkParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

// You should always use this function to get a new ListDomainLdapLinkParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewListDomainLdapLinkParams(domainid string) *ListDomainLdapLinkParams {
	p := &ListDomainLdapLinkParams{}
	p.p = make(map[string]interface{})
	p.p["domainid"] = domainid
	return p
}

// list link of domain to group or OU in ldap
func (s *AuthenticationService) ListDomainLdapLink(p *ListDomainLdapLinkParams) (*ListDomainLdapLinkResponse, error) {
	resp, err := s.cs.newRequest("listDomainLdapLink", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainLdapLinkResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListDomainLdapLinkResponse struct {
	Count          int               `json:"count"`
	DomainLdapLink []*DomainLdapLink `json:"domainldaplink"`
}

type DomainLdapLink struct {
	Accountid   string `json:"accountid,omitempty"`
	Accounttype int    `json:"accounttype,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Ldapenabled bool   `json:"ldapenabled,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}

type ListLdapConfigurationsParams struct {
	p map[string]interface{}
}

func (p *ListLdapConfigurationsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (p *ListLdapConfigurationsParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
}

func (p *ListLdapConfigurationsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListLdapConfigurationsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListLdapConfigurationsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListLdapConfigurationsParams) SetPort(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["port"] = v
}

// You should always use this function to get a new ListLdapConfigurationsParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewListLdapConfigurationsParams() *ListLdapConfigurationsParams {
	p := &ListLdapConfigurationsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all LDAP configurations
func (s *AuthenticationService) ListLdapConfigurations(p *ListLdapConfigurationsParams) (*ListLdapConfigurationsResponse, error) {
	var r ListLdapConfigurationsResponse
	for page := 2; ; page++ {
		var l ListLdapConfigurationsResponse
		resp, err := s.cs.newRequest("listLdapConfigurations", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.LdapConfigurations = append(r.LdapConfigurations, l.LdapConfigurations...)

		if r.Count == len(r.LdapConfigurations) {
			return &r, nil
		}

		p.SetPagesize(len(l.LdapConfigurations))
		p.SetPage(page)
	}
}

type ListLdapConfigurationsResponse struct {
	Count              int                  `json:"count"`
	LdapConfigurations []*LdapConfiguration `json:"ldapconfiguration"`
}

type LdapConfiguration struct {
	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
}

type LoginParams struct {
	p map[string]interface{}
}

func (p *LoginParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["domainId"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("domainId", vv)
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *LoginParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *LoginParams) SetDomainId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainId"] = v
}

func (p *LoginParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
}

func (p *LoginParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
}

// You should always use this function to get a new LoginParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLoginParams(password string, username string) *LoginParams {
	p := &LoginParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Logs a user into the CloudStack. A successful login attempt will generate a JSESSIONID cookie value that can be passed in subsequent Query command calls until the "logout" command has been issued or the session has expired.
func (s *AuthenticationService) Login(p *LoginParams) (*LoginResponse, error) {
	resp, err := s.cs.newRequest("login", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LoginResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LoginResponse struct {
	Account    string `json:"account,omitempty"`
	Domainid   string `json:"domainid,omitempty"`
	Domainname string `json:"domainname,omitempty"`
	Firstname  string `json:"firstname,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
	Registered string `json:"registered,omitempty"`
	Sessionkey string `json:"sessionkey,omitempty"`
	Timeout    int    `json:"timeout,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
	Type       string `json:"type,omitempty"`
	Userid     string `json:"userid,omitempty"`
	Username   string `json:"username,omitempty"`
}

type LogoutParams struct {
	p map[string]interface{}
}

func (p *LogoutParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

// You should always use this function to get a new LogoutParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLogoutParams() *LogoutParams {
	p := &LogoutParams{}
	p.p = make(map[string]interface{})
	return p
}

// Logs out the user
func (s *AuthenticationService) Logout(p *LogoutParams) (*LogoutResponse, error) {
	resp, err := s.cs.newRequest("logout", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LogoutResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type LogoutResponse struct {
	Description string `json:"description,omitempty"`
}
