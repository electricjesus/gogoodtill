package authentication

type AuthLoginParams struct {
	Subdomain string `json:"subdomain"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type AuthLoginResponse struct {
	UserID            string `json:"user_id"`
	UserName          string `json:"user_name"`
	UserLevel         string `json:"user_level"`
	Token             string `json:"token"`
	ClientID          string `json:"client_id"`
	ClientName        string `json:"client_name"`
	ClientSubdomain   string `json:"client_subdomain"`
	CurrentOutletID   string `json:"current_outlet_id"`
	CurrentOutletName string `json:"current_outlet_name"`
	Config            struct {
		ID                      string      `json:"id"`
		Subdomain               string      `json:"subdomain"`
		ClientName              string      `json:"client_name"`
		HasDeliveryModule       int         `json:"has_delivery_module"`
		HasTableModule          int         `json:"has_table_module"`
		EnableIngredients       int         `json:"enable_ingredients"`
		DefaultCurrency         string      `json:"default_currency"`
		LanguageCode            string      `json:"language_code"`
		ClientTimezone          string      `json:"client_timezone"`
		ClientColour            string      `json:"client_colour"`
		ClientLogo              interface{} `json:"client_logo"`
		EnableLoyaltyModule     int         `json:"enable_loyalty_module"`
		EnableStockModule       int         `json:"enable_stock_module"`
		EnablePromoModule       int         `json:"enable_promo_module"`
		EnableMorePrinters      int         `json:"enable_more_printers"`
		EnableAccountPayment    int         `json:"enable_account_payment"`
		ReportCutoffHour        int         `json:"report_cutoff_hour"`
		BoReportCutoffPoint     int         `json:"bo_report_cutoff_point"`
		EnableBetaFeature       int         `json:"enable_beta_feature"`
		ShowEatinTakeaway       int         `json:"show_eatin_takeaway"`
		DefaultCurrencySymbol   string      `json:"default_currency_symbol"`
		DefaultCurrencyHTMLCode string      `json:"default_currency_html_code"`
		IsRetail                int         `json:"is_retail"`
		IsHospitality           int         `json:"is_hospitality"`
		OutletConfig            struct {
			ID                       string `json:"id"`
			OutletName               string `json:"outlet_name"`
			HasOwnBilling            int    `json:"has_own_billing"`
			BillingCompleted         int    `json:"billing_completed"`
			ShowWelcomeTour          int    `json:"show_welcome_tour"`
			EnableOrderCollectModule int    `json:"enable_order_collect_module"`
		} `json:"outlet_config"`
		UserConfig struct {
			ID                 string      `json:"id"`
			Name               string      `json:"name"`
			Email              string      `json:"email"`
			IsMultiOutletAdmin int         `json:"is_multi_outlet_admin"`
			RestrictAccess     int         `json:"restrict_access"`
			OutletTags         interface{} `json:"outlet_tags"`
			PermissionLists    interface{} `json:"permission_lists"`
		} `json:"user_config"`
	} `json:"config"`
}

type AuthRefreshTokenResponse struct {
	Success string `json:"success"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

type AuthLogoutResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

type AuthConfigResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Token struct {
			IssuedAt     string `json:"issued_at"`
			ValidUntil   string `json:"valid_until"`
			RefreshUntil string `json:"refresh_until"`
		} `json:"token"`
		User struct {
			ID                  string      `json:"id"`
			Name                string      `json:"name"`
			Email               interface{} `json:"email"`
			Username            string      `json:"username"`
			OutletID            string      `json:"outlet_id"`
			OutletName          string      `json:"outlet_name"`
			Role                string      `json:"role"`
			StoreAdmin          bool        `json:"store_admin"`
			Status              string      `json:"status"`
			Active              int         `json:"active"`
			IsMultiOutletAdmin  int         `json:"is_multi_outlet_admin"`
			RestrictAccess      int         `json:"restrict_access"`
			OutletTags          string      `json:"outlet_tags"`
			PermissionLists     string      `json:"permission_lists"`
			HasElevatedAccess   int         `json:"has_elevated_access"`
			ElevatedPermissions string      `json:"elevated_permissions"`
		} `json:"user"`
		Client struct {
			ID                      string `json:"id"`
			ClientName              string `json:"client_name"`
			ContactName             string `json:"contact_name"`
			ClientTimezone          string `json:"client_timezone"`
			DefaultCurrency         string `json:"default_currency"`
			DefaultCurrencySymbol   string `json:"default_currency_symbol"`
			DefaultCurrencyHTMLCode string `json:"default_currency_html_code"`
			Config                  struct {
				Loyalty struct {
					Enabled bool `json:"enabled"`
					Config  struct {
						Provider             string `json:"provider"`
						ProviderName         string `json:"provider_name"`
						LoyaltySetupComplete int    `json:"loyalty_setup_complete"`
						PointsPerCurrency    int    `json:"points_per_currency"`
					} `json:"config"`
				} `json:"loyalty"`
			} `json:"config"`
		} `json:"client"`
		Outlet struct {
			ID             string      `json:"id"`
			OutletName     string      `json:"outlet_name"`
			OutletAddress  string      `json:"outlet_address"`
			OutletCity     string      `json:"outlet_city"`
			OutletCounty   interface{} `json:"outlet_county"`
			OutletPostcode string      `json:"outlet_postcode"`
			OutletCountry  string      `json:"outlet_country"`
			StoreTag       string      `json:"store_tag"`
			Status         string      `json:"status"`
			Active         int         `json:"active"`
		} `json:"outlet"`
	} `json:"data"`
	Message string `json:"message"`
}
