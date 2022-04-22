package customers

type CustomerCreateUpdateParams struct {
	CustomerGroupID  string `json:"customer_group_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	Address          string `json:"address,omitempty"`
	City             string `json:"city,omitempty"`
	County           string `json:"county,omitempty"`
	PostCode         string `json:"postcode,omitempty"`
	Country          string `json:"country,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
	Active           uint   `json:"active,omitempty"`
	ExtraNotes       string `json:"extra_nodes,omitempty"`
	EmailOptIn       uint   `json:"opt_in_email,omitempty"`
	AccountCode      string `json:"account_code,omitempty"`
	MembershipNumber string `json:"membership_no,omitempty"`
	Source           string `json:"source,omitempty"`
}

type Customer struct {
	ID                string      `json:"id"`
	CustomerGroup     string      `json:"customer_group"`
	CustomerGroupID   interface{} `json:"customer_group_id"`
	TotalOrders       int         `json:"total_orders"`
	LoyaltyPoints     int         `json:"loyalty_points"`
	Name              string      `json:"name"`
	Email             string      `json:"email"`
	Address           string      `json:"address"`
	City              string      `json:"city"`
	County            string      `json:"county"`
	Postcode          string      `json:"postcode"`
	Country           string      `json:"country"`
	Active            int         `json:"active"`
	Status            string      `json:"status"`
	Mobile            string      `json:"mobile"`
	Phone             interface{} `json:"phone"`
	Website           string      `json:"website"`
	ExtraNotes        string      `json:"extra_notes"`
	AccountCode       string      `json:"account_code"`
	OptInEmail        int         `json:"opt_in_email"`
	OptIn             string      `json:"opt_in"`
	IsAccountCustomer int         `json:"is_account_customer"`
	CreditLimit       string      `json:"credit_limit"`
	MembershipNo      string      `json:"membership_no"`
	CompanyName       string      `json:"company_name"`
	Title             string      `json:"title"`
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Gender            string      `json:"gender"`
	DateOfBirth       interface{} `json:"date_of_birth"`
	Source            string      `json:"source"`
}

type CustomerListResponse struct {
	Data   []Customer `json:"data"`
	Status bool       `json:"status,omitempty"`
}

type CustomerCreateGetUpdateResponse struct {
	Data   Customer `json:"data"`
	Status bool     `json:"status,omitempty"`
}

type CustomerGetResponse struct {
	Data struct {
		Model Customer `json:"model"`
	} `json:"data"`
	Meta []interface{} `json:"meta"`
	// todo: loyalty field
	// todo: prepayment field
	// --- https://apidoc.thegoodtill.com/#api-Customer-ShowCustomer
	Status bool `json:"status,omitempty"`
}

type CustomerCreateUpdateRequest struct {
	Name   string `json:"group_name"`
	Active uint   `json:"active"`
}

type CustomerGroupResponseData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"group_name"`
	Active    int    `json:"active"`
	Status    string `json:"status"`
}

type CustomerGroupListResponse struct {
	Status  bool                        `json:"status"`
	Data    []CustomerGroupResponseData `json:"data"`
	Message string                      `json:"message"`
}

type CustomerGroupCreateGetUpdateResponse struct {
	Status  bool                      `json:"status"`
	Data    CustomerGroupResponseData `json:"data"`
	Message string                    `json:"message"`
}

type CustomerDeleteRequest struct{}

type CustomerGroupDeleteResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Success int `json:"success"`
	} `json:"data"`
	Message string `json:"message"`
}
