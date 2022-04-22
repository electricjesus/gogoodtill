package authentication

type API interface {
	Login(params *AuthLoginParams) (*AuthLoginResponse, error)
	Token() (string, error)
	Refresh(token string) ([]AuthRefreshTokenResponse, error)
	Logout(token string) (*AuthLogoutResponse, error)
	GetConfig(token, outletID string) (*AuthConfigResponse, error)
}

type AuthenticatedClient interface {
	Get(url string, vPtr interface{}) error
	Post(url string, body interface{}, vPtr interface{}) error
	Put(url string, body interface{}, vPtr interface{}) error
	Del(url string, vPtr interface{}) error
}

type TokenSource interface {
	Token() (string, error)
}
