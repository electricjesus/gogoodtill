package products

import (
	"github.com/electricjesus/gogoodtill/pkg/authentication"
	"github.com/electricjesus/gogoodtill/pkg/util"
)

type ProductsAPI interface {
	// GET https://api.thegoodtill.com/api/products
	List() (*ProductListResponse, error)
	// GET https://api.thegoodtill.com/api/products/:id
	Get(id string) (*ProductDetailsReponse, error)
	// POST https://api.thegoodtill.com/api/products
	Create(params *ProductCreateUpdateRequest) (*ProductCreateUpdateResponse, error)
	// PUT https://api.thegoodtill.com/api/products/:id
	Update(id string, params *ProductCreateUpdateRequest) (*ProductCreateUpdateResponse, error)
	// DELETE https://api.thegoodtill.com/api/products/delete/:id
	Delete(id string) (*ProductDeleteResponse, error)
}

type productsAPI struct {
	ac authentication.AuthenticatedClient
	u  *util.URL
}

func NewProductsAPI(ac authentication.AuthenticatedClient) ProductsAPI {
	return &productsAPI{ac, &util.URL{Base: "https://api.thegoodtill.com/api/products"}}
}

func (client *productsAPI) List() (*ProductListResponse, error) {
	var result ProductListResponse
	if err := client.ac.Get(client.u.Get(), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *productsAPI) Get(id string) (*ProductDetailsReponse, error) {
	var result ProductDetailsReponse
	if err := client.ac.Get(client.u.Get(id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *productsAPI) Create(params *ProductCreateUpdateRequest) (*ProductCreateUpdateResponse, error) {
	var result ProductCreateUpdateResponse
	if err := client.ac.Post(client.u.Get(), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *productsAPI) Update(id string, params *ProductCreateUpdateRequest) (*ProductCreateUpdateResponse, error) {
	var result ProductCreateUpdateResponse
	if err := client.ac.Put(client.u.Get(id), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
func (client *productsAPI) Delete(id string) (*ProductDeleteResponse, error) {
	var result ProductDeleteResponse
	if err := client.ac.Del(client.u.Get("delete", id), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
