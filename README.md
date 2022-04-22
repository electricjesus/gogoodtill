# GoGoodTill

Goodtill API Client for Golang

## Quick start 

```golang
package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/electricjesus/gogoodtill/pkg/authentication"
	"github.com/electricjesus/gogoodtill/pkg/products"
)

func main() {
	log.SetLevel(log.DebugLevel)

	fmt.Println("logging in")
	params := &authentication.AuthLoginParams{
		Subdomain: "yoursubdomain",
		Username:  "yourusername",
		Password:  "actuallynotapassword",
	}
	auth := authentication.New(authentication.WithLoginParams(params))
	client, err := authentication.NewAuthenticatedClient(authentication.WithTokenSource(auth))
	if err != nil {
		panic(err)
	}

	prs := products.NewProductsAPI(client)
	fmt.Println(prs.List())
}
```