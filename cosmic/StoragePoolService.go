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

type CancelStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *CancelStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *CancelStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new CancelStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams {
	p := &CancelStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Cancels maintenance for primary storage
func (s *StoragePoolService) CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("cancelStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelStorageMaintenanceResponse
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

type CancelStorageMaintenanceResponse struct {
	JobID                string            `json:"jobid,omitempty"`
	Capacityiops         int64             `json:"capacityiops,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Created              string            `json:"created,omitempty"`
	Disksizeallocated    int64             `json:"disksizeallocated,omitempty"`
	Disksizetotal        int64             `json:"disksizetotal,omitempty"`
	Disksizeused         int64             `json:"disksizeused,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Overprovisionfactor  string            `json:"overprovisionfactor,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	State                string            `json:"state,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
}

type EnableStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *EnableStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new EnableStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams {
	p := &EnableStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Puts storage pool into maintenance state
func (s *StoragePoolService) EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("enableStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableStorageMaintenanceResponse
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

type EnableStorageMaintenanceResponse struct {
	JobID                string            `json:"jobid,omitempty"`
	Capacityiops         int64             `json:"capacityiops,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Created              string            `json:"created,omitempty"`
	Disksizeallocated    int64             `json:"disksizeallocated,omitempty"`
	Disksizetotal        int64             `json:"disksizetotal,omitempty"`
	Disksizeused         int64             `json:"disksizeused,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Overprovisionfactor  string            `json:"overprovisionfactor,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	State                string            `json:"state,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
}

type CreateStoragePoolParams struct {
	p map[string]interface{}
}

func (p *CreateStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["capacityiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacityiops", vv)
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["managed"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("managed", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateStoragePoolParams) SetCapacitybytes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
}

func (p *CreateStoragePoolParams) SetCapacityiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacityiops"] = v
}

func (p *CreateStoragePoolParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *CreateStoragePoolParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateStoragePoolParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *CreateStoragePoolParams) SetManaged(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managed"] = v
}

func (p *CreateStoragePoolParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateStoragePoolParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *CreateStoragePoolParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *CreateStoragePoolParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *CreateStoragePoolParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateStoragePoolParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *CreateStoragePoolParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new CreateStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewCreateStoragePoolParams(name string, url string, zoneid string) *CreateStoragePoolParams {
	p := &CreateStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Creates a storage pool.
func (s *StoragePoolService) CreateStoragePool(p *CreateStoragePoolParams) (*CreateStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("createStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateStoragePoolResponse struct {
	Capacityiops         int64             `json:"capacityiops,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Created              string            `json:"created,omitempty"`
	Disksizeallocated    int64             `json:"disksizeallocated,omitempty"`
	Disksizetotal        int64             `json:"disksizetotal,omitempty"`
	Disksizeused         int64             `json:"disksizeused,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Overprovisionfactor  string            `json:"overprovisionfactor,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	State                string            `json:"state,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
}

type DeleteStoragePoolParams struct {
	p map[string]interface{}
}

func (p *DeleteStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteStoragePoolParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DeleteStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

// You should always use this function to get a new DeleteStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewDeleteStoragePoolParams(id string) *DeleteStoragePoolParams {
	p := &DeleteStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a storage pool.
func (s *StoragePoolService) DeleteStoragePool(p *DeleteStoragePoolParams) (*DeleteStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("deleteStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteStoragePoolResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type UpdateStoragePoolParams struct {
	p map[string]interface{}
}

func (p *UpdateStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["capacityiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacityiops", vv)
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["tags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("tags", vv)
	}
	return u
}

func (p *UpdateStoragePoolParams) SetCapacitybytes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
}

func (p *UpdateStoragePoolParams) SetCapacityiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacityiops"] = v
}

func (p *UpdateStoragePoolParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *UpdateStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateStoragePoolParams) SetTags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

// You should always use this function to get a new UpdateStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewUpdateStoragePoolParams(id string) *UpdateStoragePoolParams {
	p := &UpdateStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a storage pool.
func (s *StoragePoolService) UpdateStoragePool(p *UpdateStoragePoolParams) (*UpdateStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("updateStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateStoragePoolResponse struct {
	Capacityiops         int64             `json:"capacityiops,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Created              string            `json:"created,omitempty"`
	Disksizeallocated    int64             `json:"disksizeallocated,omitempty"`
	Disksizetotal        int64             `json:"disksizetotal,omitempty"`
	Disksizeused         int64             `json:"disksizeused,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Overprovisionfactor  string            `json:"overprovisionfactor,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	State                string            `json:"state,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
}

type ListStoragePoolsParams struct {
	p map[string]interface{}
}

func (p *ListStoragePoolsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
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
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListStoragePoolsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListStoragePoolsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStoragePoolsParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *ListStoragePoolsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStoragePoolsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListStoragePoolsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStoragePoolsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStoragePoolsParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *ListStoragePoolsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListStoragePoolsParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *ListStoragePoolsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

// You should always use this function to get a new ListStoragePoolsParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStoragePoolsParams() *ListStoragePoolsParams {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListStoragePools(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.StoragePools[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.StoragePools {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolByName(name string, opts ...OptionFunc) (*StoragePool, int, error) {
	id, count, err := s.GetStoragePoolID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetStoragePoolByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolByID(id string, opts ...OptionFunc) (*StoragePool, int, error) {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStoragePools(p)
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
		return l.StoragePools[0], l.Count, nil
	}

	return nil, l.Count, fmt.Errorf("There is more then one result for StoragePool UUID: %s!", id)
}

// Lists storage pools.
func (s *StoragePoolService) ListStoragePools(p *ListStoragePoolsParams) (*ListStoragePoolsResponse, error) {
	var r, l ListStoragePoolsResponse
	for page := 2; ; page++ {
		resp, err := s.cs.newRequest("listStoragePools", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.StoragePools = append(r.StoragePools, l.StoragePools...)

		if r.Count == len(r.StoragePools) {
			return &r, nil
		}

		p.SetPagesize(len(l.StoragePools))
		p.SetPage(page)
	}
}

type ListStoragePoolsResponse struct {
	Count        int            `json:"count"`
	StoragePools []*StoragePool `json:"storagepool"`
}

type StoragePool struct {
	Capacityiops         int64             `json:"capacityiops,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Created              string            `json:"created,omitempty"`
	Disksizeallocated    int64             `json:"disksizeallocated,omitempty"`
	Disksizetotal        int64             `json:"disksizetotal,omitempty"`
	Disksizeused         int64             `json:"disksizeused,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Overprovisionfactor  string            `json:"overprovisionfactor,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	State                string            `json:"state,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
}

type FindStoragePoolsForMigrationParams struct {
	p map[string]interface{}
}

func (p *FindStoragePoolsForMigrationParams) toURLValues() url.Values {
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

func (p *FindStoragePoolsForMigrationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *FindStoragePoolsForMigrationParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *FindStoragePoolsForMigrationParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *FindStoragePoolsForMigrationParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

// You should always use this function to get a new FindStoragePoolsForMigrationParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewFindStoragePoolsForMigrationParams(id string) *FindStoragePoolsForMigrationParams {
	p := &FindStoragePoolsForMigrationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Lists storage pools available for migration of a volume.
func (s *StoragePoolService) FindStoragePoolsForMigration(p *FindStoragePoolsForMigrationParams) (*FindStoragePoolsForMigrationResponse, error) {
	resp, err := s.cs.newRequest("findStoragePoolsForMigration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r FindStoragePoolsForMigrationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type FindStoragePoolsForMigrationResponse struct {
	Capacityiops         int64             `json:"capacityiops,omitempty"`
	Clusterid            string            `json:"clusterid,omitempty"`
	Clustername          string            `json:"clustername,omitempty"`
	Created              string            `json:"created,omitempty"`
	Disksizeallocated    int64             `json:"disksizeallocated,omitempty"`
	Disksizetotal        int64             `json:"disksizetotal,omitempty"`
	Disksizeused         int64             `json:"disksizeused,omitempty"`
	Hypervisor           string            `json:"hypervisor,omitempty"`
	Id                   string            `json:"id,omitempty"`
	Ipaddress            string            `json:"ipaddress,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Overprovisionfactor  string            `json:"overprovisionfactor,omitempty"`
	Path                 string            `json:"path,omitempty"`
	Podid                string            `json:"podid,omitempty"`
	Podname              string            `json:"podname,omitempty"`
	Scope                string            `json:"scope,omitempty"`
	State                string            `json:"state,omitempty"`
	Storagecapabilities  map[string]string `json:"storagecapabilities,omitempty"`
	Suitableformigration bool              `json:"suitableformigration,omitempty"`
	Tags                 string            `json:"tags,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Zoneid               string            `json:"zoneid,omitempty"`
	Zonename             string            `json:"zonename,omitempty"`
}

type ListStorageProvidersParams struct {
	p map[string]interface{}
}

func (p *ListStorageProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
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
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ListStorageProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStorageProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStorageProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStorageProvidersParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

// You should always use this function to get a new ListStorageProvidersParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams {
	p := &ListStorageProvidersParams{}
	p.p = make(map[string]interface{})
	p.p["type"] = storagePoolType
	return p
}

// Lists storage providers.
func (s *StoragePoolService) ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error) {
	var r, l ListStorageProvidersResponse
	for page := 2; ; page++ {
		resp, err := s.cs.newRequest("listStorageProviders", p.toURLValues())
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(resp, &l); err != nil {
			return nil, err
		}

		r.Count = l.Count
		r.StorageProviders = append(r.StorageProviders, l.StorageProviders...)

		if r.Count == len(r.StorageProviders) {
			return &r, nil
		}

		p.SetPagesize(len(l.StorageProviders))
		p.SetPage(page)
	}
}

type ListStorageProvidersResponse struct {
	Count            int                `json:"count"`
	StorageProviders []*StorageProvider `json:"storageprovider"`
}

type StorageProvider struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
