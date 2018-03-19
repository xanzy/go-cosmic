//
// Copyright 2016, Sander van Harmelen
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

type AddImageStoreParams struct {
	p map[string]interface{}
}

func (p *AddImageStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddImageStoreParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *AddImageStoreParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddImageStoreParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *AddImageStoreParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddImageStoreParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new AddImageStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewAddImageStoreParams(provider string) *AddImageStoreParams {
	p := &AddImageStoreParams{}
	p.p = make(map[string]interface{})
	p.p["provider"] = provider
	return p
}

// Adds backup image store.
func (s *ImageStoreService) AddImageStore(p *AddImageStoreParams) (*AddImageStoreResponse, error) {
	resp, err := s.cs.newRequest("addImageStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddImageStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type AddImageStoreResponse struct {
	Details      []string `json:"details,omitempty"`
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	Providername string   `json:"providername,omitempty"`
	Scope        string   `json:"scope,omitempty"`
	Url          string   `json:"url,omitempty"`
	Zoneid       string   `json:"zoneid,omitempty"`
	Zonename     string   `json:"zonename,omitempty"`
}

type DeleteImageStoreParams struct {
	p map[string]interface{}
}

func (p *DeleteImageStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteImageStoreParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteImageStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewDeleteImageStoreParams(id string) *DeleteImageStoreParams {
	p := &DeleteImageStoreParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an image store or Secondary Storage.
func (s *ImageStoreService) DeleteImageStore(p *DeleteImageStoreParams) (*DeleteImageStoreResponse, error) {
	resp, err := s.cs.newRequest("deleteImageStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteImageStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteImageStoreResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type ListImageStoresParams struct {
	p map[string]interface{}
}

func (p *ListImageStoresParams) toURLValues() url.Values {
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
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListImageStoresParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListImageStoresParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListImageStoresParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListImageStoresParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListImageStoresParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListImageStoresParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *ListImageStoresParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListImageStoresParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListImageStoresParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewListImageStoresParams() *ListImageStoresParams {
	p := &ListImageStoresParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListImageStoresParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListImageStores(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ImageStores[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ImageStores {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreByName(name string, opts ...OptionFunc) (*ImageStore, int, error) {
	id, count, err := s.GetImageStoreID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetImageStoreByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreByID(id string, opts ...OptionFunc) (*ImageStore, int, error) {
	p := &ListImageStoresParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListImageStores(p)
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
		return l.ImageStores[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for ImageStore UUID: %s!", id)
}

// Lists image stores.
func (s *ImageStoreService) ListImageStores(p *ListImageStoresParams) (*ListImageStoresResponse, error) {
	var r, l ListImageStoresResponse
	for page := 2; ; page++ {
		resp, err := s.cs.newRequest("listImageStores", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.ImageStores = append(r.ImageStores, l.ImageStores...)

		if r.Count == len(r.ImageStores) {
			return &r, nil
		}

		p.SetPagesize(len(l.ImageStores))
		p.SetPage(page)
	}
}

type ListImageStoresResponse struct {
	Count       int           `json:"count"`
	ImageStores []*ImageStore `json:"imagestore"`
}

type ImageStore struct {
	Details      []string `json:"details,omitempty"`
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	Providername string   `json:"providername,omitempty"`
	Scope        string   `json:"scope,omitempty"`
	Url          string   `json:"url,omitempty"`
	Zoneid       string   `json:"zoneid,omitempty"`
	Zonename     string   `json:"zonename,omitempty"`
}

type CreateSecondaryStagingStoreParams struct {
	p map[string]interface{}
}

func (p *CreateSecondaryStagingStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateSecondaryStagingStoreParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateSecondaryStagingStoreParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *CreateSecondaryStagingStoreParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *CreateSecondaryStagingStoreParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *CreateSecondaryStagingStoreParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new CreateSecondaryStagingStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewCreateSecondaryStagingStoreParams(url string) *CreateSecondaryStagingStoreParams {
	p := &CreateSecondaryStagingStoreParams{}
	p.p = make(map[string]interface{})
	p.p["url"] = url
	return p
}

// create secondary staging store.
func (s *ImageStoreService) CreateSecondaryStagingStore(p *CreateSecondaryStagingStoreParams) (*CreateSecondaryStagingStoreResponse, error) {
	resp, err := s.cs.newRequest("createSecondaryStagingStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSecondaryStagingStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateSecondaryStagingStoreResponse struct {
	Details      []string `json:"details,omitempty"`
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	Providername string   `json:"providername,omitempty"`
	Scope        string   `json:"scope,omitempty"`
	Url          string   `json:"url,omitempty"`
	Zoneid       string   `json:"zoneid,omitempty"`
	Zonename     string   `json:"zonename,omitempty"`
}

type DeleteSecondaryStagingStoreParams struct {
	p map[string]interface{}
}

func (p *DeleteSecondaryStagingStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteSecondaryStagingStoreParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteSecondaryStagingStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewDeleteSecondaryStagingStoreParams(id string) *DeleteSecondaryStagingStoreParams {
	p := &DeleteSecondaryStagingStoreParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a secondary staging store .
func (s *ImageStoreService) DeleteSecondaryStagingStore(p *DeleteSecondaryStagingStoreParams) (*DeleteSecondaryStagingStoreResponse, error) {
	resp, err := s.cs.newRequest("deleteSecondaryStagingStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSecondaryStagingStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteSecondaryStagingStoreResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type ListSecondaryStagingStoresParams struct {
	p map[string]interface{}
}

func (p *ListSecondaryStagingStoresParams) toURLValues() url.Values {
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
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSecondaryStagingStoresParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListSecondaryStagingStoresParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListSecondaryStagingStoresParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListSecondaryStagingStoresParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListSecondaryStagingStoresParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListSecondaryStagingStoresParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *ListSecondaryStagingStoresParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListSecondaryStagingStoresParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListSecondaryStagingStoresParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewListSecondaryStagingStoresParams() *ListSecondaryStagingStoresParams {
	p := &ListSecondaryStagingStoresParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListSecondaryStagingStoresParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSecondaryStagingStores(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SecondaryStagingStores[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SecondaryStagingStores {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreByName(name string, opts ...OptionFunc) (*SecondaryStagingStore, int, error) {
	id, count, err := s.GetSecondaryStagingStoreID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSecondaryStagingStoreByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreByID(id string, opts ...OptionFunc) (*SecondaryStagingStore, int, error) {
	p := &ListSecondaryStagingStoresParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSecondaryStagingStores(p)
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
		return l.SecondaryStagingStores[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for SecondaryStagingStore UUID: %s!", id)
}

// Lists secondary staging stores.
func (s *ImageStoreService) ListSecondaryStagingStores(p *ListSecondaryStagingStoresParams) (*ListSecondaryStagingStoresResponse, error) {
	var r, l ListSecondaryStagingStoresResponse
	for page := 2; ; page++ {
		resp, err := s.cs.newRequest("listSecondaryStagingStores", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.SecondaryStagingStores = append(r.SecondaryStagingStores, l.SecondaryStagingStores...)

		if r.Count == len(r.SecondaryStagingStores) {
			return &r, nil
		}

		p.SetPagesize(len(l.SecondaryStagingStores))
		p.SetPage(page)
	}
}

type ListSecondaryStagingStoresResponse struct {
	Count                  int                      `json:"count"`
	SecondaryStagingStores []*SecondaryStagingStore `json:"secondarystagingstore"`
}

type SecondaryStagingStore struct {
	Details      []string `json:"details,omitempty"`
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	Providername string   `json:"providername,omitempty"`
	Scope        string   `json:"scope,omitempty"`
	Url          string   `json:"url,omitempty"`
	Zoneid       string   `json:"zoneid,omitempty"`
	Zonename     string   `json:"zonename,omitempty"`
}
