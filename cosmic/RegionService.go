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

type ListRegionsParams struct {
	p map[string]interface{}
}

func (p *ListRegionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
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

func (p *ListRegionsParams) SetId(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListRegionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListRegionsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListRegionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListRegionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

// You should always use this function to get a new ListRegionsParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewListRegionsParams() *ListRegionsParams {
	p := &ListRegionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Regions
func (s *RegionService) ListRegions(p *ListRegionsParams) (*ListRegionsResponse, error) {
	var r ListRegionsResponse
	for page := 2; ; page++ {
		var l ListRegionsResponse
		resp, err := s.cs.newRequest("listRegions", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.Regions = append(r.Regions, l.Regions...)

		if r.Count == len(r.Regions) {
			return &r, nil
		}

		p.SetPagesize(len(l.Regions))
		p.SetPage(page)
	}
}

type ListRegionsResponse struct {
	Count   int       `json:"count"`
	Regions []*Region `json:"region"`
}

type Region struct {
	Endpoint string `json:"endpoint,omitempty"`
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
}
