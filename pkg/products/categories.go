package products

import (
	"github.com/electricjesus/gogoodtill/pkg/authentication"
	"github.com/electricjesus/gogoodtill/pkg/util"
)

type CategoriesAPI interface {
	List() (*CategoriesListResponse, error)
	Get(id string) (*CategoriesCreateGetUpdateResponse, error)
	Create(params *CategoriesCreateUpdateRequest) (*CategoriesCreateGetUpdateResponse, error)
	Update(id string, params *CategoriesCreateUpdateRequest) (*CategoriesCreateGetUpdateResponse, error)
	Delete(id string) (*CategoriesDeleteResponse, error)
}

type CategoriesCreateUpdateRequest struct {
	Name             string `json:"category_name"`      // Category Name
	ParentCategoryID string `json:"parent_category_id"` // Parent Category ID. Has to be an UUID and must exist in the system if provided.
	AccountCode      string `json:"account_code"`       // Account Code. To be used when sending data to Accounting platform such as Xero or Exact.
	Active           uint   `json:"active"`             // O or 1, should default to 1
}

type CategoryResponseData struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	ParentCategoryID interface{} `json:"parent_category_id"`
	AccountCode      interface{} `json:"account_code"`
	Active           int         `json:"active"`
	Status           int         `json:"status"`
}

type CategoriesListResponse struct {
	Status  bool                   `json:"status"`
	Data    []CategoryResponseData `json:"data"`
	Message string                 `json:"message"`
}

type CategoriesCreateGetUpdateResponse struct {
	Status  bool                 `json:"status"`
	Data    CategoryResponseData `json:"data"`
	Message string               `json:"message"`
}

type CategoriesDeleteRequest struct{}

type CategoriesDeleteResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Success int `json:"success"`
	} `json:"data"`
	Message string `json:"message"`
}

type categoriesAPI struct {
	ac authentication.AuthenticatedClient
	u  *util.URL
}

func NewCategoriesAPI(ac authentication.AuthenticatedClient) CategoriesAPI {
	return &categoriesAPI{ac, &util.URL{Base: "https://api.thegoodtill.com/api/categories"}}
}

func (client *categoriesAPI) List() (*CategoriesListResponse, error) {
	var result CategoriesListResponse
	if err := client.ac.Get(client.u.Get(), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *categoriesAPI) Get(id string) (*CategoriesCreateGetUpdateResponse, error) {
	var result CategoriesCreateGetUpdateResponse
	if err := client.ac.Get(client.u.Get(id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *categoriesAPI) Create(params *CategoriesCreateUpdateRequest) (*CategoriesCreateGetUpdateResponse, error) {
	var result CategoriesCreateGetUpdateResponse
	if err := client.ac.Post(client.u.Get(), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *categoriesAPI) Update(id string, params *CategoriesCreateUpdateRequest) (*CategoriesCreateGetUpdateResponse, error) {
	var result CategoriesCreateGetUpdateResponse
	if err := client.ac.Put(client.u.Get(id), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *categoriesAPI) Delete(id string) (*CategoriesDeleteResponse, error) {
	var result CategoriesDeleteResponse
	if err := client.ac.Del(client.u.Get(id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
