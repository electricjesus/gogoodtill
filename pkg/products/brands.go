package products

import (
	"github.com/electricjesus/gogoodtill/pkg/authentication"
	"github.com/electricjesus/gogoodtill/pkg/util"
)

type BrandsAPI interface {
	List() (*BrandsListResponse, error)
	Get(id string) (*BrandsCreateGetUpdateResponse, error)
	Create(params *BrandsCreateUpdateRequest) (*BrandsCreateGetUpdateResponse, error)
	Update(id string, params *BrandsCreateUpdateRequest) (*BrandsCreateGetUpdateResponse, error)
	Delete(id string) (*BrandsDeleteResponse, error)
}

type BrandsCreateUpdateRequest struct {
	Name   string `json:"brand_name"` // Brand Name
	Active uint   `json:"active"`     // O or 1, should default to 1
}

type BrandResponseData struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active int    `json:"active"`
	Status int    `json:"status"`
}

type BrandsListResponse struct {
	Status  bool                `json:"status"`
	Data    []BrandResponseData `json:"data"`
	Message string              `json:"message"`
}

type BrandsCreateGetUpdateResponse struct {
	Status  bool              `json:"status"`
	Data    BrandResponseData `json:"data"`
	Message string            `json:"message"`
}

type BrandsDeleteRequest struct{}

type BrandsDeleteResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Success int `json:"success"`
	} `json:"data"`
	Message string `json:"message"`
}

type brandsAPI struct {
	ac authentication.AuthenticatedClient
	u  *util.URL
}

func NewBrandsdAPI(ac authentication.AuthenticatedClient) BrandsAPI {
	return &brandsAPI{ac, &util.URL{Base: "https://api.thegoodtill.com/api/brands"}}
}

func (client *brandsAPI) List() (*BrandsListResponse, error) {
	var result BrandsListResponse
	if err := client.ac.Get(client.u.Get(), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *brandsAPI) Get(id string) (*BrandsCreateGetUpdateResponse, error) {
	var result BrandsCreateGetUpdateResponse
	if err := client.ac.Get(client.u.Get(id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *brandsAPI) Create(params *BrandsCreateUpdateRequest) (*BrandsCreateGetUpdateResponse, error) {
	var result BrandsCreateGetUpdateResponse
	if err := client.ac.Post(client.u.Get(), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *brandsAPI) Update(id string, params *BrandsCreateUpdateRequest) (*BrandsCreateGetUpdateResponse, error) {
	var result BrandsCreateGetUpdateResponse
	if err := client.ac.Put(client.u.Get(id), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *brandsAPI) Delete(id string) (*BrandsDeleteResponse, error) {
	var result BrandsDeleteResponse
	if err := client.ac.Del(client.u.Get(id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
