package authentication

import (
	"net/http"

	"github.com/electricjesus/gogoodtill/pkg/consts"
	"github.com/electricjesus/gogoodtill/pkg/util"
)

type authenticatedClient struct {
	token       string
	httpClient  *http.Client
	headers     map[string]string
	tokenSource TokenSource
}

type ClientOption func(*authenticatedClient)

func NewAuthenticatedClient(opts ...ClientOption) (AuthenticatedClient, error) {
	o := &authenticatedClient{}
	if len(opts) == 0 {
		panic("must provide one of [WithToken, WithAuthParams]")
	}

	for _, opt := range opts {
		opt(o)
	}
	o.httpClient = &http.Client{Timeout: consts.DefaultHTTPTimeout}
	return o, nil
}

func WithToken(token string) ClientOption {
	return func(ac *authenticatedClient) {
		ac.token = token
	}
}

func WithTokenSource(tokenSource TokenSource) ClientOption {
	return func(ac *authenticatedClient) {
		ac.tokenSource = tokenSource
	}
}

func WithExtraHTTPHeaders(headers map[string]string) ClientOption {
	return func(ac *authenticatedClient) {
		ac.headers = headers
	}
}

// Get - GET request to API url and serialize result to vPtr
func (client *authenticatedClient) Get(url string, vPtr interface{}) error {
	token, err := client.tokenSource.Token()
	if err != nil {
		return err
	}
	return util.JSONHTTPRequestWithToken(
		client.httpClient,
		http.MethodGet,
		url,
		token,
		client.headers,
		nil,
		vPtr,
	)
}

// Post - POST request to API url and serialize result to vPtr
func (client *authenticatedClient) Post(url string, requestBody interface{}, vPtr interface{}) error {
	token, err := client.tokenSource.Token()
	if err != nil {
		return err
	}

	return util.JSONHTTPRequestWithToken(
		client.httpClient,
		http.MethodPost,
		url,
		token,
		client.headers,
		requestBody,
		vPtr,
	)
}

// Post - POST request to API url and serialize result to vPtr
func (client *authenticatedClient) Put(url string, requestBody interface{}, vPtr interface{}) error {
	token, err := client.tokenSource.Token()
	if err != nil {
		return err
	}

	return util.JSONHTTPRequestWithToken(
		client.httpClient,
		http.MethodPut,
		url,
		token,
		client.headers,
		requestBody,
		vPtr,
	)
}

// Del - DELETE request to API url and serialize result to vPtr
func (client *authenticatedClient) Del(url string, vPtr interface{}) error {
	token, err := client.tokenSource.Token()
	if err != nil {
		return err
	}
	return util.JSONHTTPRequestWithToken(
		client.httpClient,
		http.MethodDelete,
		url,
		token,
		client.headers,
		nil,
		vPtr,
	)
}
