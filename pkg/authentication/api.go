package authentication

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/electricjesus/gogoodtill/pkg/consts"
	"github.com/electricjesus/gogoodtill/pkg/util"
)

type client struct {
	loginParams   *AuthLoginParams
	httpClient    *http.Client
	loginDetails  *AuthLoginResponse
	configdetails *AuthConfigResponse
}
type clientOptions func(*client)

func New(options ...clientOptions) API {
	cl := &client{
		httpClient: &http.Client{Timeout: consts.DefaultHTTPTimeout},
	}
	for _, opt := range options {
		opt(cl)
	}
	return cl
}

func WithAuthenticationHTTPClient(httpClient *http.Client) clientOptions {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

func WithLoginParams(params *AuthLoginParams) clientOptions {
	return func(c *client) {
		c.loginParams = params
	}
}

const (
	loginUrl         = "https://api.thegoodtill.com/api/login"
	refreshTokenUrl  = "https://api.thegoodtill.com/api/refresh_token"
	logoutUrl        = "https://api.thegoodtill.com/api/logout"
	configDetailsUrl = "https://api.thegoodtill.com/api/config"
)

func (client *client) Login(params *AuthLoginParams) (*AuthLoginResponse, error) {
	if client.loginParams == nil {
		panic("bug: login parameters not specified")

	}
	var result AuthLoginResponse
	if err := util.JSONHTTPRequestWithToken(
		client.httpClient,
		http.MethodPost,
		loginUrl,
		"",
		nil,
		client.loginParams,
		&result,
	); err != nil {
		return nil, err
	}
	client.loginDetails = &result
	return &result, nil
}

func (client *client) login() error {
	resp, err := client.Login(client.loginParams)
	if err != nil {
		return err
	}
	client.loginDetails = resp
	return nil
}

func (client *client) Token() (string, error) {
	if client.loginDetails != nil {
		return client.loginDetails.Token, nil
	}
	if err := client.login(); err != nil {
		return "", err
	}
	return client.loginDetails.Token, nil
}

func (client *client) Logout(token string) (*AuthLogoutResponse, error) {
	if client.loginDetails == nil {
		return nil, errors.New("auth logout: already logged out")
	}
	var result AuthLogoutResponse
	cl, err := NewAuthenticatedClient(WithToken(token))
	if err != nil {
		return nil, fmt.Errorf("auth logout client err: %w", err)
	}

	if err := cl.Post(logoutUrl, nil, &result); err != nil {
		return nil, fmt.Errorf("auth logout err: %w", err)
	}
	client.loginDetails = nil
	return &result, nil
}

func (client *client) Refresh(token string) ([]AuthRefreshTokenResponse, error) {
	if client.loginDetails == nil {
		return nil, errors.New("auth refresh: please login first")
	}
	var result []AuthRefreshTokenResponse
	cl, err := NewAuthenticatedClient(WithToken(token))
	if err != nil {
		return nil, fmt.Errorf("auth refresh client err: %w", err)
	}

	if err := cl.Get(refreshTokenUrl, result); err != nil {
		return nil, fmt.Errorf("auth refresh err: %w", err)
	}
	return result, nil
}

func (client *client) GetConfig(token, outletID string) (*AuthConfigResponse, error) {
	if client.loginDetails == nil {
		return nil, errors.New("auth details: please login first")
	}

	if client.configdetails != nil {
		return client.configdetails, nil
	}
	var result AuthConfigResponse

	cl, err := NewAuthenticatedClient(
		WithToken(token),
		WithExtraHTTPHeaders(map[string]string{
			"Outlet-Id": outletID,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("auth details err: %w", err)
	}

	if err := cl.Get(configDetailsUrl, &result); err != nil {
		return nil, fmt.Errorf("auth details err: %w", err)
	}
	client.configdetails = &result
	return &result, nil
}
