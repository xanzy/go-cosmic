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

type CreateServiceOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["bytesreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadrate", vv)
	}
	if v, found := p.p["byteswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriterate", vv)
	}
	if v, found := p.p["cpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpunumber", vv)
	}
	if v, found := p.p["customizediops"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customizediops", vv)
	}
	if v, found := p.p["deploymentplanner"]; found {
		u.Set("deploymentplanner", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["hosttags"]; found {
		u.Set("hosttags", v.(string))
	}
	if v, found := p.p["hypervisorsnapshotreserve"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("hypervisorsnapshotreserve", vv)
	}
	if v, found := p.p["iopsreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadrate", vv)
	}
	if v, found := p.p["iopswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriterate", vv)
	}
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := p.p["isvolatile"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isvolatile", vv)
	}
	if v, found := p.p["limitcpuuse"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("limitcpuuse", vv)
	}
	if v, found := p.p["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := p.p["memory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("memory", vv)
	}
	if v, found := p.p["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("networkrate", vv)
	}
	if v, found := p.p["offerha"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("offerha", vv)
	}
	if v, found := p.p["provisioningtype"]; found {
		u.Set("provisioningtype", v.(string))
	}
	if v, found := p.p["serviceofferingdetails"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("serviceofferingdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("serviceofferingdetails[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["storagetype"]; found {
		u.Set("storagetype", v.(string))
	}
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	return u
}

func (p *CreateServiceOfferingParams) SetBytesreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["bytesreadrate"] = v
}

func (p *CreateServiceOfferingParams) SetByteswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["byteswriterate"] = v
}

func (p *CreateServiceOfferingParams) SetCpunumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cpunumber"] = v
}

func (p *CreateServiceOfferingParams) SetCustomizediops(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["customizediops"] = v
}

func (p *CreateServiceOfferingParams) SetDeploymentplanner(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deploymentplanner"] = v
}

func (p *CreateServiceOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *CreateServiceOfferingParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *CreateServiceOfferingParams) SetHosttags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hosttags"] = v
}

func (p *CreateServiceOfferingParams) SetHypervisorsnapshotreserve(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisorsnapshotreserve"] = v
}

func (p *CreateServiceOfferingParams) SetIopsreadrate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsreadrate"] = v
}

func (p *CreateServiceOfferingParams) SetIopswriterate(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopswriterate"] = v
}

func (p *CreateServiceOfferingParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
}

func (p *CreateServiceOfferingParams) SetIsvolatile(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isvolatile"] = v
}

func (p *CreateServiceOfferingParams) SetLimitcpuuse(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["limitcpuuse"] = v
}

func (p *CreateServiceOfferingParams) SetMaxiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxiops"] = v
}

func (p *CreateServiceOfferingParams) SetMemory(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["memory"] = v
}

func (p *CreateServiceOfferingParams) SetMiniops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["miniops"] = v
}

func (p *CreateServiceOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateServiceOfferingParams) SetNetworkrate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkrate"] = v
}

func (p *CreateServiceOfferingParams) SetOfferha(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["offerha"] = v
}

func (p *CreateServiceOfferingParams) SetProvisioningtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provisioningtype"] = v
}

func (p *CreateServiceOfferingParams) SetServiceofferingdetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingdetails"] = v
}

func (p *CreateServiceOfferingParams) SetStoragetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagetype"] = v
}

func (p *CreateServiceOfferingParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
}

func (p *CreateServiceOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

// You should always use this function to get a new CreateServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewCreateServiceOfferingParams(displaytext string, name string) *CreateServiceOfferingParams {
	p := &CreateServiceOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["name"] = name
	return p
}

// Creates a service offering.
func (s *ServiceOfferingService) CreateServiceOffering(p *CreateServiceOfferingParams) (*CreateServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("createServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateServiceOfferingResponse struct {
	Cpunumber                 int               `json:"cpunumber,omitempty"`
	Created                   string            `json:"created,omitempty"`
	Defaultuse                bool              `json:"defaultuse,omitempty"`
	Deploymentplanner         string            `json:"deploymentplanner,omitempty"`
	DiskBytesReadRate         int64             `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate        int64             `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate          int64             `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate         int64             `json:"diskIopsWriteRate,omitempty"`
	Displaytext               string            `json:"displaytext,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Hosttags                  string            `json:"hosttags,omitempty"`
	Hypervisorsnapshotreserve int               `json:"hypervisorsnapshotreserve,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Iscustomized              bool              `json:"iscustomized,omitempty"`
	Iscustomizediops          bool              `json:"iscustomizediops,omitempty"`
	Issystem                  bool              `json:"issystem,omitempty"`
	Isvolatile                bool              `json:"isvolatile,omitempty"`
	Limitcpuuse               bool              `json:"limitcpuuse,omitempty"`
	Maxiops                   int64             `json:"maxiops,omitempty"`
	Memory                    int               `json:"memory,omitempty"`
	Miniops                   int64             `json:"miniops,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Networkrate               int               `json:"networkrate,omitempty"`
	Offerha                   bool              `json:"offerha,omitempty"`
	Provisioningtype          string            `json:"provisioningtype,omitempty"`
	Serviceofferingdetails    map[string]string `json:"serviceofferingdetails,omitempty"`
	Storagetype               string            `json:"storagetype,omitempty"`
	Systemvmtype              string            `json:"systemvmtype,omitempty"`
	Tags                      string            `json:"tags,omitempty"`
}

type DeleteServiceOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteServiceOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewDeleteServiceOfferingParams(id string) *DeleteServiceOfferingParams {
	p := &DeleteServiceOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a service offering.
func (s *ServiceOfferingService) DeleteServiceOffering(p *DeleteServiceOfferingParams) (*DeleteServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteServiceOfferingResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type UpdateServiceOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	return u
}

func (p *UpdateServiceOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
}

func (p *UpdateServiceOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateServiceOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateServiceOfferingParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
}

// You should always use this function to get a new UpdateServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewUpdateServiceOfferingParams(id string) *UpdateServiceOfferingParams {
	p := &UpdateServiceOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a service offering.
func (s *ServiceOfferingService) UpdateServiceOffering(p *UpdateServiceOfferingParams) (*UpdateServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateServiceOfferingResponse struct {
	Cpunumber                 int               `json:"cpunumber,omitempty"`
	Created                   string            `json:"created,omitempty"`
	Defaultuse                bool              `json:"defaultuse,omitempty"`
	Deploymentplanner         string            `json:"deploymentplanner,omitempty"`
	DiskBytesReadRate         int64             `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate        int64             `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate          int64             `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate         int64             `json:"diskIopsWriteRate,omitempty"`
	Displaytext               string            `json:"displaytext,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Hosttags                  string            `json:"hosttags,omitempty"`
	Hypervisorsnapshotreserve int               `json:"hypervisorsnapshotreserve,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Iscustomized              bool              `json:"iscustomized,omitempty"`
	Iscustomizediops          bool              `json:"iscustomizediops,omitempty"`
	Issystem                  bool              `json:"issystem,omitempty"`
	Isvolatile                bool              `json:"isvolatile,omitempty"`
	Limitcpuuse               bool              `json:"limitcpuuse,omitempty"`
	Maxiops                   int64             `json:"maxiops,omitempty"`
	Memory                    int               `json:"memory,omitempty"`
	Miniops                   int64             `json:"miniops,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Networkrate               int               `json:"networkrate,omitempty"`
	Offerha                   bool              `json:"offerha,omitempty"`
	Provisioningtype          string            `json:"provisioningtype,omitempty"`
	Serviceofferingdetails    map[string]string `json:"serviceofferingdetails,omitempty"`
	Storagetype               string            `json:"storagetype,omitempty"`
	Systemvmtype              string            `json:"systemvmtype,omitempty"`
	Tags                      string            `json:"tags,omitempty"`
}

type ListServiceOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListServiceOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListServiceOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
}

func (p *ListServiceOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListServiceOfferingsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
}

func (p *ListServiceOfferingsParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
}

func (p *ListServiceOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListServiceOfferingsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
}

func (p *ListServiceOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListServiceOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListServiceOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListServiceOfferingsParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
}

func (p *ListServiceOfferingsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
}

// You should always use this function to get a new ListServiceOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewListServiceOfferingsParams() *ListServiceOfferingsParams {
	p := &ListServiceOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListServiceOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListServiceOfferings(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ServiceOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ServiceOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingByName(name string, opts ...OptionFunc) (*ServiceOffering, int, error) {
	id, count, err := s.GetServiceOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetServiceOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingByID(id string, opts ...OptionFunc) (*ServiceOffering, int, error) {
	p := &ListServiceOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListServiceOfferings(p)
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
		return l.ServiceOfferings[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for ServiceOffering UUID: %s!", id)
}

// Lists all available service offerings.
func (s *ServiceOfferingService) ListServiceOfferings(p *ListServiceOfferingsParams) (*ListServiceOfferingsResponse, error) {
	var r, l ListServiceOfferingsResponse
	for page := 2; ; page++ {
		resp, err := s.cs.newRequest("listServiceOfferings", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.ServiceOfferings = append(r.ServiceOfferings, l.ServiceOfferings...)

		if r.Count == len(r.ServiceOfferings) {
			return &r, nil
		}

		p.SetPagesize(len(l.ServiceOfferings))
		p.SetPage(page)
	}
}

type ListServiceOfferingsResponse struct {
	Count            int                `json:"count"`
	ServiceOfferings []*ServiceOffering `json:"serviceoffering"`
}

type ServiceOffering struct {
	Cpunumber                 int               `json:"cpunumber,omitempty"`
	Created                   string            `json:"created,omitempty"`
	Defaultuse                bool              `json:"defaultuse,omitempty"`
	Deploymentplanner         string            `json:"deploymentplanner,omitempty"`
	DiskBytesReadRate         int64             `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate        int64             `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate          int64             `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate         int64             `json:"diskIopsWriteRate,omitempty"`
	Displaytext               string            `json:"displaytext,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Hosttags                  string            `json:"hosttags,omitempty"`
	Hypervisorsnapshotreserve int               `json:"hypervisorsnapshotreserve,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Iscustomized              bool              `json:"iscustomized,omitempty"`
	Iscustomizediops          bool              `json:"iscustomizediops,omitempty"`
	Issystem                  bool              `json:"issystem,omitempty"`
	Isvolatile                bool              `json:"isvolatile,omitempty"`
	Limitcpuuse               bool              `json:"limitcpuuse,omitempty"`
	Maxiops                   int64             `json:"maxiops,omitempty"`
	Memory                    int               `json:"memory,omitempty"`
	Miniops                   int64             `json:"miniops,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Networkrate               int               `json:"networkrate,omitempty"`
	Offerha                   bool              `json:"offerha,omitempty"`
	Provisioningtype          string            `json:"provisioningtype,omitempty"`
	Serviceofferingdetails    map[string]string `json:"serviceofferingdetails,omitempty"`
	Storagetype               string            `json:"storagetype,omitempty"`
	Systemvmtype              string            `json:"systemvmtype,omitempty"`
	Tags                      string            `json:"tags,omitempty"`
}
