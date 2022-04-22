package products

type ProductsExtraAPI interface {
	// POST https://api.thegoodtill.com/api/products/:id/image.jpg
	Image()
	// POST https://api.thegoodtill.com/api/products/:id/duplicate
	Duplicate() // NOT IMPLEMENTED
	//
	//
	// POST https://api.thegoodtill.com/api/products/variants
	CreateVariant()
	// PUT https://api.thegoodtill.com/api/products/variants
	UpdateVariant()

	// GET https://api.thegoodtill.com/api/products/inventory
	GetInventory()
}
