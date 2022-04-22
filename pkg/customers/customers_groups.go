package customers

import (
	"github.com/electricjesus/gogoodtill/pkg/authentication"
	"github.com/electricjesus/gogoodtill/pkg/util"
)

type CustomerGroupsAPI interface {
	List() (*CustomerGroupListResponse, error)
	Get(id string) (*CustomerGroupCreateGetUpdateResponse, error)
	Create(params *CustomerCreateUpdateRequest) (*CustomerGroupCreateGetUpdateResponse, error)
	Update(id string, params *CustomerCreateUpdateRequest) (*CustomerGroupCreateGetUpdateResponse, error)
	Delete(id string) (*CustomerGroupDeleteResponse, error)
}

type customerGroupsAPI struct {
	ac authentication.AuthenticatedClient
	u  *util.URL
}

func NewCustomerGroupsAPIClient(ac authentication.AuthenticatedClient) CustomerGroupsAPI {
	return &customerGroupsAPI{ac, &util.URL{Base: "https://api.thegoodtill.com/api/customerGroups"}}
}

func (client *customerGroupsAPI) List() (*CustomerGroupListResponse, error) {
	var result CustomerGroupListResponse
	if err := client.ac.Get(client.u.Get(), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *customerGroupsAPI) Get(id string) (*CustomerGroupCreateGetUpdateResponse, error) {
	var result CustomerGroupCreateGetUpdateResponse
	if err := client.ac.Get(client.u.Get(id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *customerGroupsAPI) Create(params *CustomerCreateUpdateRequest) (*CustomerGroupCreateGetUpdateResponse, error) {
	var result CustomerGroupCreateGetUpdateResponse
	if err := client.ac.Post(client.u.Get(), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *customerGroupsAPI) Update(id string, params *CustomerCreateUpdateRequest) (*CustomerGroupCreateGetUpdateResponse, error) {
	var result CustomerGroupCreateGetUpdateResponse
	if err := client.ac.Put(client.u.Get(id), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *customerGroupsAPI) Delete(id string) (*CustomerGroupDeleteResponse, error) {
	var result CustomerGroupDeleteResponse
	if err := client.ac.Del(client.u.Get(id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
