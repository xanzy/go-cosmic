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

type ListOsCategoriesParams struct {
	p map[string]interface{}
}

func (p *ListOsCategoriesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListOsCategoriesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListOsCategoriesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListOsCategoriesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListOsCategoriesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListOsCategoriesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

// You should always use this function to get a new ListOsCategoriesParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListOsCategoriesParams() *ListOsCategoriesParams {
	p := &ListOsCategoriesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListOsCategoriesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListOsCategories(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.OsCategories[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.OsCategories {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryByName(name string, opts ...OptionFunc) (*OsCategory, int, error) {
	id, count, err := s.GetOsCategoryID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetOsCategoryByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryByID(id string, opts ...OptionFunc) (*OsCategory, int, error) {
	p := &ListOsCategoriesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOsCategories(p)
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
		return l.OsCategories[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for OsCategory UUID: %s!", id)
}

// Lists all supported OS categories for this cloud.
func (s *GuestOSService) ListOsCategories(p *ListOsCategoriesParams) (*ListOsCategoriesResponse, error) {
	var r ListOsCategoriesResponse
	for page := 2; ; page++ {
		var l ListOsCategoriesResponse
		resp, err := s.cs.newRequest("listOsCategories", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.OsCategories = append(r.OsCategories, l.OsCategories...)

		if r.Count == len(r.OsCategories) {
			return &r, nil
		}

		p.SetPagesize(len(l.OsCategories))
		p.SetPage(page)
	}
}

type ListOsCategoriesResponse struct {
	Count        int           `json:"count"`
	OsCategories []*OsCategory `json:"oscategory"`
}

type OsCategory struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListOsTypesParams struct {
	p map[string]interface{}
}

func (p *ListOsTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListOsTypesParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
}

func (p *ListOsTypesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListOsTypesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListOsTypesParams) SetOscategoryid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["oscategoryid"] = v
}

func (p *ListOsTypesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListOsTypesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

// You should always use this function to get a new ListOsTypesParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListOsTypesParams() *ListOsTypesParams {
	p := &ListOsTypesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsTypeByID(id string, opts ...OptionFunc) (*OsType, int, error) {
	p := &ListOsTypesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOsTypes(p)
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
		return l.OsTypes[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for OsType UUID: %s!", id)
}

// Lists all supported OS types for this cloud.
func (s *GuestOSService) ListOsTypes(p *ListOsTypesParams) (*ListOsTypesResponse, error) {
	var r ListOsTypesResponse
	for page := 2; ; page++ {
		var l ListOsTypesResponse
		resp, err := s.cs.newRequest("listOsTypes", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.OsTypes = append(r.OsTypes, l.OsTypes...)

		if r.Count == len(r.OsTypes) {
			return &r, nil
		}

		p.SetPagesize(len(l.OsTypes))
		p.SetPage(page)
	}
}

type ListOsTypesResponse struct {
	Count   int       `json:"count"`
	OsTypes []*OsType `json:"ostype"`
}

type OsType struct {
	Description   string `json:"description,omitempty"`
	Id            string `json:"id,omitempty"`
	Isuserdefined string `json:"isuserdefined,omitempty"`
	Oscategoryid  string `json:"oscategoryid,omitempty"`
}
