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
)

type ListHypervisorsParams struct {
	p map[string]interface{}
}

func (p *ListHypervisorsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListHypervisorsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListHypervisorsParams instance,
// as then you are sure you have configured all required params
func (s *HypervisorService) NewListHypervisorsParams() *ListHypervisorsParams {
	p := &ListHypervisorsParams{}
	p.p = make(map[string]interface{})
	return p
}

// List hypervisors
func (s *HypervisorService) ListHypervisors(p *ListHypervisorsParams) (*ListHypervisorsResponse, error) {
	resp, err := s.cs.newRequest("listHypervisors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHypervisorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListHypervisorsResponse struct {
	Count       int           `json:"count"`
	Hypervisors []*Hypervisor `json:"hypervisor"`
}

type Hypervisor struct {
	Name string `json:"name,omitempty"`
}
