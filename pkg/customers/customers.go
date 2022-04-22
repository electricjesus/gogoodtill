package customers

import (
	"github.com/electricjesus/gogoodtill/pkg/authentication"
	"github.com/electricjesus/gogoodtill/pkg/util"
)

type CustomersAPI interface {
	List() (*CustomerListResponse, error)
	Get(id string) (*CustomerCreateGetUpdateResponse, error)
	Create(params *CustomerCreateUpdateParams) (*CustomerCreateGetUpdateResponse, error)
	Update(id string, params *CustomerCreateUpdateParams) (*CustomerCreateGetUpdateResponse, error)
}

type customersAPI struct {
	ac authentication.AuthenticatedClient
	u  *util.URL
}

func NewCustomersAPIClient(ac authentication.AuthenticatedClient) CustomersAPI {
	return &customersAPI{ac, &util.URL{Base: "https://api.thegoodtill.com/api/customers"}}
}

func (client *customersAPI) List() (*CustomerListResponse, error) {
	var result CustomerListResponse
	if err := client.ac.Get(client.u.Get(), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *customersAPI) Get(id string) (*CustomerCreateGetUpdateResponse, error) {
	var result CustomerCreateGetUpdateResponse
	if err := client.ac.Get(client.u.Get("details", id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *customersAPI) Create(params *CustomerCreateUpdateParams) (*CustomerCreateGetUpdateResponse, error) {
	var result CustomerCreateGetUpdateResponse
	if err := client.ac.Post(client.u.Get(), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *customersAPI) Update(id string, params *CustomerCreateUpdateParams) (*CustomerCreateGetUpdateResponse, error) {
	var result CustomerCreateGetUpdateResponse
	if err := client.ac.Post(client.u.Get(id), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
